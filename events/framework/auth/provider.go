package auth

type AuthUserCredentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthUserInfo struct {
	Id string `json:"id,omitempty"`
}

type AuthProvider interface {
	AuthUser(accessToken string) (AuthUserInfo, error)
}
