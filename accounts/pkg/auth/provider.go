package auth

import "time"

type AuthProviderUser struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type AuthUserCredentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Token struct {
	AccessToken      string `json:"access_token,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	RefreshExpiresIn int    `json:"refresh_expires_in,omitempty"`
}

type AuthUserInfo struct {
	Id string `json:"id,omitempty"`
}

type AuthProvider interface {
	CreateUser(user AuthProviderUser) error
	SignInUser(credentials AuthUserCredentials) (Token, error)
	AuthUser(accessToken string) (AuthUserInfo, error)
	RefreshAuth(refreshToken string) (Token, error)
}
