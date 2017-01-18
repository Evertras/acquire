package acquire

// StateMerge is the state when a merge between hotel chains is triggered
type StateMerge struct {
	TriggeringPiece Piece
}

// NewStateMerge creates a new merge state instance
func NewStateMerge(triggeringPiece Piece) StateMerge {
	return StateMerge{triggeringPiece}
}

// Do chooses which hotel will obtain the other(s) and proceeds to the next step
func (s StateMerge) Do(g *Game) State {
	neighbors := g.Board.GetNeighbors(s.TriggeringPiece)

	biggestHotelSize := 0
	contenders := make([]Hotel, 4)[:0]

	for _, n := range neighbors {
		h := g.Board.Tiles[n.Row][n.Col]
		c := g.CurrentChainSizes[h]

		if c >= biggestHotelSize {
			biggestHotelSize = g.CurrentChainSizes[h]
		}

		if h != HotelEmpty && h != HotelNeutral {
			contenders = append(contenders, h)
		}
	}

	biggest := make([]Hotel, 4)[:0]

	for _, h := range contenders {
		if g.CurrentChainSizes[h] == biggestHotelSize {
			present := false
			for _, b := range biggest {
				if b == h {
					present = true
				}
			}

			if !present {
				biggest = append(biggest, h)
			}
		}
	}

	winner := g.Players[g.CurrentPlayerIndex].Merge(g, biggest)

	defunctValues := make(map[Hotel]HotelWorth)

	for _, h := range contenders {
		if _, exists := defunctValues[h]; h != winner && !exists {
			defunctValues[h] = g.GetWorth(h)
			g.CurrentChainSizes[h] = 0
			g.AvailableChains = append(g.AvailableChains, h)
		}
	}

	grewBy := g.Board.Fill(s.TriggeringPiece, winner)

	g.CurrentChainSizes[winner] += grewBy

	return NewStateSell(defunctValues, winner)
}
