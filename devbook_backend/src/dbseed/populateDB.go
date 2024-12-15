package dbseed

import (
	"devbook_backend/src/database"
	"devbook_backend/src/models"
	"devbook_backend/src/repository"
	"fmt"
	"log"

	"github.com/go-faker/faker/v4"
)

func populateUserTable(quantity uint) {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("populate user function: connect database error\n%v\n", err)
		return
	}
	defer db.Close()

	NewUsersRepository := repository.NewUsersRepository(db)

	for i := 0; i < int(quantity); i++ {
		newUser := models.User{
			Name:     faker.Name(),
			Nick:     faker.Username(),
			Email:    faker.Email(),
			Password: "testing123",
		}

		newUser.Prepare("register")

		if err = NewUsersRepository.Create(newUser); err != nil {
			log.Fatal("populate user function: create user error")
			return
		}

		fmt.Println(newUser)
	}
}

func populateFollowersTable() {

}

func PopulateDatabase(quantity uint) {
	populateUserTable(quantity)
	populateFollowersTable()
}
