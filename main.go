package main

import (
	"fmt"
	"log"

	"example.com/albums"
	"example.com/dbConnect"
	_ "github.com/lib/pq"
)

func main() {
	dbConnect.TestConnection()

	albums, err := albums.AlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}
