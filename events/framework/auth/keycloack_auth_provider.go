package auth

import (
	"context"
	"events/config/env"
	"fmt"

	"github.com/Nerzal/gocloak/v7"
)

var (
	clientId     = env.KEYCLOACK_CLIENT_ID
	clientSecret = env.KEYCLOACK_CLIENT_SECRET
	realm        = env.KEYCLOACK_REALM
	hostname     = fmt.Sprintf("http://%s:%d", env.KEYCLOACK_HOST, env.KEYCLOACK_PORT)
)

type KeyclaockAuthProvider struct {
	Client gocloak.GoCloak
}

func NewKeycloackauthProvider() *KeyclaockAuthProvider {
	return &KeyclaockAuthProvider{Client: gocloak.NewClient(hostname)}
}

func (k *KeyclaockAuthProvider) AuthUser(accessToken string) (AuthUserInfo, error) {

	keycloackUserInfo, err := k.Client.GetRawUserInfo(context.Background(), accessToken, realm)
	if err != nil {
		return AuthUserInfo{}, err
	}

	userInfo := AuthUserInfo{
		Id: keycloackUserInfo["account_id"].(string),
	}

	return userInfo, nil
}
