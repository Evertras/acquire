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
	players := []acquire.Player{acquire.NewPlayerRandom(r), acquire.NewPlayerRandom(r)}

	g := acquire.NewGame(r, players)

	for {
		fmt.Printf("Player %d\n", g.CurrentPlayerIndex+1)
		g.Board.PrintBoard(os.Stdout)
		g.Advance()
	}
}
