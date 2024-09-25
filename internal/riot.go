package riot

import (
	"encoding/json"
	"fmt"
	"leagueinform/internal/types"
	"log"
	"net/http"
	"os"
)

type RiotApi struct {
	key string
}

func NewRiotApi() *RiotApi {
	return &RiotApi{
		key: os.Getenv("RIOT_KEY"),
	}
}

// GetsID takes an account and returns the puuid of the account after accesing the riot api
func GetId(acc *types.Account) string {

	riotKey := os.Getenv("RIOT_KEY")

	req, err := http.NewRequest("GET", fmt.Sprintf(`https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s`, acc.Name, acc.Tag), nil)
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

// GetsMatchestakes an account and returns the id of the last 20 matches of the account
func GetMatchesById(acc *types.Account) []string {
	var matches []string

	riotKey := os.Getenv("RIOT_KEY")

	req, err := http.NewRequest("GET", fmt.Sprintf(`https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20`, acc.Puuid), nil)
	if err != nil {
		log.Fatal(err)
	}

	// acc.Matches := make([]string, 20)
	req.Header.Set("X-Riot-Token", riotKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//This get requests returns a string of match ID's they should go into the account match structure.
	err = json.NewDecoder(res.Body).Decode(&matches)
	if err != nil {
		log.Fatal(err)
	}

	return matches

}

func GetMatchInfo(acc *types.Account, match string) types.Match {

	riotKey := os.Getenv("RIOT_KEY")

	//Should probably use a go routine for this, iterating a get request should probably not be responsability of this function

	var matchType types.Match

	req, err := http.NewRequest("GET", fmt.Sprintf(`https://americas.api.riotgames.com/lol/match/v5/matches/%s`, match), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Riot-Token", riotKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(res.Body).Decode(&matchType)
	if err != nil {
		log.Fatal(err)
	}
	matchType.MatchId = match

	return matchType
}
