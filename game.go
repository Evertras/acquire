package acquire

// Some useful game constants
const (
	// StartingMoney is how much each player starts with
	StartingMoney = 6000

	// BuyStocksPerTurn is how many stocks a player can buy in one turn
	BuyStocksPerTurn = 3
)

// Game is the current state of a game of Acquire
type Game struct {
	Board             Board
	State             State
	AvailableStocks   [HotelCount]int
	AvailableChains   []Hotel
	CurrentChainSizes [HotelCount]int
	PieceBag          *PieceCollection
}

// GetWorth returns the current stock price and bonuses for a given Hotel
func (g *Game) GetWorth(h Hotel) HotelWorth {
	return GetWorth(h, g.CurrentChainSizes[h])
}
