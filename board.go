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
		row := make([]Hotel, 17)
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
