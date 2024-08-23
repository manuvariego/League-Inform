package api

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)



func RunServer() {
	err := godotenv.Load()
  http.HandleFunc("/matches", HandleUserMatches)
	err = http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
