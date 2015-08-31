package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
	"time"
)

// TODO: RETURN ERRORS INSTEAD OF PANICKING

func getLoginString() string {
	str, err := ioutil.ReadFile("mysql.txt")
	if err != nil {
		panic(err)
	}
	arr := strings.Split((strings.Replace(string(str), "\n", "", -1)), ", ")
	conn := fmt.Sprint(arr[0], ":", arr[1], "@/xmu")
	return string(conn)
}

func getDBConn(conn string) *sql.DB {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getLastPlayedSong(db *sql.DB) (string, string, error) {
	var (
		song   string
		artist string
	)
	err := db.QueryRow("select song, artist from xmu2 order by datetime desc limit 1").Scan(&song, &artist)
	if err != nil {
		panic(err.Error())
	}
	return song, artist, nil
}

func main() {
	conn := getLoginString()
	db := getDBConn(conn)
	defer db.Close()

	last_song, last_artist, _ := getLastPlayedSong(db)
	stmt, err := db.Prepare("insert into xmu2 (song, artist, datetime) VALUES (?, ?, NOW())")
	if err != nil {
		panic(err.Error())
	}
	for {
		song, artist, err := getCurrentSong()
		if err != nil {
			panic(err.Error())
		}
		if last_song != song || last_artist != artist {
			_, err := stmt.Exec(song, artist)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Wrote: %s - %s\n", song, artist)
			last_song, last_artist = song, artist
			time.Sleep(120 * time.Second)
		} else {
			fmt.Printf("Nothing new yet...\n")
			time.Sleep(30 * time.Second)
		}
	}
}
