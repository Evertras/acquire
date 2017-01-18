package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	acquire "github.com/evertras/acquire/lib"
)

const (
	numPlayers = 5
	iterations = 100000
)

func runGames(c chan *acquire.Game) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for {
		players := make([]acquire.Player, numPlayers)[:0]
		for i := 0; i < numPlayers; i++ {
			players = append(players, acquire.NewPlayerRandom(r))
		}

		g := acquire.NewGame(r, players)

		phases := 0
		lastIndex := 0

		for g.State != nil {
			g.Advance()
			if g.CurrentPlayerIndex != lastIndex {
				lastIndex = g.CurrentPlayerIndex
			}
			phases++
			if phases > 10000 {
				for h := acquire.HotelFirst; h < acquire.HotelLast; h++ {
					fmt.Printf("%c: %d\n", acquire.GetHotelInitial(h), g.CurrentChainSizes[h])
				}
				g.Board.PrintBoard(os.Stdout)
				for i, p := range g.Players {
					fmt.Printf("Player %d: %v (%d)\n", i+1, p.CanPlayPiece(g), p.GetPieces())
				}
				fmt.Printf("Remaining pieces in bag: %d\n", len(g.PieceBag.Pieces))
				fmt.Printf("Remaining available chains: %d\n", len(g.AvailableChains))
				panic("Infinite loop")
			}
		}

		c <- g
	}
}

func main() {
	totals := make([]struct {
		funds int
		win   int
	}, numPlayers)

	resultCh := make(chan *acquire.Game)

	cpus := runtime.NumCPU()

	runtime.GOMAXPROCS(cpus)

	start := time.Now()

	for i := 0; i < cpus; i++ {
		go runGames(resultCh)
	}

	for i := 0; i < iterations; i++ {
		result := <-resultCh

		winner := result.Players[0]
		winnerIndex := 0

		for p := 0; p < len(totals); p++ {
			funds := result.Players[p].GetFunds()
			totals[p].funds += funds

			if funds > winner.GetFunds() {
				winner = result.Players[p]
				winnerIndex = p
			}
		}

		totals[winnerIndex].win++
	}

	elapsed := time.Since(start)

	fmt.Printf("Completed %d games in %s\n", iterations, elapsed)

	for p := 0; p < len(totals); p++ {
		fmt.Printf("Player %d: Won %.1f%% with avg $%d\n", p+1, float64(totals[p].win*100)/iterations, totals[p].funds/iterations)
	}
}
