package acquire

import (
	"math/rand"
	"testing"
)

func TestNewPieceCollection(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	c := NewPieceCollection(r)
	tileCount := countTiles(c)

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			if tileCount[row][col] != 1 {
				t.Errorf("Unexpected tile count for row %d col %d: %d", row, col, tileCount[col][row])
			}
		}
	}
}

func TestDrawPiece(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	c := NewPieceCollection(r)
	for len(c.Pieces) > 0 {
		tileCountBefore := countTiles(c)
		piecesBefore := len(c.Pieces)

		drawn := c.Draw()
		tileCountAfter := countTiles(c)
		piecesAfter := len(c.Pieces)

		if tileCountBefore[drawn.Row][drawn.Col] != 1 {
			t.Errorf("Should have one piece in collection for %v", drawn)
			t.Fail()
		}

		if tileCountAfter[drawn.Row][drawn.Col] != 0 {
			t.Errorf("Should no longer have a piece for tile %v", drawn)
		}

		if piecesAfter != piecesBefore-1 {
			t.Errorf("%d (before) should be 1 more than %d (after)", piecesBefore, piecesAfter)
		}

		for _, p := range c.Pieces {
			if p.Row == drawn.Row && p.Col == drawn.Col {
				t.Error("Found duplicate piece in bag after drawing")
			}
		}
	}
}
