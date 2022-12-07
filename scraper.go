package main

import (
	"github.com/mmcdole/gofeed"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type FeedItem struct {
	Title string
	URL   string
}

func readFile(fname string) string {
	databyte, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(databyte)
}

func ParseRSS() {
	blogList := [1]string{"https://pwnlab.me/feed/"}
	fp := gofeed.NewParser()
	fp.Client = &http.Client{Timeout: time.Second * 5}

	feed_items := make([]FeedItem, 1)

	for true {
		for k := 0; k < len(blogList); k++ {
			feed, err := fp.ParseURL(blogList[k])
			if err == nil {
				l.Printf("[INFO] RSS Parse %s icin basladi", blogList[k])
				items := feed.Items
				for i := 0; i < len(items); i++ {
					if strings.Contains(readFile("feed_item.list"), items[i].Link) {
						l.Printf("[WARN] FEddItem zaten olusturuldu. Link : %s", items[i].Link)
					} else {
						feedItem := FeedItem{Title: items[i].Title, URL: items[i].Link}
						feed_items = append(feed_items, feedItem)
						l.Printf("[INFO] FeedItem olusturuldu. Title : %s, URL: %s", feedItem.Title, feedItem.URL)
						file, err := os.OpenFile("feed_item.list", os.O_APPEND|os.O_WRONLY, 0644)
						if err != nil {
							panic(err)
						}
						defer file.Close()
						if _, err := file.WriteString(items[i].Link + "\n"); err != nil {
							l.Fatal(err)
						}
						msg := "Yeni bir gonderi paylasildi : ***" + items[i].Title + "\n" + items[i].Link
						Dg.ChannelMessageSend(botChID, msg)
					}
				}
			} else {
				feed_items = make([]FeedItem, 1)
				l.Printf("[ERR] FeedItem olusturulmadi! Link veya title bos. Url : %s", blogList[k])
			}
		}
	}
}
