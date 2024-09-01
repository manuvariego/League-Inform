package analyze


import (
  "leagueinform/internal/types"
  "leagueinform/internal/riot"
  "fmt"


)


//takes some matches and analyzes them
func AnalyzeMatches(acc *types.Account, matches []string) int {
  var matchesWon int
  for i:=0; i < len(matches); i++ { 
    match := riot.GetMatchInfo(acc, matches[i])
    for _, participant:= range match.Info.Participants {
      if participant.Puuid== acc.Puuid {
        switch participant.Win  {
        case true:
          matchesWon += 1
          fmt.Println("This is the match", match.MatchId)
          fmt.Println("You won this match")
        case false: 
          fmt.Println("This is the match", match.MatchId)
          fmt.Println("You lost this match") 
        }
      }
    }

  }
  return matchesWon
}



















