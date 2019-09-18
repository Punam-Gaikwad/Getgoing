package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

//Connect - connects to the mysql database
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/todo_api_db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database...")
	con = db
	return db
}
