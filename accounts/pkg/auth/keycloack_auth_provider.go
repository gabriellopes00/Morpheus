package auth

import (
	"accounts/config/env"
	"accounts/pkg/encrypter"
	"context"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v7"
	"go.uber.org/zap"
)

var (
	clientId     = env.KEYCLOACK_CLIENT_ID
	clientSecret = env.KEYCLOACK_CLIENT_SECRET
	realm        = env.KEYCLOACK_REALM
	hostname     = fmt.Sprintf("http://%s:%d", env.KEYCLOACK_HOST, env.KEYCLOACK_PORT)
)

type KeyclaockAuthProvider struct {
	Client    gocloak.GoCloak
	Encrypter encrypter.Encrypter
}

func NewKeycloackauthProvider(Encrypter encrypter.Encrypter) *KeyclaockAuthProvider {
	return &KeyclaockAuthProvider{
		Client:    gocloak.NewClient(hostname),
		Encrypter: Encrypter,
	}
}

func (k *KeyclaockAuthProvider) CreateUser(user AuthProviderUser) error {

	clientToken, err := k.Client.LoginClient(context.Background(), clientId, clientSecret, realm)
	if err != nil {
		return err
	}

	keycloackUser := gocloak.User{
		FirstName:        gocloak.StringP(strings.Split(user.Name, " ")[0]),
		LastName:         gocloak.StringP(strings.Split(user.Name, " ")[1]),
		Email:            gocloak.StringP(user.Email),
		Enabled:          gocloak.BoolP(true),
		EmailVerified:    gocloak.BoolP(false),
		Username:         gocloak.StringP(strings.Split(user.Email, "@")[0]),
		CreatedTimestamp: gocloak.Int64P(user.CreatedAt.Unix()),
		Attributes:       &map[string][]string{"account_id": {user.Id}},
	}

	newUserId, err := k.Client.CreateUser(context.Background(), clientToken.AccessToken, realm, keycloackUser)
	if err != nil {
		return err
	}

	err = k.Client.SetPassword(context.Background(), clientToken.AccessToken, newUserId, realm, user.Password, false)
	if err != nil {
		return err
	}

	return nil
}

func (k *KeyclaockAuthProvider) SignInUser(email, password string) (Token, error) {

	keycloackToken, err := k.Client.Login(
		context.Background(),
		clientId,
		clientSecret,
		realm,
		strings.Split(email, "@")[0],
		password,
	)
	if err != nil {
		return Token{}, err
	}

	token := Token{
		AccessToken:      keycloackToken.AccessToken,
		ExpiresIn:        keycloackToken.ExpiresIn,
		RefreshToken:     keycloackToken.RefreshToken,
		RefreshExpiresIn: keycloackToken.RefreshExpiresIn,
	}

	return token, nil
}

func (k *KeyclaockAuthProvider) AuthUser(accessToken string) (AuthUserInfo, error) {
	accountId, err := k.Encrypter.DecryptToken(accessToken)
	if err != nil {
		zap.Error(err)
		// return an custom error
		return AuthUserInfo{}, err
	}

	userInfo := AuthUserInfo{
		Id: accountId,
	}

	return userInfo, nil
}

func (k *KeyclaockAuthProvider) RefreshAuth(refreshToken string) (Token, error) {

	keycloackToken, err := k.Client.RefreshToken(context.Background(), refreshToken, clientId, clientSecret, realm)
	if err != nil {
		zap.Error(err)
		// return an custom error
		return Token{}, err
	}

	token := Token{
		AccessToken:      keycloackToken.AccessToken,
		ExpiresIn:        keycloackToken.ExpiresIn,
		RefreshToken:     keycloackToken.RefreshToken,
		RefreshExpiresIn: keycloackToken.RefreshExpiresIn,
	}

	return token, nil
}
