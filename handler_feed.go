package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/MohammedMogeab/rssagger/internal/database"
	"github.com/google/uuid"
	"log"
	"github.com/go-chi/chi/v5"
)

func (apicfing *apiconfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request,user database.User) {
	// send an empty JSON object {} with status 200

   type request struct{
	Name string `json:"name"`
	Url string `json:"url"`
   }

    decoder:=  json.NewDecoder(r.Body)
	params :=request{}
	err:=decoder.Decode(&params)
	if err != nil {
		respondWithError(w,http.StatusBadRequest,"invalid request payload")
		return
	}
	feed, err :=apicfing.db.CreateFeed(r.Context(),database.CreateFeedParams{
		ID:uuid.New(),
		Name: params.Name,
		CreatedAt: sql.NullTime{
			Time:time.Now(),
			Valid:true,
		},
		UpdatedAt: sql.NullTime{
			Time:time.Now(),
			Valid:true,
		},
		Url: params.Url,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w,http.StatusInternalServerError,"failed to create feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedtoFeed(feed))
}


func (apicfing *apiconfig) HandlerGetfeed(w http.ResponseWriter, r *http.Request) {
	feeds, err :=apicfing.db.GetFeeds(r.Context())
		if err != nil {
		respondWithError(w,http.StatusInternalServerError,"failed to get feed")
		return
	}

 

respondWithJSON(w, http.StatusOK,databaseFeedstoFeeds(feeds))

}



func (apicfing *apiconfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request,user database.User) {
	// send an empty JSON object {} with status 200

   type request struct{
	Feed_id string `json:"feed_id"`
   }

    decoder:=  json.NewDecoder(r.Body)
	params :=request{}
	err:=decoder.Decode(&params)
	if err != nil {
		respondWithError(w,http.StatusBadRequest,"invalid request payload")
		return
	}
	Feedsfollow, err :=apicfing.db.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID:uuid.New(),
		
		CreatedAt: sql.NullTime{
			Time:time.Now(),
			Valid:true,
		},
		UpdatedAt: sql.NullTime{
			Time:time.Now(),
			Valid:true,
		},
		UserID: user.ID,
		FeedID: uuid.MustParse(params.Feed_id),
	})
	if err != nil {
		respondWithError(w,http.StatusInternalServerError,"failed to follow feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowtoFeed(Feedsfollow))
}



func (apicfing *apiconfig) HandlerGetFeedFollow(w http.ResponseWriter, r *http.Request,user database.User) {
	// send an empty JSON object {} with status 200

	Feedsfollow, err :=apicfing.db.GetFeedFollowsByUserID(r.Context(),user.ID)
	
	if err != nil {
		respondWithError(w,http.StatusInternalServerError,"failed to follow feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowstoFeedsFollow(Feedsfollow))
}


func (apicfing *apiconfig) DeleteFeed(w http.ResponseWriter, r *http.Request,user database.User) {
	// send an empty JSON object {} with status 200

log.Println("path:", r.URL.Path)
	feedId := chi.URLParam(r,"feedfollowId")

    log.Println("feed id:",feedId)
	feed,err := uuid.Parse(feedId)
	if err != nil {
		respondWithError(w,http.StatusBadRequest,"invalid feed id")
		return
	}
	err =apicfing.db.DeleteFeedFollowByID(r.Context(),database.DeleteFeedFollowByIDParams{
		ID:feed,
		UserID:user.ID,
	})
	if err != nil {
		respondWithError(w,http.StatusInternalServerError,"failed to delete feed")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result":"feed deleted"})


}
