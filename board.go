package acquire

import (
	"fmt"
	"io"
)

// Board describes a snapshot of the board state
type Board struct {
	Tiles [][]Tile
}

// NewBoard creates a new Board and initializes it as empty
func NewBoard() Board {
	b := Board{
		[][]Tile{},
	}

	for i := 0; i < 11; i++ {
		row := make([]Tile, 17)
		b.Tiles = append(b.Tiles, row)
	}

	return b
}

// PrintBoard prints the board state to stdout
func (b *Board) PrintBoard(out io.Writer) {
	fmt.Fprintln(out, "   A B C D E F G H I")
	for i := 0; i < 11; i++ {
		fmt.Fprintf(out, "%2d ", i+1)
		for j := 0; j < 9; j++ {
			fmt.Fprintf(out, "%c ", GetHotelInitial(b.Tiles[i][j].Hotel))
		}
		fmt.Fprintln(out)
	}
}
