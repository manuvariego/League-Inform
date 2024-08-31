package riot 

import (
  "os"
  "log"
  "fmt"
  "net/http"
  "leagueinform/internal/types"
  "encoding/json"
)

type RiotApi struct{
  key string
  
};

func NewRiotApi() *RiotApi {
  return &RiotApi{
    key: os.Getenv("RIOT_KEY"),
  }
}

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

//GetsMatchestakes an account and returns the id of the last 20 matches of the account 
func GetMatches(acc *types.Account){
  // var match types.Match 

  riotKey := os.Getenv("RIOT_KEY")

  req, err := http.NewRequest("GET",fmt.Sprintf(`https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20`, acc.Puuid), nil) 
  if err != nil {
    log.Fatal(err)
  }

  sliceMatches := make([]type.Match, 20)

  // acc.Matches := make([]string, 20)
  req.Header.Set("X-Riot-Token", riotKey)
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }


  //This get requests returns a string of match ID's they should go into the account match structure. 
  err = json.NewDecoder(res.Body).Decode(&sliceMatches)
  if err != nil{
    log.Fatal(err)
  }

  acc.Matches = sliceMatches

}

func GetMatchInfo(acc *types.Account) {

  riotKey := os.Getenv("RIOT_KEY")

  //Should probably use a go routine for this, iterating a get request should probably not be responsability of this function 

  for i := 0; i < len(acc.Matches); i++ {
    var match types.Match
    fmt.Println(len(acc.Matches))
    fmt.Println(i)

    req, err := http.NewRequest("GET",fmt.Sprintf(`https://americas.api.riotgames.com/lol/match/v5/matches/%s`, acc.Matches[i]), nil) 
    if err != nil {
      log.Fatal(err)
    }

    req.Header.Set("X-Riot-Token", riotKey)
    res, err := http.DefaultClient.Do(req)
    if err != nil {
      log.Fatal(err)
    }

    err = json.NewDecoder(res.Body).Decode(&match)
    if err != nil {
      log.Fatal(err)
    }

    for i, participant:= range match.Info.Participants {

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

  }
}
