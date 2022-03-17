package auth

import (
	"events/config/env"
	"events/framework/encrypter"
	"fmt"

	"github.com/Nerzal/gocloak/v7"
)

var (
	hostname = fmt.Sprintf("http://%s:%d", env.KEYCLOACK_HOST, env.KEYCLOACK_PORT)
)

type KeyclaockAuthProvider struct {
	client    gocloak.GoCloak
	encrypter encrypter.Encrypter
}

func NewKeycloackauthProvider(encrypter encrypter.Encrypter) *KeyclaockAuthProvider {
	return &KeyclaockAuthProvider{
		client:    gocloak.NewClient(hostname),
		encrypter: encrypter,
	}
}

func (k *KeyclaockAuthProvider) AuthUser(accessToken string) (AuthUserInfo, error) {
	accountId, err := k.encrypter.DecryptToken(accessToken, env.KEYCLOACK_PUBLIC_RSA_KEY)
	if err != nil {
		return AuthUserInfo{}, err
	}

	userInfo := AuthUserInfo{
		Id: accountId,
	}

	return userInfo, nil
}
