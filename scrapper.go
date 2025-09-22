package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/MohammedMogeab/rssagger/internal/database"
)



func startscrapper(db *database.Queries,concurrecy int ,intervalSeconds time.Duration) {
log.Printf("Starting scrapper with concurrency %d and interval %d seconds",concurrecy,intervalSeconds)
	ticker := time.NewTicker(intervalSeconds)
	for ; ;<-ticker.C {
       feed,err:= db.GetNextFeedForFetch(context.Background(),int32(concurrecy))
	   if err != nil {
		log.Printf("Error getting next feed for fetch: %v",err)
		continue
	   }
	   wg:=&sync.WaitGroup{}
	  
	   for _,fed:=range feed {
          wg.Add(1)
          go scrapeFeed(wg,fed,db)
	   }

	   wg.Wait()


	}





}

func parsePubDate(s string) (time.Time, error) {
	layouts := []string{
		time.RFC1123,    // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC1123Z,   // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC3339,    // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,
		time.ANSIC,      // some feeds do weird things
	}
	var lastErr error
	for _, l := range layouts {
		if t, err := time.Parse(l, s); err == nil {
			return t, nil
		} else {
			lastErr = err
		}
	}
	return time.Time{}, lastErr
}

func scrapeFeed(wg *sync.WaitGroup, feed database.Feed, q *database.Queries) {
	defer wg.Done()

	// 1) Fetch/parse RSS
	newFeed, err := rssToUrl(feed.Url) // your function returning parsed RSS
	if err != nil {
		log.Printf("Error fetching rss feed from url %s: %v", feed.Url, err)
		return
	}

	// 2) Insert posts
	for _, item := range newFeed.Channel.Item {
		// nullable strings
		title := sql.NullString{String: item.Title,        Valid: item.Title != ""}
		desc  := sql.NullString{String: item.Description,  Valid: item.Description != ""}

		// nullable time
		var publishedAt sql.NullTime
		if item.PubDate != "" {
			if t, err := parsePubDate(item.PubDate); err == nil {
				publishedAt = sql.NullTime{Time: t, Valid: true}
			} else {
				log.Printf("Warning: could not parse pubDate %q for feed %s: %v", item.PubDate, feed.Name, err)
			}
		}

	  _,err = q.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Name:        item.Title,   // choose a meaningful value for "name"
			Url:         item.Link,
			Title:       title,
			Description: desc,         // ok because column is nullable now
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			// Skip duplicates on url (unique index)
			if strings.Contains(err.Error(),"duplicate key") {
			 continue
			}
		}

		log.Printf("Feed Item: %s - %s", item.Title, item.Link)
	}

	// 3) Mark feed as fetched AFTER processing
	if _, err := q.MarkFeedAsFetched(context.Background(), feed.ID); err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
	}


	log.Printf("Finished scraping feed %s", feed.Name)
}