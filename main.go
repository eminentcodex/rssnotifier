package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ctcpip/notifize"
)

type (
	RSSNotifier struct {
		RSSPool []iRSSFeed
	}
)

func main() {
	var pool RSSNotifier

	ec := EnglishClub{}
	ec.SetUrl("https://www.englishclub.com/ref/idiom-of-the-day.xml")

	pool.RSSPool = append(pool.RSSPool, iRSSFeed(&ec))
	iconPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	ticker := time.NewTicker(time.Minute * 30)

	for range ticker.C {
		for _, rss := range pool.RSSPool {
			if err := rss.GetFeed(); err == nil {
				notifize.Display("Title : "+rss.GetTitle(), rss.GetFeedLink(), false, iconPath+"/icon.png")
			}
		}
	}

}
