package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter,status int ,message string) {
  if status > 499{
	log.Println("Response with error 5xx",message)

  }

  type reponseErr struct{
	Err string `json:"error"`
  }
  respondWithJSON(w,status,reponseErr{Err: message})


}

func respondWithJSON(w http.ResponseWriter,status int ,data interface{}) {

jsonData,err :=json.Marshal(data)
if err != nil {
	log.Printf("Error marshalling JSON: %v , data: %v", err, data)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("HTTP 500: Internal Server Error"))
	return
}

w.Header().Set("Content-Type","application/json")
w.WriteHeader(status)
log.Printf("Responding with JSON: %v",string(jsonData))
w.Write(jsonData)

}