package acquire

import (
	"math/rand"
	"os"
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

	for g.State != nil {
		g.Advance()
	}
}

func TestGameStocksAlwaysStableCount(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)

	g := NewGame(r, []Player{p1, p2})

	for j := 0; j < 10000; j++ {
		for i := 0; i < 10000; i++ {
			for g.State != nil {
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
}

func TestGamePlacementValidOnEmptyBoard(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			if !g.IsValidPlacement(Piece{row, col}) {
				t.Errorf("Piece at %d,%d should be valid", row, col)
			}
		}
	}
}

func TestGameExistingHotelMakesInvalidPlacement(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < BoardWidth; col++ {
			// Overkill, but hey, why not
			for h := HotelFirst; h < HotelLast; h++ {
				p := Piece{row, col}
				g.Board.Tiles[row][col] = h
				if g.IsValidPlacement(p) {
					t.Errorf("Piece at %d,%d should not be valid while taken up by %c", row, col, GetHotelInitial(h))
				}
			}
		}
	}
}

func TestGamePieceThatWouldMergeTwoBigChainsIsInvalid(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)

	g := NewGame(r, []Player{p1})

	g.Board.Tiles[0][0] = HotelLuxor
	g.Board.Tiles[0][2] = HotelAmerican

	triggerPiece := Piece{0, 1}

	// Cheating, the pieces don't exist but this is enough to check
	g.CurrentChainSizes[HotelLuxor] = 10
	g.CurrentChainSizes[HotelAmerican] = 10

	if !g.IsValidPlacement(triggerPiece) {
		t.Error("Should be valid to join two hotel chains with size 10")
	}

	g.CurrentChainSizes[HotelLuxor] = 11
	g.CurrentChainSizes[HotelAmerican] = 11

	if g.IsValidPlacement(triggerPiece) {
		t.Error("Should not be valid to join two giant hotel chains at size 11")
	}
}

func TestGameCanGrowBigChain(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	g.Board.Tiles[0][0] = HotelLuxor
	// Cheat about how big it is
	g.CurrentChainSizes[HotelLuxor] = 20

	if !g.IsValidPlacement(Piece{0, 1}) {
		t.Error("Should be able to place tile at 0,1")
	}
}

func BenchmarkCanPlaceSomewhere(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	b.ResetTimer()

	g.CanPlaceSomewhere()
}

func TestGameCanPlaceAtStart(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	if !g.CanPlaceSomewhere() {
		t.Error("Should be able to place somewhere at the start of the match")
	}
}

func TestGameCanPlaceWhenBigHotelsCanStillGrow(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < 4; col++ {
			g.Board.Tiles[row][col] = HotelLuxor
			g.CurrentChainSizes[HotelLuxor]++
		}

		// Leave a 2 wide gutter in between, can technically place in either side
		for col := 6; col < BoardWidth; col++ {
			g.Board.Tiles[row][col] = HotelAmerican
			g.CurrentChainSizes[HotelAmerican]++
		}
	}

	if !g.CanPlaceSomewhere() {
		g.Board.PrintBoard(os.Stdout)
		t.Error("Should be able to place on board")
	}
}

func TestGameCannotPlaceWhenBigHotelsOnlyOneApart(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < 4; col++ {
			g.Board.Tiles[row][col] = HotelLuxor
			g.CurrentChainSizes[HotelLuxor]++
		}

		// Leave a 1 wide gutter in between, can't place in gutter now
		for col := 5; col < BoardWidth; col++ {
			g.Board.Tiles[row][col] = HotelAmerican
			g.CurrentChainSizes[HotelAmerican]++
		}
	}

	if g.Board.Tiles[0][4] != HotelEmpty {
		t.Error("0,4 should be empty but isn't, test isn't set up right")
	}

	if g.CanPlaceSomewhere() {
		g.Board.PrintBoard(os.Stdout)
		t.Error("Shouldn't be able to place on board")
	}
}

func TestGameCanPlaceWhenBigHotelWouldEatBigNeutralBlock(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	g := NewGame(r, []Player{p1})

	for row := 0; row < BoardHeight; row++ {
		for col := 0; col < 4; col++ {
			g.Board.Tiles[row][col] = HotelLuxor
			g.CurrentChainSizes[HotelLuxor]++
		}

		// Leave a 1 wide gutter in between, can't place in gutter now
		for col := 5; col < BoardWidth; col++ {
			g.Board.Tiles[row][col] = HotelNeutral
			// Shouldn't actually do anything, but just to be safe...
			g.CurrentChainSizes[HotelNeutral]++
		}
	}

	if g.Board.Tiles[0][4] != HotelEmpty {
		t.Error("0,4 should be empty but isn't, test isn't set up right")
	}

	if !g.CanPlaceSomewhere() {
		g.Board.PrintBoard(os.Stdout)
		t.Error("Should be able to place on board")
	}
}
