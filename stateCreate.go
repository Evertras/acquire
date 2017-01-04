package acquire

// StateCreate is the state that triggers when a new hotel chain needs to be
// chosen and then created
type StateCreate struct {
	TriggeringPiece Piece
	ActivePlayer    *Player
}

// NewStateCreate creates a new StateCreate instance
func NewStateCreate(activePlayer *Player, triggeringPiece Piece) StateCreate {
	return StateCreate{triggeringPiece, activePlayer}
}

// Do creates a new hotel chain and proceeds with the next state
func (s StateCreate) Do(g *Game) State {
	return nil
}
