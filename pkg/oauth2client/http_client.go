package oauth2client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// NewBasicClient returns a client which always sends along basic auth
// credentials.
func NewBasicClient(clientID string, clientSecret string) *basicClient {
	return &basicClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		client: http.Client{
			Timeout: time.Second * 5,
		},
	}
}

type basicClient struct {
	clientID     string
	clientSecret string

	client http.Client
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// Post sends a request to the given uri with a payload of url values.
func (c *basicClient) Post(uri string, payload url.Values) (res *http.Response, token RefreshTokenResponse, err error) {
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader([]byte(payload.Encode())))
	if err != nil {
		return
	}

	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err = c.client.Do(req)
	if err != nil {
		return
	}

	// reset body for re-reading
	err = json.NewDecoder(res.Body).Decode(&token)

	return res, token, err
}
