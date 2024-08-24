package api

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)



func RunServer() {
	err := godotenv.Load()
  //need a router to grab ID from the URL and pass it to the function
	http.HandleFunc("/matches", GetRiotPuuid)
	err = http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
