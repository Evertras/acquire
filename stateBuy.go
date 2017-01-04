package acquire

// StateBuy is the state where a player can buy stocks
type StateBuy struct {
}

// NewStateBuy creates a new StateBuy instance
func NewStateBuy() StateBuy {
	return StateBuy{}
}

// Do buys stocks for the active player and passes the turn
func (s StateBuy) Do(g *Game) State {
	return nil
}
