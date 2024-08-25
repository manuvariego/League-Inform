package discord

import (
  "fmt"
  "github.com/bwmarrin/discordgo"
  "log"
  "leagueinform/internal/types"
  "leagueinform/internal/riot"
)
//sess is a discord session initialized in main.go
func DiscordBot( sess *discordgo.Session) {

  //Create a new account 
  acc := &types.Account{}
  i := 0

  //Adds a handler to the bot
  sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
    fmt.Println(m.Content)
    if m.Author.ID == s.State.User.ID {
      return 
    }
    if m.Content == "matches please" {

      s.ChannelMessageSend(m.ChannelID, "Please enter your riot name and tag, with this format 'riot-name#tag'")

      //Hardcoded for now because I do not know how to wait for user input before continuing
      m.Content = "Krazie#LAS"

      //Divides the content of the message into the name and tag of the account 
      for i < len(m.Content) {
        if m.Content[i] == '#' {
          acc.Name = m.Content[:i]
          acc.Tag = m.Content[i+1:]
          break
        }
        i++
      }

      acc.Puuid = riot.GetId(acc)
      acc.Matches = riot.GetMatches(acc)

      //only printing the first match for now
      fmt.Println(acc.Matches[1])

      //!!Create function that takes the matches and sends them as a sum of strings (probably xd) 
      s.ChannelMessageSend(m.ChannelID, "Here are your matches: " + acc.Matches[1])
    }
  })



  sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
  err := sess.Open()
  if err != nil {
    log.Fatal(err)
  }
  defer sess.Close()

  fmt.Println("Bot is running")

}
