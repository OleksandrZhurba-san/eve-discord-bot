package auth

import (
	"net/url"
)

const (
	AuthEndpoint = "https://login.eveonline.com/v2/oauth/authorize"
)

func BuidlAuthUrl(cliendID, callbackURL, state, scope string) string {
	u, _ := url.Parse(AuthEndpoint)
	query := u.Query()
	query.Set("response_type", "code")
	query.Set("redirect_uri", callbackURL)
	query.Set("client_id", cliendID)
	query.Set("scope", scope)
	query.Set("state", state) // for callback validation

	u.RawQuery = query.Encode()
	return u.String()
}
