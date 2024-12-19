package dbseed

import (
	"devbook_backend/src/database"
	"devbook_backend/src/models"
	"devbook_backend/src/repository"
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

func populateUsersTable(quantity uint) {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("populate USERS function: connect database error\n%v\n", err)
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
			log.Fatalf("populate USERS function: create user error\n%v\n", err)
			return
		}
	}
}

func populateFollowersTable(quantity uint) {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("populate FOLLOWERS function: connect database error\n%v\n", err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into FOLLOWERS (FOLLOWING_ID, FOLLOWED_ID) values (?, ?);")
	if err != nil {
		log.Fatalf("populate FOLLOWERS function: create statement followers error\n%v\n", err)
		return
	}
	defer statement.Close()

	var userIds []string
	rows, err := db.Query("select ID from USERS;")
	if err != nil {
		log.Fatalf("populate FOLLOWERS function: query users UUIDs error\n%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var uuid string
		if err := rows.Scan(&uuid); err != nil {
			log.Fatalf("populate FOLLOWERS function: scan user UUID error\n%v\n", err)
			return
		}
		userIds = append(userIds, uuid)
	}

	for i := 0; i < int(quantity); i++ {
		followingIndex := rand.Intn(len(userIds))
		followedIndex := rand.Intn(len(userIds))

		if followingIndex == followedIndex {
			followedIndex = (followedIndex + 1) % len(userIds)
		}

		followingID := userIds[followingIndex]
		followedID := userIds[followedIndex]

		_, err := statement.Exec(followingID, followedID)
		if err != nil {
			log.Fatalf("populate FOLLOWERS function: execute statement followers error\n%v\n", err)
			return
		}
	}

}

func populatePostsTable() {
	defaultPost := models.Post{
		Title: "Default Post",
		Content: `Post default para testar o banco de dados...
		quero ver como isso pode fcar...
		`,
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("populate POSTS function: connect database error\n%v\n", err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into POSTS (TITLE, CONTENT, AUTHOR_ID, AUTHOR_NICK) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("populate POSTS function: create statement followers error\n%v\n", err)
		return
	}
	defer statement.Close()

	rows, err := db.Query("select ID, NICK from USERS where NAME like '%Dr.%';")
	if err != nil {
		log.Fatalf("populate POSTS function: query error\n%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userId string
		var userNick string

		if err := rows.Scan(&userId, &userNick); err != nil {
			log.Fatalf("populate POSTS function: scan user UUID error\n%v\n", err)
			return
		}

		userIdUUID, err := uuid.Parse(userId)
		if err != nil {
			log.Fatalf("populate POSTS function: convert string to uuid\n%v\n", err)
			return
		}

		defaultPost.AuthorId = userIdUUID
		defaultPost.AuthorNick = userNick

		_, err = statement.Exec(defaultPost.Title, defaultPost.Content, defaultPost.AuthorId, defaultPost.AuthorNick)
		if err != nil {
			log.Fatalf("populate POSTS function: execute statement followers error\n%v\n", err)
			return
		}
	}
}

func populateLikesTable() {
	var usersIdList []string
	var postIdList []string

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("populate LIKES function: connect database error\n%v\n", err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into LIKES (POST_ID, USER_ID) values (?, ?)")
	if err != nil {
		log.Fatalf("populate LIKES function: create statement followers error\n%v\n", err)
		return
	}
	defer statement.Close()

	userRows, err := db.Query("SELECT ID FROM USERS WHERE NAME LIKE '%queen%';")
	if err != nil {
		log.Fatalf("populate LIKES function: query error\n%v\n", err)
		return
	}
	defer userRows.Close()

	postsRows, err := db.Query("SELECT ID FROM POSTS")
	if err != nil {
		log.Fatalf("populate LIKES function: query error\n%v\n", err)
		return
	}
	defer postsRows.Close()

	for userRows.Next() {
		var userId string
		if err = userRows.Scan(&userId); err != nil {
			log.Fatalf("populate LIKES function: scan error\n%v\n", err)
			return
		}

		usersIdList = append(usersIdList, userId)
	}

	for postsRows.Next() {
		var postId string
		if err = postsRows.Scan(&postId); err != nil {
			log.Fatalf("populate LIKES function: scan error\n%v\n", err)
			return
		}

		postIdList = append(postIdList, postId)
	}

	for _, userId := range usersIdList {
		for _, postId := range postIdList {
			_, err = statement.Exec(postId, userId)
			if err != nil {
				log.Fatalf("populate LIKE function: statement exec error\n%v\n", err)
				return
			}
		}
	}
}

func PopulateDatabase(users, followers uint) {
	populateUsersTable(users)
	populateFollowersTable(followers)
	populatePostsTable()
	populateLikesTable()
}
