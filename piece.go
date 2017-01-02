package acquire

// Piece represents a playing piece that can be put onto a certain tile
type Piece struct {
	Row int
	Col int
}

// PieceCollection is a full set of Pieces for every possible place on the board
type PieceCollection struct {
	Pieces []Piece
}

// NewPieceCollection creates a full set of Pieces; one for each tile
func NewPieceCollection() *PieceCollection {
	p := &PieceCollection{[]Piece{}}

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			p.Pieces = append(p.Pieces, Piece{row, col})
		}
	}

	return p
}
