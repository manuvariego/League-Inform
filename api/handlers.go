package api

import (
	"net/http"
  "io"
  "fmt"
  "os"
  "log"
)

// func HandleUserMatches(w http.ResponseWriter, r *http.Request) {
//   services.GetRiotPuuid(w, r)
//
//
//
// }


func GetRiotPuuid(w http.ResponseWriter, r *http.Request){
  var accountName, accountTag string


  riotKey := os.Getenv("RIOT_KEY")

  accountName = "Krazie"
  accountTag = "LAS"

  accountUrl := fmt.Sprintf(`https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s`, accountName, accountTag, riotKey)

	resp, err := http.Get(accountUrl)
  if err != nil{
    log.Fatal(err)
  }

  w.Header().Set("Content-Type", "application/json")

  io.Copy(w, resp.Body)



}



