package acquire

import (
	"fmt"
	"testing"
)

func TestStateMergeFillsCorrectly(t *testing.T) {
	r, ps := genGameParams()
	g := NewGame(r, ps)

	g.Board.Tiles[0][0] = HotelLuxor
	g.Board.Tiles[0][2] = HotelAmerican
	g.Board.Tiles[0][3] = HotelAmerican

	g.CurrentChainSizes[HotelLuxor] = 1
	g.CurrentChainSizes[HotelAmerican] = 2

	s := NewStateMerge(Piece{0, 1})

	s.Do(g)

	for i := 0; i < 4; i++ {
		if h := g.Board.Tiles[0][i]; h != HotelAmerican {
			t.Errorf("Piece at 0,%d is %d, should be %d", i, h, HotelAmerican)
		}
	}

	if s := g.CurrentChainSizes[HotelLuxor]; s != 0 {
		t.Errorf("Should have 0 pieces of HotelLuxor, have %d", s)
	}

	if s := g.CurrentChainSizes[HotelAmerican]; s != 4 {
		t.Errorf("Should have 4 pieces of HotelAmerican, have %d", s)
	}
}

func TestStateMergeHandlesFourWayCorrectly(t *testing.T) {
	r, ps := genGameParams()
	g := NewGame(r, ps)

	fmt.Println("-----")

	g.Board.Tiles[1][0] = HotelLuxor
	g.Board.Tiles[0][1] = HotelLuxor
	g.Board.Tiles[0][0] = HotelLuxor
	g.Board.Tiles[1][2] = HotelAmerican
	g.Board.Tiles[2][1] = HotelAmerican

	g.CurrentChainSizes[HotelLuxor] = 3
	g.CurrentChainSizes[HotelAmerican] = 2

	s := NewStateMerge(Piece{1, 1})

	s.Do(g)

	if s := g.CurrentChainSizes[HotelLuxor]; s != 6 {
		t.Errorf("HotelLuxor should have 6, but has %d", s)
	}

	if s := g.CurrentChainSizes[HotelAmerican]; s != 0 {
		t.Errorf("HotelAmerican should have 0, but has %d", s)
	}
}
