package acquire

// SellInfo contains information on how a Player wishes to sell their stocks
type SellInfo struct {
	Sell  int
	Trade int
	Hold  int
}

// Player is anything that can make decisions on a Game based on its current
// State and Board
type Player interface {
	// Stateful decisions that do not alter data
	PlayTile(g *Game) Piece
	BuyStocks(g *Game) []Hotel
	Merge(g *Game, choices []Hotel) Hotel
	Sell(g *Game, defunct Hotel, acquiredBy Hotel) SellInfo
	Create(g *Game, triggeringPiece Piece) Hotel

	// Held pieces
	AddPiece(p Piece)
	CanPlayPiece(g *Game) bool
	GetPieces() []Piece

	// Funds
	GetFunds() int
	AddFunds(funds int)

	// Stocks the Player owns
	GetStocks() [HotelSize]int
	GiveStocks(h Hotel, count int)
}
