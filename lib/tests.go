package acquire

import "math/rand"

func countTiles(c *PieceCollection) [BoardHeight][BoardWidth]int {
	var tileCount [BoardHeight][BoardWidth]int

	for _, p := range c.Pieces {
		tileCount[p.Row][p.Col]++
	}

	return tileCount
}

func _genTestGame() (*Game, *PlayerRandom) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	return NewGame(r, []Player{p1}), p1
}

func genGameParams() (r *rand.Rand, players []Player) {
	r = rand.New(rand.NewSource(0))
	players = []Player{NewPlayerRandom(r)}
	return
}
