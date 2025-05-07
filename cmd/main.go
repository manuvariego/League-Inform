package main

import (
	"fmt"
	"leagueinform/api"
	"leagueinform/discord"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := StartDatabase()
	//Starts the api server
	go api.RunServer(db)

	//Starts the websocket connection!
	ws := discord.ConnectToDiscord()
	ws.Reader()
	ws.Identify()

}
