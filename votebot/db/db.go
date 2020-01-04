package db

import (
	"Votebot/votebot/cfg"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // SQL driver.
)

// Instance is the publically available database.
var Instance *sql.DB

// Song is a struct container for a song in the database.
type Song struct {
	ID        int    `json:"ID"`
	Cover     string `json:"cover"`
	Timestamp string `json:timestamp`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
}

// Init initiallizes the database.
func Init() {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", cfg.Cfg.DatabaseUser+":"+cfg.Cfg.DatabasePassword+"@tcp("+cfg.Cfg.DatabaseIP+")/"+cfg.Cfg.Database)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	Instance = db
	fmt.Println("Database connection: Success!")
}

// GetAllSongs returns a slice of Song structs that exist in the database.
func GetAllSongs() ([]Song, error) {

	rows, err := Instance.Query("SELECT * FROM songs;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	songs := make([]Song, 0)
	for rows.Next() {
		s := Song{}
		err = rows.Scan(&s.ID, &s.Cover, &s.Timestamp, &s.Title, &s.Artist, &s.Album)
		if err != nil {
			break
		}
		songs = append(songs, s)
	}
	return songs, nil
}
