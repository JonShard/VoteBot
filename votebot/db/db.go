package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Init initiallizes the database.
func Init() {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3308)/votebotdb")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	insert, err := db.Query("INSERT INTO users (discordID, username) VALUES ('Hans', '238407932238799118');")
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	fmt.Println("Success!")
}
