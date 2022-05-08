package oauth2client

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	config "app-api/configs"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// The following provides the setup required for the client to perform the "Authorization Code" flow with PKCE in order
// to obtain an access token for public/untrusted clients.

const cookiePKCE = "isPKCE"

var (
	// pkceCodeVerifier stores the generated random value which the client will on-send to the auth server with the received
	// authorization code. This way the oauth server can verify that the base64URLEncoded(sha265(codeVerifier)) matches
	// the stored code challenge, which was initially sent through with the code+PKCE authorization request to ensure
	// that this is the original user-agent who requested the access token.
	pkceCodeVerifier string

	// pkceCodeChallenge stores the base64(sha256(codeVerifier)) which is sent from the
	// client to the auth server as required for PKCE.
	pkceCodeChallenge string
)

var Provider *oidc.Provider
var Verifier *oidc.IDTokenVerifier

// TODO: move to config package
var ClientConf = oauth2.Config{
	ClientID:     "es-app-api-real",
	ClientSecret: "some-secret",
	RedirectURL:  os.Getenv("SERVICE_ENDPOINT") + "/callback",
	Scopes:       []string{oidc.ScopeOpenID, "offline"},
	Endpoint: oauth2.Endpoint{
		TokenURL:  os.Getenv("TOKEN_URL"),
		AuthURL:   os.Getenv("TOKEN_AUTH_URL"),
		AuthStyle: oauth2.AuthStyleInHeader,
	},
}

