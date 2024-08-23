package main

import (
	"encoding/json"
	"log"
	"net/http"
  "leagueinform/api"
)

func main() {
  api.RunServer()

}

// func handleUserMatches(w http.ResponseWriter, r *http.Request) {
// 	//ver lo de var y type
//   var account struct {
// 		Name  string `json:"name"`
// 		Tag   string `json:"tag"`
// 		Puuid string `json:"puuid"`
// 	}
// 	//Estas cosas vendran de una validacion previa, auth token
//   fmt.Println("Ingrese Nombre y despues tag")
// 	fmt.Scanln(&account.Name)
//   fmt.Scanln(&account.Tag)
// 	//
//
// 	fmt.Println("ALL GOOD ON THAT")
// 	riotKey := os.Getenv("RIOT_KEY")
// 	accountUrl := fmt.Sprintf(`https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s`, account.Name, account.Tag, riotKey)
//
// 	resp, err := http.Get(accountUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//decoded json from the get function to the riot api
// 	err = json.NewDecoder(resp.Body).Decode(&account)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	//set headers to expect json format
// 	w.Header().Set("Content-Type", "application/json")
// 	//encoded json with writer, and account struct
// 	json.NewEncoder(w).Encode(account)
//
// }
func handleHealth(w http.ResponseWriter, r *http.Request) {
	var nashe struct {
		Name     string `json:"name"`
		Lastname string `json:"lastname"`
	}
	err := json.NewDecoder(r.Body).Decode(&nashe)
	if err != nil {
		log.Fatalf("This is the error %d", err)
	}
	println("All good, here is the name you passed", nashe.Lastname)

}
