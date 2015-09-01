package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {

	// Characters that indicate some between-song promo
	var junk = []string{"xmu", "@", "#", "p et", "a et"}
	// Get and Verify DB Connection
	conn := getLoginString()
	db := getDBConn(conn)
	defer db.Close()

	// Make sure we aren't restarting and putting in the same value
	last_song, last_artist, err := getLastPlayedSong(db)
	if err != nil {
		panic(err.Error())
	}

	// Prepare our insertion query
	stmt, err := db.Prepare("insert into xmu2 (song, artist, datetime) VALUES (?, ?, NOW())")
	if err != nil {
		panic(err.Error())
	}

	// Get currently playing song and insert if it hasn't been already
	fmt.Printf("Starting...\n")
	for {
		song, artist, err := getCurrentSong()
		if err != nil {
			panic(err.Error())
		}
		if (last_song != song || last_artist != artist) && !isJunk(song, artist, junk) {
			_, err := stmt.Exec(song, artist)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Wrote: %s - %s\n", song, artist)
			last_song, last_artist = song, artist
			time.Sleep(120 * time.Second)
		} else {
			fmt.Printf("Collision...\n")
			time.Sleep(60 * time.Second)
		}
	}
}
