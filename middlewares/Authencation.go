package middlewares

import (
	graphModel "app-api/graph/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"app-api/configs"
	config "app-api/configs"
	"app-api/constants"
	"app-api/db"
	"app-api/ent/shop"
	"app-api/models"
	"app-api/pkg/oauth2client"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

var PublicKey string

func init() {
	rsaPublicKey := configs.RSA.PublicKey
	l, _ := base64.StdEncoding.DecodeString(rsaPublicKey)
	PublicKey = string(l)

}

type IDTokenClaims struct {
	Foo      string   `json:"Foo"`
	AtHash   string   `json:"at_hash"`
	Aud      []string `json:"aud"`
	AuthTime int      `json:"auth_time"`
	Email    string   `json:"email"`
	Exp      int      `json:"exp"`
	Iat      int      `json:"iat"`
	Iss      string   `json:"iss"`
	Jti      string   `json:"jti"`
	Nonce    string   `json:"nonce"`
	Rat      int      `json:"rat"`
	Sid      string   `json:"sid"`
	Sub      string   `json:"sub"`
	Username string   `json:"username"`
}

func AuthenRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		var rawIDToken, cookieValue string
		var err error
		// Public graphql schema for gateway
		if path == "/graphql/query" || c.Request.Method == "POST" {
			// Read request body
			byteBody, _ := ioutil.ReadAll(c.Request.Body)

			bodyData := models.ApolloGraphqlGatewayBody{}
			if err := json.Unmarshal(byteBody, &bodyData); err != nil {
				c.JSON(400, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			// Reassign the body because the body will empty after parsing to bytes array
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))

			if bodyData.Query == constants.GRAPHQL_GATEWAY_BODY_REQUEST {
				c.Next()
				return
			}
		}
		header := c.GetHeader("Authorization")
		if strings.HasPrefix(header, "Bearer ") {
			rawIDToken = header[7:]
		}
		cookieValue, err = c.Cookie("__shop_id_token")
		if err != nil && rawIDToken == "" {
			c.JSON(401, gin.H{
				"status":  false,
				"message": "token not found",
			})
			c.Abort()
			return
		} else {
			if cookieValue > "" {
				rawIDToken = cookieValue
			}
		}
		// we don't really need refresh token, we only need them in case id token expired
		refreshToken, _ := c.Cookie("__shop_refresh_token")

		idToken, err := oauth2client.Verifier.Verify(c, rawIDToken)
		fmt.Println(rawIDToken)
		// fake err
		// err = errors.New("huhu")
		if err != nil && refreshToken != "" {
			// TODO request new idtoken by refresh token
			var token oauth2client.RefreshTokenResponse
			fmt.Println(refreshToken)
			clientConf := oauth2client.ClientConf
			client := oauth2client.NewBasicClient(clientConf.ClientID, clientConf.ClientSecret)

			payload := url.Values{
				"grant_type":    {"refresh_token"},
				"refresh_token": {refreshToken},
				"scope":         clientConf.Scopes,
			}
			_, token, err = client.Post(clientConf.Endpoint.TokenURL, payload)
			if err != nil {
				c.JSON(401, gin.H{
					"status":  false,
					"message": "can't get access token with refresh token",
				})
				c.Abort()
				return
				// TODO wrap err
				c.JSON(401, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			logrus.Info(token.IDToken)
			logrus.Info(token.RefreshToken)
			rawIDToken = token.IDToken
			if config.IsProduction() {
				logrus.Info("setting cookieeee")
				c.SetSameSite(http.SameSiteLaxMode)
				logrus.Info(os.Getenv("SAMESITE"))
				c.SetCookie("__shop_id_token", rawIDToken, 28800, "/", os.Getenv("SAMESITE"), false, false)
				c.SetCookie("__shop_refresh_token", token.RefreshToken, 28800, "/", os.Getenv("SAMESITE"), false, false)
				logrus.Info("setting cookieeee done")
			} else if config.IsStaging() {
				logrus.Info("setting cookieeee none")
				c.SetSameSite(http.SameSiteNoneMode)
				c.SetCookie("__shop_id_token", rawIDToken, 28800, "/", "", true, true)
				c.SetCookie("__shop_refresh_token", token.RefreshToken, 28800, "/", "", true, true)
				logrus.Info("setting cookieeee done")
			} else {
				logrus.Info("setting cookieeee")
				c.SetSameSite(http.SameSiteLaxMode)
				logrus.Info(os.Getenv("SAMESITE"))
				c.SetCookie("__shop_id_token", rawIDToken, 28800, "/", "", false, false)
				c.SetCookie("__shop_refresh_token", token.RefreshToken, 28800, "/", "", false, false)
				logrus.Info("setting cookieeee done")
			}
			idToken, _ = oauth2client.Verifier.Verify(c, rawIDToken)
		}

		entClient := db.GetClient()

		var claims IDTokenClaims
		// parse token to struct
		idToken.Claims(&claims)
		if err != nil {
			c.JSON(401, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		// we place shop domain in nonce
		shopDomain := claims.Nonce
		if err != nil {
			c.JSON(400, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		shopData, err := entClient.Shop.Query().Where(shop.DefaultDomain(shopDomain)).Only(c.Request.Context())
		if err != nil {
			c.JSON(400, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// TODO: set currentUser also
		c.Set("currentShop", shopData)
		uId, _ := strconv.ParseUint(claims.Sub, 0, 64)
		currentUser := &graphModel.User{
			ID:       uId,
			UserName: &claims.Username,
			Email:    &claims.Email,
		}
		c.Set("currentUser", currentUser)
		c.Next()
	}
}

func AuthenShopRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func ValidateTokenRSA(tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key"
		r := strings.NewReader(PublicKey)
		hmacSampleSecret, _ := ioutil.ReadAll(r)
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(hmacSampleSecret)
		if err != nil {
			return nil, err

		}
		return verifyKey, nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("%+v\n", claims)
		return claims, nil
	}
	return jwt.MapClaims{}, err
}
