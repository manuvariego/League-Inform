package types


type Account struct {
  Name string `json:"name"`  
  Tag  string `json:"tag"`
  Puuid string `json:"puuid"`
  Matches []string `json:"matches"`
}

// type Match struct {
//   MatchId string `json: "matchId"`
// }
