package acquire

import "math/rand"

// Piece represents a playing piece that can be put onto a certain tile
type Piece struct {
	Row int
	Col int
}

// PieceCollection is a full set of Pieces for every possible place on the board
type PieceCollection struct {
	Pieces []Piece
	r      *rand.Rand
}

// NewPieceCollection creates a full set of Pieces; one for each tile
func NewPieceCollection(r *rand.Rand) *PieceCollection {
	p := &PieceCollection{[]Piece{}, r}

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			p.Pieces = append(p.Pieces, Piece{row, col})
		}
	}

	return p
}

// Draw destructively draws a random piece and returns it, removing it from
// the collection
func (p *PieceCollection) Draw() Piece {
	l := len(p.Pieces)
	i := p.r.Intn(l)
	drawn := p.Pieces[i]

	// Swap whatever's last into what we drew, order doesn't matter
	p.Pieces[i] = p.Pieces[l-1]
	p.Pieces = p.Pieces[:l-1]

	return drawn
}