func init() {
	var initErr error
	// todo provider get value from config
	Provider, initErr = oidc.NewProvider(context.Background(), os.Getenv("TOKEN_PROVIDER"))
	if initErr != nil {
		panic(initErr)
	}
	Verifier = Provider.Verifier(&oidc.Config{ClientID: ClientConf.ClientID})

}
func CallbackHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		c := ClientConf
		// rw := ctx.Writer
		req := ctx.Request

		//
		fmt.Println("call?")
		AuthCode := req.FormValue("code")
		fmt.Println(AuthCode)
		// time.Sleep(10*time.Second)
		//
		// codeVerifier := resetPKCE(rw)
		// rw.Write([]byte(`<h1>Callback site</h1><a href="/">Go back</a>`))
		// rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		// if req.URL.Query().Get("error") != "" {
		// 	rw.Write([]byte(fmt.Sprintf(`<h1>Error!</h1>
		// 	Error: %s<br>
		// 	Error Hint: %s<br>
		// 	Description: %s<br>
		// 	<br>`,
		// 		req.URL.Query().Get("error"),
		// 		req.URL.Query().Get("error_hint"),
		// 		req.URL.Query().Get("error_description"),
		// 	)))
		// 	return
		// }

		// client := newBasicClient(c.ClientID, c.ClientSecret)
		// if req.URL.Query().Get("revoke") != "" {
		// 	revokeURL := strings.Replace(c.Endpoint.TokenURL, "token", "revoke", 1)
		// 	payload := url.Values{
		// 		"token_type_hint": {"refresh_token"},
		// 		"token":           {req.URL.Query().Get("revoke")},
		// 	}
		// 	resp, body, err := client.Post(revokeURL, payload)
		// 	if err != nil {
		// 		rw.Write([]byte(fmt.Sprintf(`<p>Could not revoke token %s</p>`, err)))
		// 		return
		// 	}

		// 	rw.Write([]byte(fmt.Sprintf(`<p>Received status code from the revoke endpoint:<br><code>%d</code></p>`, resp.StatusCode)))
		// 	if body != "" {
		// 		rw.Write([]byte(fmt.Sprintf(`<p>Got a response from the revoke endpoint:<br><code>%s</code></p>`, body)))
		// 	}

		// 	rw.Write([]byte(fmt.Sprintf(`<p>These tokens have been revoked, try to use the refresh token by <br><a href="%s">by clicking here</a></p>`, "?refresh="+url.QueryEscape(req.URL.Query().Get("revoke")))))
		// 	rw.Write([]byte(fmt.Sprintf(`<p>Try to use the access token by <br><a href="%s">by clicking here</a></p>`, "/protected?token="+url.QueryEscape(req.URL.Query().Get("access_token")))))

		// 	return
		// }

		// if req.URL.Query().Get("refresh") != "" {
		// 	payload := url.Values{
		// 		"grant_type":    {"refresh_token"},
		// 		"refresh_token": {req.URL.Query().Get("refresh")},
		// 		"scope":         {"openid", "offline", "photos.read"},
		// 	}
		// 	_, body, err := client.Post(c.Endpoint.TokenURL, payload)
		// 	if err != nil {
		// 		rw.Write([]byte(fmt.Sprintf(`<p>Could not refresh token %s</p>`, err)))
		// 		return
		// 	}
		// 	rw.Write([]byte(fmt.Sprintf(`<p>Got a response from the refresh grant:<br><code>%s</code></p>`, body)))
		// 	return
		// }

		// if req.URL.Query().Get("code") == "" {
		// 	rw.Write([]byte(fmt.Sprintln(`<p>Could not find the authorize code. If you've used the implicit grant, check the
		// 	browser location bar for the
		// 	access token <small><a href="http://en.wikipedia.org/wiki/Fragment_identifier#Basics">(the server side does not have access to url fragments)</a></small>
		// 	</p>`,
		// 	)))
		// 	return
		// }

		// rw.Write([]byte(fmt.Sprintf(`<p>Amazing! You just got an authorize code!:<br><code>%s</code></p>
		// <p>Click <a href="/">here to return</a> to the front page</p>`,
		// 	req.URL.Query().Get("code"),
		// )))

		// // We'll check whether we sent a code+PKCE request, and if so, send the code_Verifier along when requesting the access token.
		var opts []oauth2.AuthCodeOption

		// if isPKCE(req) {
		// 	opts = append(opts, oauth2.SetAuthURLParam("code_Verifier", codeVerifier))
		// }

		token, err := c.Exchange(context.Background(), req.URL.Query().Get("code"), opts...)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			logrus.Info("missing ID token")
		}
		logrus.Info("***************************")
		logrus.Info(rawIDToken)

		// Parse and verify ID Token payload.
		idToken, err := Verifier.Verify(ctx, rawIDToken)
		if err != nil {
			// handle error
			logrus.Info(err)
		}
		logrus.Info(idToken)
		// Extract custom claims
		var claims struct {
			Email    string `json:"email"`
			Verified bool   `json:"email_verified"`
		}
		if err := idToken.Claims(&claims); err != nil {
			logrus.Info(err)
			// handle error
		}

		// rw.Write([]byte(fmt.Sprintf(`<p>Cool! You are now a proud token owner.<br>
		// <ul>
		// 	<li>
		// 		Access token (click to make <a href="%s">authorized call</a>):<br>
		// 		<code>%s</code>
		// 	</li>
		// 	<li>
		// 		Refresh token (click <a href="%s">here to use it</a>) (click <a href="%s">here to revoke it</a>):<br>
		// 		<code>%s</code>
		// 	</li>
		// 	<li>
		// 		Extra info: <br>
		// 		<code>%s</code>
		// 	</li>
		// </ul>`,
		// 	"/protected?token="+token.AccessToken,
		// 	token.AccessToken,
		// 	"?refresh="+url.QueryEscape(token.RefreshToken),
		// 	"?revoke="+url.QueryEscape(token.RefreshToken)+"&access_token="+url.QueryEscape(token.AccessToken),
		// 	token.RefreshToken,
		// 	token,
		// )))
		// set cookie for authen router
		// todo samesite in config
		logrus.Info(os.Getenv("MODE"))
		if config.IsProduction() {
			logrus.Info("setting cookieeee")
			ctx.SetSameSite(http.SameSiteLaxMode)
			logrus.Info(os.Getenv("SAMESITE"))
			ctx.SetCookie("__shop_id_token", rawIDToken, 28800, "/", os.Getenv("SAMESITE"), false, false)
			ctx.SetCookie("__shop_refresh_token", token.RefreshToken, 28800, "/", os.Getenv("SAMESITE"), false, false)
			logrus.Info("setting cookieeee done")
		} else if config.IsStaging() {
			logrus.Info("setting cookieeee none")
			ctx.SetSameSite(http.SameSiteNoneMode)
			ctx.SetCookie("__shop_id_token", rawIDToken, 28800, "/", "", true, true)
			ctx.SetCookie("__shop_refresh_token", token.RefreshToken, 28800, "/", "", true, true)
			logrus.Info("setting cookieeee done")
		} else {
			logrus.Info("setting cookieeee")
			ctx.SetSameSite(http.SameSiteLaxMode)
			logrus.Info(os.Getenv("SAMESITE"))
			ctx.SetCookie("__shop_id_token", rawIDToken, 28800, "/", "", false, false)
			ctx.SetCookie("__shop_refresh_token", token.RefreshToken, 28800, "/", "", false, false)
			logrus.Info("setting cookieeee done")
		}

		shopDomain := idToken.Nonce
		frontendHost := shopDomain
		if config.IsDev() {
			frontendHost = fmt.Sprintf("http://%s:%s", shopDomain, config.Dev.FrontendDevPort)
		} else {
			frontendHost = fmt.Sprintf("https://%s", frontendHost)
		}
		ctx.Redirect(http.StatusMovedPermanently, frontendHost)
		return
		// ctx.JSON(200, gin.H{
		// 	"access_token":  token.AccessToken,
		// 	"id_token":      rawIDToken,
		// 	"refresh_token": token.RefreshToken,
		// })
	}
}

// isPKCE detects whether a PKCE auth request was made.
func isPKCE(r *http.Request) bool {
	cookie, err := r.Cookie(cookiePKCE)
	if err != nil {
		return false
	}

	return cookie.Value == "true"
}

// resetPKCE cleans up PKCE details and returns the code Verifier.
func resetPKCE(w http.ResponseWriter) (codeVerifier string) {
	// remove cookie that informs the client the callback request was a PKCE
	// request.
	http.SetCookie(w, &http.Cookie{
		Name:    cookiePKCE,
		Path:    "/",
		Expires: time.Unix(0, 0),
	})

	codeVerifier = pkceCodeVerifier
	pkceCodeVerifier = ""

	return codeVerifier
}
