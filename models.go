package main

import (
	"database/sql"
	"time"

	"github.com/MohammedMogeab/rssagger/internal/database"

	"github.com/google/uuid"
)

type User struct{
	ID        uuid.UUID `json:"id"`
	Name      string `json:"name"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	ApiKey  string `json:"api_key,omitempty"`
}

type Feed struct{
	ID        uuid.UUID `json:"id"`
	Name      string `json:"name"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}
func databaseUserToUser(dbUser database.User) User {
 return User{
	ID: dbUser.ID, 
	Name: dbUser.Name,
	CreatedAt: dbUser.CreatedAt,
	UpdatedAt: dbUser.UpdatedAt,
	ApiKey: dbUser.ApiKey,

 }

 
}

type Feedsfollow struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databaseFeedtoFeed(dbfeed database.Feed) Feed {
 return Feed{
	ID: dbfeed.ID, 
	Name: dbfeed.Name,
	CreatedAt: dbfeed.CreatedAt,
	UpdatedAt: dbfeed.UpdatedAt,
	Url: dbfeed.Url,
	UserID: dbfeed.UserID,

 }





}


func databaseFeedstoFeeds(dbfeed []database.Feed) []Feed {
  feeds:= []Feed{}
  for _,dbfeed :=range dbfeed{
	feeds = append(feeds, databaseFeedtoFeed(dbfeed))
  }
  return  feeds;




}

func databaseFeedFollowtoFeed(dbfeed database.Feedsfollow) Feedsfollow {
 return Feedsfollow{
	ID: dbfeed.ID, 
	CreatedAt: dbfeed.CreatedAt,
	UpdatedAt: dbfeed.UpdatedAt,
	FeedID: dbfeed.FeedID,
	UserID: dbfeed.UserID,

 }
}

func databaseFeedFollowstoFeedsFollow(dbfeedfollow []database.Feedsfollow) []Feedsfollow {
  feedsfollow:= []Feedsfollow{}
  for _,dbfeedfollow :=range dbfeedfollow{
	feedsfollow = append(feedsfollow, databaseFeedFollowtoFeed(dbfeedfollow))
  }
  return  feedsfollow;




}


type Post struct{
	ID        uuid.UUID `json:"id"`
	Name      string `json:"name"`
	CreatedAt *string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
	Url string `json:"url"`
	FeedID uuid.UUID `json:"feed_id"`
	Description *string `json:"description"`
	PublishedAt sql.NullTime `json:"published_at"`
	Title *string `json:"title"`
}

func databasePosttoPost(dbpost database.Post) Post {
 return Post{
	ID: dbpost.ID,
	Name: dbpost.Name,
	CreatedAt: nultimeptr(dbpost.CreatedAt),
	UpdatedAt: nultimeptr(dbpost.UpdatedAt),
	Url: dbpost.Url,
	FeedID: dbpost.FeedID,	
	Description: &dbpost.Description.String,
	PublishedAt: dbpost.PublishedAt,
	Title: &dbpost.Title.String,
 }



}

func nultimeptr(t time.Time) *string {
	s := t.Format(time.RFC3339)
	if s == "" {
		return nil
	}
	return &s

}

func databasePosttoPostArray(dbpost []database.Post) []Post {
  posts:= []Post{}
  for _,dbpost :=range dbpost{
	posts = append(posts, databasePosttoPost(dbpost))
  }
  return  posts;		


}