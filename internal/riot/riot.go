package riot 


import (
  "os"
  "log"
  "fmt"
  "net/http"
  "leagueinform/internal/types"
  "encoding/json"

)

//GetsID takes an account and returns the puuid of the account after accesing the riot api
func GetId(acc *types.Account) string {

  riotKey := os.Getenv("RIOT_KEY")

  req, err := http.NewRequest("GET",fmt.Sprintf(`https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s`, acc.Name, acc.Tag), nil) 
  if err != nil {
    log.Fatal(err)
  }

  req.Header.Set("X-Riot-Token", riotKey)
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  //Closes the response body after the function ends, defer is a 'late' return it executes after the function ends
  defer res.Body.Close()

  //Decodes the response body into the account struct

  err = json.NewDecoder(res.Body).Decode(&acc)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(acc)
  return acc.Puuid

}

//GetsMatches takes an account and returns the id of the last 20 matches of the account 
func GetMatches(acc *types.Account) ([]types.Match) {
  var matchIds []string

  riotKey := os.Getenv("RIOT_KEY")

  req, err := http.NewRequest("GET",fmt.Sprintf(`https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20`, acc.Puuid), nil) 
  req.Header.Set("X-Riot-Token", riotKey)
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }


  //This get requests returns a string of match ID's they should go into the account match structure. 
  err = json.NewDecoder(res.Body).Decode(&matchIds)
  if err != nil{
    log.Fatal(err)
  }

  //Makes a slice of matches the length of matchIds (normally 20)) 
  acc.Matches = make([]types.Match, len(matchIds))

  for i := 0; i<len(matchIds); i++{
    acc.Matches[i] = types.Match{
      MatchId: matchIds[i],
    }
  }
  return acc.Matches
}
func GetMatchInfo(acc *types.Account) {
  // i = 0

  riotKey := os.Getenv("RIOT_KEY")

  for i := 0; i < len(acc.Matches); i++ {

    req, err := http.NewRequest("GET",fmt.Sprintf(`https://americas.api.riotgames.com/lol/match/v5/matches/%s`, acc.Matches[i].MatchId), nil) 
    if err != nil {
      log.Fatal(err)
    }

    req.Header.Set("X-Riot-Token", riotKey)
    res, err := http.DefaultClient.Do(req)
    if err != nil {
      log.Fatal(err)
    }

    err = json.NewDecoder(res.Body).Decode(&acc.Matches[i])
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(acc.Matches[i])

  }

}
