package acquire

import "testing"

func TestNewPieceCollection(t *testing.T) {
	var tileCount [BoardWidth][BoardHeight]int
	c := NewPieceCollection()

	for _, p := range c.Pieces {
		tileCount[p.Col][p.Row]++
	}

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			if tileCount[col][row] != 1 {
				t.Errorf("Unexpected tile count for row %d col %d: %d", row, col, tileCount[col][row])
			}
		}
	}
}
