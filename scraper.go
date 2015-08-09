package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func scrape(url string) string {
	doc, _ := goquery.NewDocument(url)
	var ret string
	doc.Find("td").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if strings.HasPrefix(s.Text(), "\n35") {
			ret = s.Text()
			return false
		}
		return true
	})
	return ret
}

func parse(dump string) (string, string) {
	splits := strings.Split(dump, "\n")
	return splits[len(splits)-1], strings.Replace(splits[len(splits)-2], " - ", "", 1)
}

func main() {
	url := "http://www.dogstarradio.com/now_playing.php"
	song, artist := parse(scrape(url))
	fmt.Printf("%s - %s", artist, song)
}
