package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func scrape(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}
	var ret string
	doc.Find("td").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if strings.HasPrefix(s.Text(), "\n35") {
			ret = s.Text()
			return false
		}
		return true
	})
	return ret, nil
}

func parse(dump string) (string, string) {
	splits := strings.Split(dump, "\n")
	return splits[len(splits)-1], strings.Replace(splits[len(splits)-2], " - ", "", 1)
}

func getCurrentSong() (string, string, error) {
	url := "http://www.dogstarradio.com/now_playing.php"
	dump, err := scrape(url)
	song, artist := parse(dump)
	if err != nil {
		return "", "", err
	}
	return song, artist, nil
}
