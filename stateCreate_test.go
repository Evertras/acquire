package acquire

import "testing"

func TestStateCreatePicksFromAvailableHotels(t *testing.T) {
	r, p := genGameParams()

	g := NewGame(r, p)

	g.Board.Tiles[0][0] = HotelNeutral

	createdHotel := HotelLuxor

	g.AvailableChains = []Hotel{createdHotel}
	g.State = NewStateCreate(&p[0], Piece{0, 1})

	g.State.Do(g)

	if g.Board.Tiles[0][0] != createdHotel {
		t.Error("Did not turn corner 0,0 into HotelLuxor")
	}

	if g.Board.Tiles[0][1] != createdHotel {
		t.Error("Did not turn placed piece into HotelLuxor")
	}
}

func TestStateCreateGivesFreeStock(t *testing.T) {
	r, p := genGameParams()

	g := NewGame(r, p)

	g.Board.Tiles[0][0] = HotelNeutral

	createdHotel := HotelLuxor

	g.AvailableChains = []Hotel{HotelLuxor}
	g.State = NewStateCreate(&p[0], Piece{0, 1})

	g.State.Do(g)

	if stocksOwned := p[0].GetStocks()[createdHotel]; stocksOwned != 1 {
		t.Errorf("Should have 1 stock, but have %d", stocksOwned)
	}
}

func TestStateCreateDoesNotGiveStockWhenNoneAvailable(t *testing.T) {
	r, p := genGameParams()

	g := NewGame(r, p)

	g.Board.Tiles[0][0] = HotelNeutral

	createdHotel := HotelLuxor

	g.AvailableChains = []Hotel{HotelLuxor}
	g.AvailableStocks[createdHotel] = 0
	g.State = NewStateCreate(&p[0], Piece{0, 1})

	g.State.Do(g)

	if stocksOwned := p[0].GetStocks()[createdHotel]; stocksOwned != 0 {
		t.Errorf("Should have 0 stock, but have %d", stocksOwned)
	}
}
