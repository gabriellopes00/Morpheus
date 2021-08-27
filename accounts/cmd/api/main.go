package main

import (
	"accounts/entities"
	"accounts/infra"
	"accounts/infra/db"
	"accounts/usecases"
	"fmt"
	"log"
)

func main() {
	connection := db.NewPostgresDb()

	database, err := connection.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer database.Close()

	usecase := usecases.NewCreateAccount(
		infra.NewUUIDGenerator(),
		infra.NewBcryptHasher(),
		db.NewPgAccountRepository(database),
	)

	account, err := usecase.Create(entities.Account{
		Name:      "Gabriel Luís Lopes",
		Email:     "gabriellopes00@gmail.com",
		Password:  "my_passwd_123",
		AvatarUrl: "https://avatars.githubusercontent.com/u/69465943?v=4",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(*account)
}
