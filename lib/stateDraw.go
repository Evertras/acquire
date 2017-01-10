package acquire

// StateDraw is the last phase of the turn where a Piece is drawn
type StateDraw struct{}

// NewStateDraw returns a new StateDraw
func NewStateDraw() StateDraw {
	return StateDraw{}
}

// Do draws a tile from the bag if possible and gives it to the active Player
func (s StateDraw) Do(g *Game) State {
	if len(g.PieceBag.Pieces) > 0 {
		g.Players[g.CurrentPlayerIndex].AddPiece(g.PieceBag.Draw())
	}

	return NewStateEndTurn()
}
