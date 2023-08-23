package nicrudns

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

func (client *Client) GetOauth2Client() (*http.Client, error) {
	ctx := context.TODO()

	if client.oauth2client != nil {
		return client.oauth2client, nil
	}

	oauth2Config := oauth2.Config{
		ClientID:     client.provider.OAuth2ClientID,
		ClientSecret: client.provider.OAuth2SecretID,
		Endpoint: oauth2.Endpoint{
			TokenURL:  TokenURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		Scopes: []string{OAuth2Scope},
	}

	oauth2Token, err := oauth2Config.PasswordCredentialsToken(ctx, client.provider.Username, client.provider.Password)
	if err != nil {
		return nil, errors.Wrap(err, AuthorizationError.Error())
	}

	client.oauth2client = oauth2Config.Client(ctx, oauth2Token)

	return client.oauth2client, nil
}
