package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	acquire "github.com/evertras/acquire/lib"
)

func runGame(r *rand.Rand) *acquire.Game {
	players := []acquire.Player{
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
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

	return g
}

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	totals := []int{0, 0, 0, 0, 0}
	iters := 100000

	for i := 0; i < iters; i++ {
		result := runGame(r)

		for p := 0; p < len(totals); p++ {
			totals[p] += result.Players[p].GetFunds()
		}
	}

	for p := 0; p < len(totals); p++ {
		fmt.Printf("Player %d: $%d\n", p+1, totals[p]/iters)
	}
}
