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
			if h != HotelNeutral {
				if !isNeighboring[h] {
					isNeighboring[h] = true
					uniqueNeighbors++
					fillType = h
				}
			}
		}
	}

	if uniqueNeighbors == 0 {
		g.Board.Tiles[p.Row][p.Col] = fillType
	} else if uniqueNeighbors == 1 {
		if fillType == HotelNeutral {

		} else {
			g.Board.Tiles[p.Row][p.Col] = fillType
		}
	} else {
		// TODO: Merge
	}

	// TODO: "paint bucket" neutral tiles if touching ANY, regardless of spread
	// or merger results

	return nil
}
