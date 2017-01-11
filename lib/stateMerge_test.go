package acquire

import "testing"

func TestStateMergeFillsInCorrectly(t *testing.T) {
	r, ps := genGameParams()
	g := NewGame(r, ps)

	g.Board.Tiles[0][0] = HotelLuxor
	g.Board.Tiles[0][2] = HotelAmerican
	g.Board.Tiles[0][3] = HotelAmerican

	s := NewStateMerge(Piece{0, 1})

	s.Do(g)

	for i := 0; i < 4; i++ {
		if g.Board.Tiles[0][i] != HotelAmerican {
			t.Errorf("Piece at 0,%d is %d, should be %d", i, g.Board.Tiles[0][i], HotelAmerican)
		}
	}
}
