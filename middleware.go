package main

import (
	"github.com/MohammedMogeab/rssagger/internal/database"
	"net/http"
	"database/sql"

	"github.com/MohammedMogeab/rssagger/internal/auth"
)
type AuthHandler func(http.ResponseWriter,*http.Request,  database.User)


func (apicfing *apiconfig) MiddlewareAuth(next AuthHandler) http.HandlerFunc {
	return func(r http.ResponseWriter,rq *http.Request) {
    user,err := auth.GetApiToken(rq.Header)
   if err != nil {
	respondWithError(r,http.StatusUnauthorized,"missing or invalid api key")
	return   
   }
 
  userc,err:=apicfing.db.GetUserByAPIKey(rq.Context(),user)
if err != nil {
	if err == sql.ErrNoRows{
		respondWithError(r,http.StatusUnauthorized,"invalid api key")
		return   
	}
	respondWithError(r,http.StatusInternalServerError,"failed to get user")
	return   
}

    next(r,rq,userc)


	}
  }








