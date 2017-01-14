package acquire

// StateEndTurn checks for end game states and advances to the next Player's
// turn if the Game should continue
type StateEndTurn struct{}

// NewStateEndTurn returns a new StateEndTurn
func NewStateEndTurn() StateEndTurn {
	return StateEndTurn{}
}

// Do checks the Game for an end state and advances the turn if necessary
func (s StateEndTurn) Do(g *Game) State {
	// If any chain is >40, the game immediately ends
	for h := HotelFirst; h < HotelLast; h++ {
		if g.CurrentChainSizes[h] > 40 {
			return nil
		}
	}

	if g.CanPlaceSomewhere() {
		if g.CurrentPlayerIndex == len(g.Players)-1 {
			g.CurrentPlayerIndex = 0
		} else {
			g.CurrentPlayerIndex++
		}

		cur := g.Players[g.CurrentPlayerIndex]

		if cur.CanPlayPiece(g) {
			return NewStatePlayTile(&cur)
		}
		return NewStateBuy()
	}

	// This is never reached even by a bajillion random iterations... is this
	// actually theoretically reachable?  I don't know.  May remove it.
	return nil
}
