package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	acquire "github.com/evertras/acquire/lib"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	players := []acquire.Player{
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
		acquire.NewPlayerRandom(r),
	}

	g := acquire.NewGame(r, players)

	curPlayer := -1

	for g.State != nil {
		if curPlayer != g.CurrentPlayerIndex {
			curPlayer = g.CurrentPlayerIndex
			//fmt.Printf("Player %d\n", g.CurrentPlayerIndex+1)
			//g.Board.PrintBoard(os.Stdout)
		}

		g.Advance()
	}

	g.Board.PrintBoard(os.Stdout)
	for i, p := range players {
		fmt.Printf("Player %d - $%d\n", i+1, p.GetFunds())
	}
}
