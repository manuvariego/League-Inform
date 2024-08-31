package analyze


import (
  "leagueinform/internal/types"


)


//takes some matches and analyzes them
func AnalyzeMatches(acc *types.Match) string {
      if participant.Puuid== acc.Puuid {
        switch participant.Win  {
        case true:
          fmt.Println("This is the match", acc.Matches[i])
          fmt.Println("You won this match")
        case false: 
          fmt.Println("This is the match", acc.Matches[i])
          fmt.Println("You lost this match") 
        }
      }
} 










