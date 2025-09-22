package main

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "time"
    "strconv"

	"github.com/MohammedMogeab/rssagger/internal/database"
	"github.com/google/uuid"
)

func (apicfing *apiconfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// send an empty JSON object {} with status 200

   type request struct{
	Name string `json:"name"`
   }

    decoder:=  json.NewDecoder(r.Body)
	params :=request{}
	err:=decoder.Decode(&params)
	if err != nil {
		respondWithError(w,http.StatusBadRequest,"invalid request payload")
		return
	}
	user, err :=apicfing.db.CreateUser(r.Context(),database.CreateUserParams{
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
	})
	if err != nil {
		respondWithError(w,http.StatusInternalServerError,"failed to create user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}


func (apicfing *apiconfig) HandlerGetUser(w http.ResponseWriter, r *http.Request,user database.User) {
	
respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apicfing *apiconfig) HandlerGetPostForUser(w http.ResponseWriter, r *http.Request,user database.User) {
     // Optional ?limit=n query param; defaults to previous behavior (4)
     limit := int32(4)
     if v := r.URL.Query().Get("limit"); v != "" {
        if n, err := strconv.Atoi(v); err == nil {
            if n < 1 { n = 1 }
            if n > 100 { n = 100 }
            limit = int32(n)
        }
     }

     posts,err:= apicfing.db.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
        UserID: user.ID,
        Limit:  limit,
      })
      if err != nil {
        respondWithError(w,http.StatusInternalServerError,"failed to get posts for user")
        return
      }

respondWithJSON(w, http.StatusOK, databasePosttoPostArray(posts))
}
