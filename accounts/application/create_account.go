package application

import (
	"accounts/domain/entities"
	"accounts/pkg/db"
	"context"
	"fmt"
	"strings"

	// "golang.org/x/crypto/bcrypt"
	"github.com/Nerzal/gocloak/v8"
)

const (
	ClientId     = "accounts_service"
	ClientSecret = "BwJtYnEqzwgyKldF71DbcEVX1Yd0lvPW"
	Realm        = "morpheus"
	Hostname     = "http://localhost:8080"
)

type CreateAccountDTO struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Document  string `json:"document,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}

type CreateAccount struct {
	Repository db.Repository
}

func NewCreateAccount(Repository db.Repository) *CreateAccount {
	return &CreateAccount{
		Repository: Repository,
	}
}

func (c *CreateAccount) Create(data *CreateAccountDTO) (*entities.Account, error) {
	accountExists, err := c.Repository.Exists(data.Email)
	if err != nil {
		return nil, err
	}

	if accountExists {
		return nil, ErrEmailAlreadyInUse
	}

	account, err := entities.NewAccount(
		data.Name,
		data.Email,
		data.Password,
		data.AvatarUrl,
		data.BirthDate,
		data.Document,
	)
	if err != nil {
		return nil, err
	}

	// hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }
	// account.Password = string(hash)

	keycloack := gocloak.NewClient(Hostname)

	clientToken, err := keycloack.LoginClient(context.Background(), ClientId, ClientSecret, Realm)
	if err != nil {
		fmt.Println(err)
		fmt.Println("client error")
		return nil, err
	}

	keycloackUser := gocloak.User{
		FirstName:     gocloak.StringP(strings.Split(account.Name, " ")[0]),
		LastName:      gocloak.StringP(strings.Split(account.Name, " ")[1]),
		Email:         gocloak.StringP(account.Email),
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(false),
		Username:      gocloak.StringP(strings.Join(strings.Split(account.Name, " "), "")),
	}

	err = c.Repository.Create(account)
	if err != nil {
		return nil, err
	}

	newUserId, err := keycloack.CreateUser(context.Background(), clientToken.AccessToken, Realm, keycloackUser)
	if err != nil {
		fmt.Println("user error")
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("New user created %s\n", newUserId)

	return account, nil
}
