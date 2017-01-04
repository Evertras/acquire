package acquire

import "math/rand"

// Some useful game constants
const (
	// StartingMoney is how much each player starts with
	StartingMoney = 6000

	// StartingStocks is how many stocks are available per hotel chain
	StartingStocks = 26

	// StartingPieces is how many pieces each player starts with
	StartingPieces = 6

	// BuyStocksPerTurn is how many stocks a player can buy in one turn
	BuyStocksPerTurn = 3
)

// Game is the current state of a game of Acquire
type Game struct {
	Board              Board
	State              State
	AvailableStocks    [HotelSize]int
	AvailableChains    []Hotel
	CurrentChainSizes  [HotelSize]int
	PieceBag           *PieceCollection
	Players            []Player
	CurrentPlayerIndex int
}

// NewGame creates a new game with the supplied RNG and Player list
func NewGame(r *rand.Rand, players []Player) *Game {
	var availableStocks [HotelSize]int

	for i := HotelFirst; i < HotelLast; i++ {
		availableStocks[i] = 26
	}

	availableChains := []Hotel{
		HotelTower,
		HotelLuxor,
		HotelAmerican,
		HotelFestival,
		HotelWorldwide,
		HotelImperial,
		HotelContinental,
	}

	g := &Game{
		Board:           NewBoard(),
		AvailableStocks: availableStocks,
		AvailableChains: availableChains,
		PieceBag:        NewPieceCollection(r),
		Players:         players,
	}

	for _, p := range players {
		for i := 0; i < StartingPieces; i++ {
			p.Draw(g)
		}
	}

	g.State = NewStatePlayTile(&g.Players[g.CurrentPlayerIndex])

	return g
}

// GetWorth returns the current stock price and bonuses for a given Hotel
func (g *Game) GetWorth(h Hotel) HotelWorth {
	return GetWorth(h, g.CurrentChainSizes[h])
}

// Advance advances the game from its current state, asking its active player
// for choices depending on the current state
func (g *Game) Advance() {
	g.State = g.State.Do(g)
}
