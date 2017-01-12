package acquire

import (
	"math/rand"
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

func BenchmarkPlayGame(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)

	g := NewGame(r, []Player{p1, p2})

	for len(g.PieceBag.Pieces) > 3 {
		g.Advance()
	}
}

func TestGameStocksAlwaysStableCount(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)

	g := NewGame(r, []Player{p1, p2})

	for i := 0; i < 10000; i++ {
		for len(g.PieceBag.Pieces) > 3 {
			g.Advance()

			for h := HotelFirst; h < HotelLast; h++ {
				totalStocks := g.AvailableStocks[h] + p1.GetStocks()[h] + p2.GetStocks()[h]

				if totalStocks != StartingStocks {
					t.Errorf("Stocks for %c have improper total", GetHotelInitial(h))
					t.Error(reflect.TypeOf(g.State).Name())
					t.Error(g.AvailableStocks)
					t.Error(p1.stocksOwned)
					t.Error(p2.stocksOwned)
					t.FailNow()
				}
			}
		}
	}
}
