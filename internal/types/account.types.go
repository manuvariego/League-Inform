package types


type Account struct {
  Name string `json:"name"`  
  Tag  string `json:"tag"`
  Puuid string `json:"puuid"`
  Matches []Match `json:"matches"`
}

//Probably not necessary, commented for now, (why do I need this?)
type Match struct {
  MatchId string `json:"matchId"`
  GameResult string `json:"endOfGameResult"`
  GameMode string `json:"gameMode"`
  GameStartTimestamp string `json:"gameStartTimestamp"`
}
