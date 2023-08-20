package albums

import (
	"database/sql"
	"fmt"

	"example.com/dbConnect"
	_ "github.com/lib/pq"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func AlbumsByArtist(name string) ([]Album, error) {
	db, err := sql.Open("postgres", dbConnect.ConnectionConfig())
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var records []Album

	rows, err := db.Query("SELECT * FROM albums WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
		}
		records = append(records, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
	}
	return records, nil
}
