package acquire

import (
	"math/rand"
	"reflect"
	"testing"
)

func _genTestGame() (*Game, *PlayerRandom) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	return NewGame(r, []Player{p1}), p1
}

func TestStatePlayTilePlacesNeutralTile(t *testing.T) {
	var state State
	g, _ := _genTestGame()
	state = NewStatePlayTile(&g.Players[0])

	state.Do(g)

	emptyCount := 0
	neutralCount := 0

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			h := g.Board.Tiles[row][col]
			if h == HotelEmpty {
				emptyCount++
			} else if h == HotelNeutral {
				neutralCount++
			} else {
				t.Errorf("Unexpected hotel type: %d", h)
			}
		}
	}

	targetEmpty := BoardWidth*BoardHeight - 1
	targetNeutral := 1

	if emptyCount != targetEmpty {
		t.Errorf("Expected %d empty tiles, but had %d", targetEmpty, emptyCount)
	}

	if neutralCount != targetNeutral {
		t.Errorf("Expected %d neutral tiles, but had %d", targetNeutral, neutralCount)
	}

	nextState := state.Do(g)
	nextStateType := reflect.TypeOf(nextState)

	if nextStateType == nil {
		t.Error("Next state missing")
		t.Error(g.Board.Tiles[0][1])
	} else if n := nextStateType.Name(); n != "StateBuy" {
		t.Errorf("Unexpected state %s, should be StateBuy", n)
	}
}

func TestStatePlayTileExpandsHotel(t *testing.T) {
	g, p := _genTestGame()

	g.Board.Tiles[0][0] = HotelLuxor
	g.CurrentChainSizes[HotelLuxor] = 1

	p.piecesHeld = []Piece{Piece{0, 1}}

	state := NewStatePlayTile(&g.Players[0])

	state.Do(g)

	if g.Board.Tiles[0][0] != HotelLuxor {
		t.Error("Somehow turned the corner tile into something else")
	} else if g.Board.Tiles[0][1] != HotelLuxor {
		t.Error("Did not expand Luxor as expected")
	}

	if g.CurrentChainSizes[HotelLuxor] != 2 {
		t.Errorf("Current chain size should be 2, but is instead %d", g.CurrentChainSizes[HotelLuxor])
	}
}

func TestStatePlayTileExpandsHotelIntoNeutralChain(t *testing.T) {
	g, p := _genTestGame()

	g.Board.Tiles[0][0] = HotelLuxor
	g.Board.Tiles[0][2] = HotelNeutral
	g.Board.Tiles[0][3] = HotelNeutral
	g.Board.Tiles[0][4] = HotelNeutral
	g.CurrentChainSizes[HotelLuxor] = 1

	p.piecesHeld = []Piece{Piece{0, 1}}

	state := NewStatePlayTile(&g.Players[0])

	state.Do(g)

	if g.Board.Tiles[0][0] != HotelLuxor {
		t.Error("Somehow turned the corner tile into something else")
	} else {
		for i := 0; i < 5; i++ {
			if g.Board.Tiles[0][i] != HotelLuxor {
				t.Errorf("Board at 0,%d should be %d, but is %d", i, HotelLuxor, g.Board.Tiles[0][i])
			}
		}
	}

	if g.CurrentChainSizes[HotelLuxor] != 5 {
		t.Errorf("Current chain size should be 5, but is instead %d", g.CurrentChainSizes[HotelLuxor])
	}
}

func TestStatePlayTileCreatesHotelChain(t *testing.T) {
	g, p := _genTestGame()

	g.Board.Tiles[0][0] = HotelNeutral
	p.piecesHeld = []Piece{Piece{0, 1}}

	s := NewStatePlayTile(&g.Players[0])

	nextState := s.Do(g)

	nextStateType := reflect.TypeOf(nextState)

	if nextStateType == nil {
		t.Error("Next state missing")
		t.Error(g.Board.Tiles[0][1])
	} else if n := nextStateType.Name(); n != "StateCreate" {
		t.Errorf("Unexpected state %s, should be StateCreate", n)
	}
}

func TestStatePlayTileCreatesHotelChainWhenSandwichedIn(t *testing.T) {
	g, p := _genTestGame()

	g.Board.Tiles[0][0] = HotelNeutral
	g.Board.Tiles[2][0] = HotelNeutral
	g.Board.Tiles[1][1] = HotelNeutral
	p.piecesHeld = []Piece{Piece{0, 1}}

	s := NewStatePlayTile(&g.Players[0])

	nextState := s.Do(g)

	nextStateType := reflect.TypeOf(nextState)

	if nextStateType == nil {
		t.Error("Next state missing")
		t.Error(g.Board.Tiles[0][1])
	} else if n := nextStateType.Name(); n != "StateCreate" {
		t.Errorf("Unexpected state %s, should be StateCreate", n)
	}
}

func TestStatePlayTileTriggersMerge(t *testing.T) {
	g, p := _genTestGame()

	g.Board.Tiles[0][0] = HotelLuxor
	g.Board.Tiles[0][1] = HotelLuxor
	g.Board.Tiles[2][0] = HotelContinental
	g.Board.Tiles[2][1] = HotelContinental
	p.piecesHeld = []Piece{Piece{1, 0}}

	s := NewStatePlayTile(&g.Players[0])

	nextState := s.Do(g)

	nextStateType := reflect.TypeOf(nextState)

	if nextStateType == nil {
		t.Error("Next state missing")
		t.Error(g.Board.Tiles[0][1])
	} else if n := nextStateType.Name(); n != "StateMerge" {
		t.Errorf("Unexpected state %s, should be StateMerge", n)
	}
}
