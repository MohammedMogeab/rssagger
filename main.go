package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"

	"github.com/MohammedMogeab/rssagger/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiconfig struct{
    db *database.Queries
    dbConn *sql.DB
}



func main() {


	fmt.Println("Hello, World!")
	godotenv.Load(".env")
	
    portString:=os.Getenv("PORT")
    if(portString==""){
        log.Fatal("PORT not set in .env file")
    }
 
  

	dbUrl:=os.Getenv("DB_URL")
	if(dbUrl==""){
		log.Fatal("dbUrl not set in .env file")
	}


      conn,err:= sql.Open("postgres",dbUrl)
      if err != nil {
        log.Fatal("cannot connect to db:",err)

      }

      apiCfg:=apiconfig{
        db: database.New(conn),
        dbConn: conn,
      }
    

 
	fmt.Println("Server starting on port:",portString)
  r:=chi.NewRouter()


  r.Use(cors.Handler(cors.Options{
	AllowedHeaders: [] string{"*"},
	AllowedMethods: [] string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedOrigins: [] string{"http://localhost:8000","http://"},
	AllowCredentials: false,
	MaxAge: 300,
  }))

 routerv2:=chi.NewRouter()
 routerv2.Get("/healthz",apiCfg.HandlerHealthz)
 routerv2.Get("/error",HandlerError)
 routerv2.Post("/users",apiCfg.HandlerCreateUser)
 routerv2.Get("/users",apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
 routerv2.Post("/feeds",apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
 routerv2.Get("/feeds",apiCfg.HandlerGetfeed)
 routerv2.Post("/feeds/follow",apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
routerv2.Get("/feeds/follow",apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollow))
routerv2.Delete("/feeds/follow/{feedfollowId}",apiCfg.MiddlewareAuth(apiCfg.DeleteFeed))
routerv2.Get("/posts",apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostForUser))



r.Mount("/v1",routerv2)


 // Scraper configuration from environment
 // SCRAPER_CONCURRENCY: int (default 3)
 // SCRAPER_INTERVAL_SECONDS: int seconds (default 60)
 concurrency := 3
 if v := os.Getenv("SCRAPER_CONCURRENCY"); v != "" {
     if n, err := strconv.Atoi(v); err == nil && n > 0 {
         concurrency = n
     }
 }
 intervalSeconds := 60
 if v := os.Getenv("SCRAPER_INTERVAL_SECONDS"); v != "" {
     if n, err := strconv.Atoi(v); err == nil && n > 0 {
         intervalSeconds = n
     }
 }
 go startscrapper(apiCfg.db, concurrency, time.Duration(intervalSeconds)*time.Second)
  

 r.Get("/hello",func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
 })
 
 err = http.ListenAndServe(":"+portString,r)
 if err != nil {
	log.Fatal(err)
 }






 

} 
