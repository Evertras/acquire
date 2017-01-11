package acquire

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrintEmptyBoard(t *testing.T) {
	buf := &bytes.Buffer{}
	bufExpected := &bytes.Buffer{}

	fmt.Fprintln(bufExpected, "   A B C D E F G H I")
	fmt.Fprintln(bufExpected, " 1 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 2 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 3 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 4 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 5 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 6 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 7 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 8 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 9 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, "10 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, "11 - - - - - - - - - ")

	b := NewBoard()

	b.PrintBoard(buf)

	result := buf.String()
	expected := bufExpected.String()

	if result != expected {
		t.Errorf("Did not get expected board result:\n\n%s\n\n%s", result, expected)

		for i, r := range result {
			if expected[i] != result[i] {
				t.Errorf("Badness %c vs %c at %d", r, expected[i], i)
			}
		}
	}
}

func TestBoardGetNeighbors(t *testing.T) {
	b := NewBoard()

	if middleNeighbors := b.GetNeighbors(Piece{BoardHeight / 2, BoardWidth / 2}); len(middleNeighbors) != 4 {
		t.Errorf("Unexpected number of neighbors in middle: %d (should be 4)", len(middleNeighbors))
	}

	sidePieces := []struct {
		name  string
		piece Piece
	}{
		{
			"top",
			Piece{0, BoardWidth / 2},
		},
		{
			"right",
			Piece{BoardHeight / 2, BoardWidth - 1},
		},
		{
			"bottom",
			Piece{BoardHeight - 1, BoardWidth / 2},
		},
		{
			"left",
			Piece{BoardHeight / 2, 0},
		},
	}

	for _, pieces := range sidePieces {
		if neighbors := b.GetNeighbors(pieces.piece); len(neighbors) != 3 {
			t.Errorf("Unexpected number of neighbors at %s: %d (should be 3)", pieces.name, len(neighbors))
		}
	}

	if cornerNeighbors := b.GetNeighbors(Piece{0, 0}); len(cornerNeighbors) != 2 {
		t.Errorf("Unexpected number of neighbors in corner: %d (should be 2)", len(cornerNeighbors))
	}
}

func TestBoardFillFull(t *testing.T) {
	b := NewBoard()

	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
			b.Tiles[i][j] = HotelNeutral
		}
	}

	filledWith := HotelLuxor

	fillCount := b.Fill(Piece{3, 3}, filledWith)

	if fillCount != BoardWidth*BoardHeight {
		t.Errorf("Should have filled %d, but filled %d", BoardWidth*BoardHeight, fillCount)
	}

	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
			if b.Tiles[i][j] != filledWith {
				t.Errorf("Tile (%d,%d) should be %d, is %d", i, j, filledWith, b.Tiles[i][j])
			}
		}
	}
}

func TestBoardFillSandwich(t *testing.T) {
	b := NewBoard()

	b.Tiles[0][0] = HotelNeutral
	b.Tiles[0][1] = HotelNeutral
	b.Tiles[0][3] = HotelNeutral
	b.Tiles[0][4] = HotelNeutral

	filledWith := HotelLuxor

	fillCount := b.Fill(Piece{0, 2}, filledWith)

	if fillCount != 5 {
		t.Errorf("Should have filled 5, but filled %d", fillCount)
	}
}

func TestBoardFillExpanded(t *testing.T) {
	b := NewBoard()
	filledWith := HotelLuxor

	b.Tiles[0][0] = filledWith
	b.Tiles[0][1] = filledWith
	b.Tiles[0][3] = HotelAmerican
	b.Tiles[0][4] = HotelAmerican

	fillCount := b.Fill(Piece{0, 2}, filledWith)

	if fillCount != 3 {
		t.Errorf("Should have filled 3, but filled %d", fillCount)
	}
}
