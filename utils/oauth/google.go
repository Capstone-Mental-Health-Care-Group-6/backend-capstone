package oauth

import (
	"FinalProject/configs"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthGoogleInterface interface {
	GenerateState() (string, error)
	AuthCodeURL(state string) (string, error)
	Exchange(code string) (*oauth2.Token, error)
	GetEmail(token *oauth2.Token) (string, error)
}

type OauthGoogleConfig struct {
	OAuthConf *oauth2.Config
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func NewOauthGoogleConfig(c configs.ProgrammingConfig) OauthGoogleInterface {
	return &OauthGoogleConfig{
		OAuthConf: &oauth2.Config{
			ClientID:     c.OauthGoogleClientID,
			ClientSecret: c.OauthGoogleClientSecret,
			RedirectURL:  c.OauthGoogleRedirectURL,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (o *OauthGoogleConfig) GenerateState() (string, error) {
	const stateLength = 32
	bytes := make([]byte, stateLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (o *OauthGoogleConfig) AuthCodeURL(state string) (string, error) {
	return o.OAuthConf.AuthCodeURL(state), nil
}

func (o *OauthGoogleConfig) Exchange(code string) (*oauth2.Token, error) {
	return o.OAuthConf.Exchange(context.Background(), code)
}

func (o *OauthGoogleConfig) GetEmail(token *oauth2.Token) (string, error) {
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return "", fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed decoding response: %s", err.Error())
	}

	email := result["email"]

	return email.(string), nil
}
