package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type RSSfeed struct {
	Channel struct {
		Title       string        `xml:"title"`
		Link        string        `xml:"link"`
		Description string        `xml:"description"`
		Language    string        `xml:"language"`
		Item        []RSSfeedItem `xml:"item"`
	} `xml:"channel"`
}
type RSSfeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func rssToUrl(rss string) (RSSfeed, error) {
	httpClient :=http.Client{
		Timeout: 10 * 1e9, // 10 seconds
	}
	resp, err := httpClient.Get(rss)
	if err != nil {
		return RSSfeed{}, err
	}
	defer resp.Body.Close()
	data,err:= io.ReadAll(resp.Body)
	if err != nil {
		return RSSfeed{}, err
	}
    rssfeeds:=RSSfeed{}
	xml.Unmarshal(data, &rssfeeds)
	return rssfeeds, nil

}