package main

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/goDotEnvVariable"
	_ "github.com/lib/pq"
)

// var db *sql.DB

const (
	host   = "localhost"
	port   = 5432
	dbname = "learngo"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	postgresqlDbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, goDotEnvVariable.GetEnvVariable("DBUSER"), goDotEnvVariable.GetEnvVariable("DBPASSWORD"), dbname,
	)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")

	albums, err := albumsByArtist("John Coltrane", db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

// albumsByArtist queries fo Albums that have the specified artist name.
func albumsByArtist(name string, db *sql.DB) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM albums WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
