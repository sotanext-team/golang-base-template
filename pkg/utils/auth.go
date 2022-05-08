package utils

import (
	"app-api/ent"
	graphModel "app-api/graph/models"
	"context"
	"encoding/base64"
	"fmt"
	"html"
	"strings"
	"time"

	"app-api/configs"
	"app-api/db"
	"app-api/ent/shop"
	"app-api/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	esUtils "github.com/es-hs/es-helper/utils"
)

// TokenClaim This is the cliam object which gets parsed from the authorization header
type TokenClaim struct {
	*jwt.StandardClaims
	User models.UserToken `json:"user"`
	Shop models.ShopToken `json:"shop"`
}

// ShopTokenClaim This is the cliam object which gets parsed from the authorization header
type ShopTokenClaim struct {
	*jwt.StandardClaims
	models.ShopToken
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

// SecretKey get secret key
func SecretKey() string {
	// key := os.Getenv("RSA_PRIVATE_KEY")
	rsaPrivateKey := configs.RSA.PrivateKey
	l, _ := base64.StdEncoding.DecodeString(rsaPrivateKey)
	return string(l)
	// return os.Getenv("RSA_PRIVATE_KEY")
}

// CreateLoginToken login token generator
func CreateLoginToken(user models.UserCheck) (string, error) {
	expires := time.Now().Add(time.Hour * 1000000).Unix()
	token := jwt.New(jwt.SigningMethodRS256)

	client := db.GetClient()
	ctx := context.Background()
	shopResult, err := client.Debug().Shop.Query().Where(shop.DefaultDomainEQ("noragem.eshs.com")).First(ctx)

	shopId := esUtils.NumberToString(shopResult.ID)
	// TODO: remove it when DragonBorn finish auth
	token.Claims = &TokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expires,
		},
		User: models.UserToken{
			ID:       1,
			Username: "noragem",
			Email:    "nora@gempages.help",
		},
		Shop: models.ShopToken{
			ID:            shopId,
			DefaultDomain: shopResult.DefaultDomain,
			UserID:        1,
		},
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(SecretKey()))
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
		// f, fn, l := helpers.WhereAmI()
	}

	return tokenString, err
}

// CreateShopLoginToken login token generator
func CreateShopLoginToken(user models.UserCheck, shop models.Shop) (string, error) {
	expires := time.Now().Add(time.Hour * 18).Unix()
	token := jwt.New(jwt.SigningMethodRS256)

	token.Claims = &ShopTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expires,
		},
		models.ShopToken{
			// ID:            shop.ID,
			DefaultDomain: shop.DefaultDomain,
			UserID:        uint(user.ID),
		},
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(SecretKey()))
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
		// f, fn, l := helpers.WhereAmI()
	}

	return tokenString, err
}

// GetAuthenInfo get current user and current shop from context
func GetAuthenInfo(c context.Context) (currentUser *graphModel.User, currentShop *ent.Shop, err error) {
	userValue := c.Value("currentUser")
	var ok bool
	if userValue != nil {
		if currentUser, ok = userValue.(*graphModel.User); !ok {
			err = fmt.Errorf("%s", "Wrong type for current user value in context")
			return nil, nil, err
		}
	}
	shopValue := c.Value("currentShop")
	if shopValue != nil {
		if currentShop, ok = shopValue.(*ent.Shop); !ok {
			err = fmt.Errorf("%s", "Wrong type for current shop value in context")
			return nil, nil, err
		}
	}
	return currentUser, currentShop, err
}
