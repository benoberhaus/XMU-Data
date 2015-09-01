package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
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

func isJunk(song string, artist string, junk []string) bool {
	for i, _ := range junk {
		if strings.Contains(song, junk[i]) || strings.Contains(artist, junk[i]) {
			return true
		}
	}
	return false
}
