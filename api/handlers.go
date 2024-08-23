package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
  "leagueinform/models"
)

func HandleUserMatches(w http.ResponseWriter, r *http.Request) {
	//ver lo de var y type
	//  var account struct {
	// 	Name  string `json:"name"`
	// 	Tag   string `json:"tag"`
	// 	Puuid string `json:"puuid"`
	// }
  account1 := &models.Account {}
  err := json.NewDecoder(r.Body).Decode(&account1)
	if err != nil  {
		log.Fatal(err)
	}

	fmt.Println("ALL GOOD ON THAT")
	riotKey := os.Getenv("RIOT_KEY")
	accountUrl := fmt.Sprintf(`https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s`, account1.Name, account1.Tag, riotKey)

	resp, err := http.Get(accountUrl)
	if err != nil {
    fmt.Println("Puuid not found or error")
		log.Fatal(err)
	}
	//decoded json from the get function to the riot api
	err = json.NewDecoder(resp.Body).Decode(&account1)
	if (err != nil) || (account1.Puuid == "")  {
		log.Fatal(err)
	}

	//set headers to expect json format
	w.Header().Set("Content-Type", "application/json")
	//encoded json with writer, and account struct
	json.NewEncoder(w).Encode(account1)

}
