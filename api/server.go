package api

import (
	"github.com/joho/godotenv"
	"log"
  "fmt"
	"net/http"
  "context"
  "time"
)



func RunServer() {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
	err := godotenv.Load()
  mux := http.NewServeMux()
  //need a router to grab ID from the URL and pass it to the function
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/matches", GetRiotPuuid)
  fmt.Println(ctx)

  go func() {
    err = http.ListenAndServe("localhost:3000", mux)
    if err != nil {
      log.Fatal(err)
    }
  } () 

  <- ctx.Done()
}
