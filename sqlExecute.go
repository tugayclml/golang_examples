package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_DSN = "user=postgres password=123456 dbname=test sslmode=disable"
)

type User struct {
	Id int
	Email string
	Password string
}

func main()  {
	db, err := sql.Open("postgres", DB_DSN)

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	var user User
	user.Id = 1
	user.Email = "test6@test.com"
	user.Password = "123456789"

	insertSql := fmt.Sprintf("INSERT INTO users (id, email, password) VALUES (%d, '%s', '%s')", user.Id, user.Email, user.Password)
	fmt.Println(insertSql)
	_,err = db.Exec(insertSql)

	if err != nil{
		log.Fatal("Failed to execute insert query: ", err)
	}

	userSql := "SELECT Id, Email, Password FROM Users where id = $1"

	err = db.QueryRow(userSql, 1).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil{
		log.Fatal("Failed to execute query: ", err)
	}
	fmt.Printf("Hi %s, welcome back!\n", user.Email)
}
