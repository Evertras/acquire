package acquire

import (
	"fmt"
	"io"
)

const (
	// BoardWidth is how wide the board is (from A - I)
	BoardWidth = 9

	// BoardHeight is how tall the board is
	BoardHeight = 11
)

// Board describes a snapshot of the board state
type Board struct {
	Tiles [][]Hotel
}

// NewBoard creates a new Board and initializes it as empty
func NewBoard() Board {
	b := Board{
		[][]Hotel{},
	}

	for i := 0; i < BoardHeight; i++ {
		row := make([]Hotel, BoardWidth)
		b.Tiles = append(b.Tiles, row)
	}

	return b
}

// PrintBoard prints the board state to stdout
func (b *Board) PrintBoard(out io.Writer) {
	fmt.Fprintln(out, "   A B C D E F G H I")
	for i := 0; i < BoardHeight; i++ {
		fmt.Fprintf(out, "%2d ", i+1)
		for j := 0; j < BoardWidth; j++ {
			fmt.Fprintf(out, "%c ", GetHotelInitial(b.Tiles[i][j]))
		}
		fmt.Fprintln(out)
	}
}

// GetNeighbors returns coordinates of all valid neighbors to the given piece
func (b *Board) GetNeighbors(p Piece) []Piece {
	n := make([]Piece, 4)[:0]

	if p.Col > 0 {
		n = append(n, Piece{Row: p.Row, Col: p.Col - 1})
	}

	if p.Col < BoardWidth-1 {
		n = append(n, Piece{Row: p.Row, Col: p.Col + 1})
	}

	if p.Row > 0 {
		n = append(n, Piece{Row: p.Row - 1, Col: p.Col})
	}

	if p.Row < BoardHeight-1 {
		n = append(n, Piece{Row: p.Row + 1, Col: p.Col})
	}

	return n
}

// Fill will flood fill the board with the given hotel, returns how many Tiles
// were filled
func (b *Board) Fill(origin Piece, fillWith Hotel) int {
	total := 1
	b.Tiles[origin.Row][origin.Col] = fillWith

	pieces := b.GetNeighbors(origin)

	for _, p := range pieces {
		if h := b.Tiles[p.Row][p.Col]; h != HotelEmpty && h != fillWith {
			total += b.Fill(p, fillWith)
		}
	}

	return total
}
