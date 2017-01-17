package acquire

// StateEndGame cleans up the game state and cashes out all remaining held stocks
type StateEndGame struct{}

// NewStateEndGame returns a new StateEndGame
func NewStateEndGame() StateEndGame {
	return StateEndGame{}
}

// Do cleans up the game state and cashes out all remaining held stocks
func (s StateEndGame) Do(g *Game) State {
	worths := make(map[Hotel]HotelWorth)

	for h := HotelFirst; h < HotelLast; h++ {
		if g.AvailableStocks[h] != StartingStocks {
			worths[h] = g.GetWorth(h)
		}
	}

	for h, d := range worths {
		holders := make([]Player, len(g.Players))[:0]

		for _, p := range g.Players {
			if p.GetStocks()[h] > 0 {
				holders = append(holders, p)
			}
		}

		sortHolders(holders, h)

		handleBonuses(holders, h, d)

		// Goes from highest holder to smallest
		for _, player := range holders {
			held := player.GetStocks()[h]

			// Cash out
			player.AddFunds(held * d.PricePerStock)
			player.GiveStocks(h, -held)
			g.AvailableStocks[h] += held
		}
	}

	return nil
}
