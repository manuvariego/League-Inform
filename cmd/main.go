package main

import (
	"leagueinform/api"
)

func main() {

	ws := api.ConnectToDiscord()
	ws.Reader(ws.Conn)

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }

	// dstoken := os.Getenv("DISCORD_KEY")

	// Create a new Discord session using the bot token.
	// sess, err := discordgo.New("Bot " + dstoken)
	// if err != nil {
	// 	fmt.Println("Error creating Discord session: ", err)
	// 	return
	// }

	//Calls 'DiscordBot' from the discord package
	// discord.DiscordBot(sess)

	//Temp code to keep the program running
	// fmt.Scanln()

}
