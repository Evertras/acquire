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
			drawn := g.PieceBag.Draw()
			p.AddPiece(drawn)
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

// IsValidPlacement checks if a given piece is a valid placement, returning
// true if a piece can be placed there and false if not
func (g *Game) IsValidPlacement(p Piece) bool {
	// Must be empty
	if g.Board.Tiles[p.Row][p.Col] != HotelEmpty {
		return false
	}

	// If it would cause a merge, only one participating hotel chain can be
	// 11+ in size
	neighbors := g.Board.GetNeighbors(p)
	alreadySaw := HotelEmpty
	for _, n := range neighbors {
		h := g.Board.Tiles[n.Row][n.Col]
		if h != HotelEmpty && h != HotelNeutral {
			if g.CurrentChainSizes[h] >= 11 {
				if alreadySaw != HotelEmpty {
					return false
				}

				alreadySaw = h
			}
		}
	}

	return true
}

// CanPlaceSomewhere returns true if there is still an empty tile on the board
// where a piece can be played, false if there are no valid plays left
func (g *Game) CanPlaceSomewhere() bool {
	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			if g.IsValidPlacement(Piece{row, col}) {
				return true
			}
		}
	}

	return false
}
