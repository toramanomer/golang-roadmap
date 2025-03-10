package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func (album *Album) scan(row *sql.Row) error {
	return row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
}

func main() {
	log.SetFlags(0)

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n\n", albums)

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of the added album: %v\n", albID)
}

const (
	queryAlbumByID     = "select * from album where id = ?"
	queryAlbumsByArist = "select * from album where artist = ?"
	createAlbum        = "insert into album (title, artist, price) values (?, ?, ?)"
)

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query(queryAlbumsByArist, name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

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

func albumByID(id int64) (Album, error) {
	if id <= 0 {
		return Album{}, fmt.Errorf("invalid album ID: %d", id)
	}

	var album Album

	row := db.QueryRow(queryAlbumByID, id)

	if err := album.scan(row); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return album, fmt.Errorf("albumByID %d: no such album", id)
		}
		return album, fmt.Errorf("albumByID %d: %v", id, err)
	}

	return album, nil
}

func addAlbum(album Album) (int64, error) {
	result, err := db.Exec(createAlbum, album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}
