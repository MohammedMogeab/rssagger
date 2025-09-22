package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiToken(r http.Header)(string ,error){

val:=r.Get("Authorization")
   if val ==""{
	  return "",errors.New("missing api key")   
    }
 vals:= strings.Split(val, " ")
 if len(vals) !=2 || vals[0] !="Bearer"{
	return "",errors.New("invalid authorization header")
 }



return  vals[1],nil


}