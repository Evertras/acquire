package acquire

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	r, players := genGameParams()
	g := NewGame(r, players)

	for i := HotelFirst; i < HotelLast; i++ {
		if g.AvailableStocks[i] != StartingStocks {
			t.Errorf("Should have %d stocks for %d, but instead have %d", StartingStocks, i, g.AvailableStocks[i])
		}
	}

	if len(g.AvailableChains) != HotelCount {
		t.Errorf("Did not have expected number of available hotel chains at start: %v", g.AvailableChains)
	}

	numPlayers := len(g.Players)

	if g.Players == nil {
		t.Error("Missing players")
	} else if numPlayers != len(players) {
		t.Error("Players not assigned")
	}

	targetPieces := BoardWidth*BoardHeight - 6*numPlayers

	if g.PieceBag == nil {
		t.Errorf("Piece bag missing")
	} else if len(g.PieceBag.Pieces) == BoardWidth*BoardHeight {
		t.Errorf("Piece bag started full, should have %d pieces but instead has %d", targetPieces, len(g.PieceBag.Pieces))
	} else if len(g.PieceBag.Pieces) != targetPieces {
		t.Errorf("Piece bag did not get drawn from correctly, has %d pieces but should have %d", len(g.PieceBag.Pieces), targetPieces)
	}

	if startingStateType := reflect.TypeOf(g.State); startingStateType == nil {
		t.Error("Starting state not set")
	} else {
		stateName := reflect.TypeOf(g.State).Name()
		if stateName != "StatePlayTile" {
			t.Errorf("Unexpected starting state: %v (should be StatePlayTile)", reflect.TypeOf(g.State))
		}
	}

	for h, s := range g.CurrentChainSizes {
		if s != 0 {
			t.Errorf("Unexpected chain size of %d for hotel %d, should be 0", s, h)
		}
	}
}
