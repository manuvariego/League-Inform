package main

import (
  "fmt"
  "github.com/joho/godotenv"
  "os"
  "github.com/bwmarrin/discordgo"
  "leagueinform/internal/discord"
)


func main() {

  err := godotenv.Load("../.env")
  if err != nil {
    fmt.Println("Error loading .env file")
  }

  dstoken := os.Getenv("DISCORD_KEY")

  // Create a new Discord session using the bot token.
  sess, err := discordgo.New("Bot " + dstoken)
  if err != nil {
    fmt.Println("Error creating Discord session: ", err)
    return
  }

  //Calls 'DiscordBot' from the discord package
  discord.DiscordBot(sess)
  fmt.Scanln()



}

