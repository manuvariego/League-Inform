package tests

import (
	"fmt"
	"leagueinform/internal/riot"
	"leagueinform/internal/types"
	"sync"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error("Error loading .env file")
	}
}

func TestGetId(t *testing.T) {
	acc := types.Account{Name: "Krazie", Tag: "LAS"}
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		puuid := riot.GetId(&acc) // Assuming GetId now returns an error
		fmt.Println(puuid)
		acc.Puuid = puuid
	}()

	wg.Wait()

	if acc.Puuid == "" {
		t.Error("Expected puuid to be filled, but it was empty")
	}
}

func TestGetMatches(t *testing.T) {
	acc := types.Account{Name: "Krazie", Tag: "LAS", Puuid: "QRAqaToSh_Ut7FaodO9DS_ZxrEQP6k4BpZ1MKDpvRWL9mC2wZ51OLgwUadyTWnLq50Qi6Vt-NjnZ4w"}
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		match := riot.GetMatchesById(&acc) // Assuming GetId now returns an error
		if match == nil {
			t.Error("Expected match to be filled, but it was empty")
		}
		fmt.Println(match)
	}()

	wg.Wait()

}
