package tests

import (
  "testing"
  "leagueinform/internal/types"
  "sync"
  "fmt"
  "github.com/joho/godotenv"
  "leagueinform/internal/riot"

)

func TestGetId(t *testing.T) {
  err := godotenv.Load("../.env")
  if err != nil {
    t.Error("Error loading .env file")
  }

	acc := types.Account{Name: "Krazie", Tag: "LAS"}
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		puuid:= riot.GetId(&acc) // Assuming GetId now returns an error
    fmt.Println(puuid)
		acc.Puuid = puuid
	}()

	wg.Wait()

	if acc.Puuid == "" {
		t.Error("Expected puuid to be filled, but it was empty")
	}
}


// func TestGetId(t *testing.T) {
//   acc := types.Account{Name: "Krazie", Tag: "LAS"}
//   var wg sync.WaitGroup
//   wg.Add(2)
//   go func(){
//     defer wg.Done()
//     puuid:=GetId(&acc)
//     acc.Puuid = puuid
//
//   }()
//   wg.Wait()
//
//   if acc.Puuid == "" {
//    t.Error("Expected puuid to be filled")
//   }
// }
