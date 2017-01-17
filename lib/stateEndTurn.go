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
	chainPresent := false
	someChainStillSmall := false

	// If any chain is >40 or all present chains have >10, the game can be ended
	for h := HotelFirst; h < HotelLast; h++ {
		size := g.CurrentChainSizes[h]
		if size > 40 {
			return nil
		}

		if size > 0 {
			chainPresent = true

			if size <= 10 {
				someChainStillSmall = true
			}
		}
	}

	if chainPresent && !someChainStillSmall {
		return nil
	}

	if g.CurrentPlayerIndex == len(g.Players)-1 {
		g.CurrentPlayerIndex = 0
	} else {
		g.CurrentPlayerIndex++
	}

	cur := g.Players[g.CurrentPlayerIndex]

	return NewStatePlayTile(&cur)
}
