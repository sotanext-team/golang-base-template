package http

import (
	"net/http"

	"app-api/pkg/oauth2client"

	"github.com/gin-gonic/gin"
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

type AuthHandler struct {
}

func NewAuthHandler() AuthHandler {
	// shopUseCase := usecase.NewShopUseCase(db)
	return AuthHandler{
		// shopUseCase: shopUseCase,
	}
}

func (o *AuthHandler) Index(c *gin.Context) {

	// rawIDToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6InB1YmxpYzo2NGMwZDgyMC03NTkxLTQzYWItYmMyMS01NDJkYTIxYjEzOTIifQ.eyJGb28iOiJoZWhlIiwiYXRfaGFzaCI6Im9VQ1VSaHhDYXVSbG1ZQTRiRlFkNnciLCJhdWQiOlsiZXMtYXBwLWFwaS1yZWFsIl0sImF1dGhfdGltZSI6MTY0MDY2MTQ1NSwiZXhwIjoxNjQwNjY3NjM0LCJpYXQiOjE2NDA2NjQwMzQsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6OTAwMC8iLCJqdGkiOiIwYzcxOWJhZi02YmJlLTQ3OTUtYTE4My01NTRjNzAyMzJhMjEiLCJub25jZSI6InNvbWUtcmFuZG9tLW5vbmNlIiwicmF0IjoxNjQwNjY0MDI5LCJzaWQiOiI4NGYwZmNhYi0wZjYyLTQ2NTktYTNiMi00ZWY5YjQwNWYxZjEiLCJzdWIiOiJmb29AYmFyLmNvbSJ9.4057ElDxktXLmw1C8W3fvYltbkTZrKhy17bACX5Bd6eGpC1-p7SnlNCGKhkI6Nh4WLzaX8eB7Pb3sHXKOgsVxYeSxY0UkwNpeKpjDekO89Q6nUY6XTbqgTCHZdC457JNBoVHVqZQaSRs3H_7_gz0plrFy8NTpd5tLo3iFKIbQ_y1W6wC8A1s88NnHgf8gJzHs4ksQzAXyqrMUetCPRzDW4oKcGSlYfBGPoJIRvlBGVSY543EHsrKQKWJoWWbm4nMfAUOyrOB39bDOyioIZU2u-iiN5mWARBxDjfsTdV7ikTFqkt0EE6mxe9DpHu8lCfFOaWar2mBSOIZG2QV3Vl5JZKHsbCF_pmRFH_GSbrFnFiyM1SIt-4pWWagalw7ULDBEaRbBSVExo67cDOnA0sab0to7IrGS47MNJOr-QelCkCD9BgAGIu5p91l30KeX8EbgQwCbCGzI-z8u8AyUewGniWC9MWMHr_68umX9gixhnRwI_F7SB5ORke1kPdRdsaWYVy-IruXc06ZfKnTJ8CBaMrnMKav9edGwIlE5MK6vQWn2Ol_sPvSFZiLt10yLWA_1Rv2SRJfcfcMCExHoxMAmFMg6JuTM6Gso0ejh9orrpnwC082fn1Hl2QpzHo8mIopih6a2KX7-L9SvrEiAgSqyqYjKTIx_VTO_1Nfngn3vI8"
	// // Parse and verify ID Token payload.
	// idToken, err := oauth2client.Verifier.Verify(c, rawIDToken)
	// if err != nil {
	// 	// handle error
	// 	logrus.Info(err)
	// }
	// fmt.Printf("%+v\n", idToken)

	authURL := oauth2client.ClientConf.AuthCodeURL("some-random-state") + "&audience=es-app-api-real" + "&nonce=dragon-born-xyzkmkssss"
	c.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"AuthURL": authURL,
		},
	)
}
