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
  accountUrl := fmt.Sprintf(`https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s`, acc.Name, acc.Tag, riotKey)
  resp, err := http.Get(accountUrl)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  //Closes the response body after the function ends, defer is a 'late' return it executes after the function ends
  defer resp.Body.Close()


  //Decodes the response body into the account struct

  err = json.NewDecoder(resp.Body).Decode(&acc)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(acc)
  return acc.Puuid

}

//GetsMatches takes an account and returns the id of the last 20 matches of the account 
func GetMatches(acc *types.Account) ([]string) {

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
  return acc.Matches

}
