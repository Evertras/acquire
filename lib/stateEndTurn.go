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
	if g.CurrentPlayerIndex == len(g.Players)-1 {
		g.CurrentPlayerIndex = 0
	} else {
		g.CurrentPlayerIndex++
	}
	return NewStatePlayTile(&g.Players[g.CurrentPlayerIndex])
}
