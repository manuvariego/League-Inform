package api

import (
	"net/http"
  // "io"
  "fmt"
  "os"
  "log"
  "leagueinform/types"
  "encoding/json"
)

// func HandleUserMatches(w http.ResponseWriter, r *http.Request) {
//   services.GetRiotPuuid(w, r)
//
//
//
// }


func GetRiotPuuid(w http.ResponseWriter, r *http.Request){

  account1:= &types.Account{}

  riotKey := os.Getenv("RIOT_KEY")

  account1.Name = "Krazie"
  account1.Tag = "LAS"

  accountUrl := fmt.Sprintf(`https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s`, account1.Name, account1.Tag, riotKey)

	resp, err := http.Get(accountUrl)
  if err != nil{
    log.Fatal(err)
  }

  json.NewDecoder(resp.Body).Decode(&account1)

  // w.Header().Set("Content-Type", "application/json")
  //
  // io.Copy(w, resp.Body)
  fmt.Println(account1)

  GetRiotMatchesByPuuid(w, r, account1)
  fmt.Println("test")

}


func GetRiotMatchesByPuuid(w http.ResponseWriter, r *http.Request, acc *types.Account) {

  riotKey := os.Getenv("RIOT_KEY")

  accountUrl := fmt.Sprintf(`https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20&api_key=%s`,acc.Puuid , riotKey)

  resp, err := http.Get(accountUrl)
  if err != nil{
    log.Fatal(err)
  }

  err = json.NewDecoder(resp.Body).Decode(&acc.Matches)
  if err != nil{
    log.Fatal(err)
  }

  fmt.Println(acc.Matches)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(acc.Matches)
}

