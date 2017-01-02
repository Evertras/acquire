package acquire

// Player is anything that can make decisions on a Game based on its current
// State and Board
type Player interface {
	// Stateful actions
	PlayTile(g *Game) (row int, col int)
	BuyStocks(g *Game) []Hotel
	Merge(g *Game, choices []Hotel) Hotel
	Sell(g *Game, defunct Hotel, acquiredBy Hotel) (sell int, trade int, hold int)
	Create(g *Game, rowPlayed int, colPlayed int) Hotel
	EndGame(g *Game)

	// Informational
	GetFunds() int
}
