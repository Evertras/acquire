package acquire

// StatePlayTile is the start of the turn where a player needs to play a tile
type StatePlayTile struct {
	ActivePlayer *Player
}

// NewStatePlayTile creates a new StatePlayTile with the given active player
func NewStatePlayTile(activePlayer *Player) StatePlayTile {
	return StatePlayTile{activePlayer}
}

// Do goes through the play tile step for the active player
func (s StatePlayTile) Do(g *Game) State {
	activePlayer := g.Players[g.CurrentPlayerIndex]
	p := activePlayer.PlayTile(g)

	neighbors := g.Board.GetNeighbors(p)
	var isNeighboring [HotelSize]bool
	uniqueNeighbors := 0
	fillType := HotelNeutral

	for _, n := range neighbors {
		h := g.Board.Tiles[n.Row][n.Col]
		if h != HotelEmpty {
			if !isNeighboring[h] {
				isNeighboring[h] = true
				uniqueNeighbors++

				if h != HotelNeutral {
					fillType = h
				}
			}
		}
	}

	if uniqueNeighbors > 1 && !(uniqueNeighbors == 2 && isNeighboring[HotelNeutral]) {
		return NewStateMerge(p)
	}

	fillCount := g.Board.Fill(p, fillType)

	if fillType != HotelNeutral {
		g.CurrentChainSizes[fillType] += fillCount
	}

	if uniqueNeighbors == 1 && fillType == HotelNeutral && len(g.AvailableChains) > 0 {
		return NewStateCreate(s.ActivePlayer, p)
	}

	return NewStateBuy()
}
