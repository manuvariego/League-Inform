package analyze

import (
	"fmt"
	"leagueinform/internal/riot"
	"leagueinform/internal/types"
	"sync"
	"sync/atomic"
)

func AnalyzeMatches(acc *types.Account, matches []string) uint64 {
	var wg sync.WaitGroup
	// var matchesWon int
	// var mu sync.Mutex
	var epica atomic.Uint64

	for i := 0; i < len(matches); i++ {
		wg.Add(1)
		go func(matchID string) {
			defer wg.Done()
			match := riot.GetMatchInfo(acc, matchID)
			for _, participant := range match.Info.Participants {
				if participant.Puuid == acc.Puuid {
					if participant.Win {
						epica.Add(1)
						// mu.Lock()
						// matchesWon++
						// mu.Unlock()
						fmt.Println("This is the match", match.MatchId)
						fmt.Println("You won this match")
					} else {
						fmt.Println("This is the match", match.MatchId)
						fmt.Println("You lost this match")
					}
				}
			}
		}(matches[i])
	}

	wg.Wait()
	result := epica.Load()
	// return matchesWon
	return result
}
