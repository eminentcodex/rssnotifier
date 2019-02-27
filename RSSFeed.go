package main

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"golang.org/x/net/html/charset"
)

type iRSSFeed interface {
	GetFeedLink() string
	GetTitle() string
	GetDecription() string
	SetUrl(url string)
	setFeedLink(feedLink string)
	setTitle(title string)
	setDecription(desc string)
	GetFeed() (err error)
	Call(requestBody []byte) (xml []byte, err error)
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
}

type Channel struct {
	XMLName     xml.Name `xml:channel`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Language    string   `xml:"language"`
	Copyright   string   `xml:"copyright"`
	Link        string   `xml:"link"`
	Item        Item     `xml:"item"`
}

type EnglishClubFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type EnglishClub struct {
	HttpCall
	FeedLink    string
	Description string
	Title       string
}

func (ec *EnglishClub) GetFeedLink() string {
	return ec.FeedLink
}

func (ec *EnglishClub) GetTitle() string {
	return ec.Title
}

func (ec *EnglishClub) GetDecription() string {
	return ec.Description
}

func (ec *EnglishClub) setFeedLink(feedLink string) {
	ec.FeedLink = feedLink
}

func (ec *EnglishClub) SetUrl(url string) {
	ec.Url = url
}

func (ec *EnglishClub) setTitle(title string) {
	ec.Title = title
}

func (ec *EnglishClub) setDecription(desc string) {
	ec.Description = desc
}

func (ec *EnglishClub) GetFeed() (err error) {
	var (
		data []byte
		feed EnglishClubFeed
	)

	if data, err = ec.Call(nil); err != nil {
		return
	}

	// parse xml
	decoder := xml.NewDecoder(bytes.NewBuffer(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err = decoder.Decode(&feed); err != nil {
		fmt.Println(err)
		return
	}

	ec.setTitle(feed.Channel.Item.Title)
	ec.setDecription(feed.Channel.Item.Description)
	ec.setFeedLink(feed.Channel.Item.Link)

	return

}
