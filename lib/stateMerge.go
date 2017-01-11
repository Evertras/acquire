package acquire

// StateMerge is the state when a merge between hotel chains is triggered
type StateMerge struct {
	TriggeringPiece Piece
}

// NewStateMerge creates a new merge state instance
func NewStateMerge(triggeringPiece Piece) StateMerge {
	return StateMerge{triggeringPiece}
}

// Do chooses which hotel will obtain the other(s) and proceeds to the next step
func (s StateMerge) Do(g *Game) State {
	return nil
}
