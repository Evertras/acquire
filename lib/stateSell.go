package acquire

// StateSell handles the sale of all defunct hotel stocks for all players
type StateSell struct {
	defunct    map[Hotel]HotelWorth
	acquiredBy Hotel
}

// NewStateSell creates a new StateSell with the defunct hotels
func NewStateSell(defunct map[Hotel]HotelWorth, acquiredBy Hotel) *StateSell {
	return &StateSell{defunct, acquiredBy}
}

func sortHolders(holders []Player, h Hotel) {
	// Bubble sort is fine on a tiny data set of 4-5 max length
	l := len(holders)
	madeSwap := true
	for i := l - 1; i > 0 && madeSwap; i-- {
		madeSwap = false
		for j := 0; j < i; j++ {
			h1 := holders[j].GetStocks()[h]
			h2 := holders[j+1].GetStocks()[h]

			if h1 < h2 {
				tmp := holders[j+1]
				holders[j+1] = holders[j]
				holders[j] = tmp
				madeSwap = true
			}
		}
	}
}

// Returns the funds split evenly and rounded up to the nearest 100
func getFundSplit(total int, split int) int {
	f := total / split

	if f%100 != 0 {
		return (f/100 + 1) * 100
	}

	return f
}

func handleBonuses(holders []Player, h Hotel, d HotelWorth) {
	l := len(holders)

	if l == 1 {
		holders[0].AddFunds(d.MajorityHolderBonusFirst + d.MajorityHolderBonusSecond)
	} else {
		highest := holders[0].GetStocks()[h]

		// If at least two people are tied for first, split evenly
		if holders[0].GetStocks()[h] == holders[1].GetStocks()[h] {
			tied := make([]Player, l)[:0]
			for i := 0; i < l && holders[i].GetStocks()[h] == highest; i++ {
				tied = append(tied, holders[i])
			}

			award := getFundSplit(d.MajorityHolderBonusFirst+d.MajorityHolderBonusSecond, len(tied))

			for _, p := range tied {
				p.AddFunds(award)
			}
		} else {
			// Top holder is lone winner of biggest bonus
			holders[0].AddFunds(d.MajorityHolderBonusFirst)

			second := holders[1].GetStocks()[h]

			if l > 2 && second == holders[2].GetStocks()[h] {
				tied := make([]Player, l-1)[:0]
				for i := 1; i < l && holders[i].GetStocks()[h] == second; i++ {
					tied = append(tied, holders[i])
				}

				award := getFundSplit(d.MajorityHolderBonusSecond, len(tied))

				for _, p := range tied {
					p.AddFunds(award)
				}
			} else {
				holders[1].AddFunds(d.MajorityHolderBonusSecond)
			}
		}
	}
}

// Do sells all defunct stocks for all players before continuing
func (s *StateSell) Do(g *Game) State {
	for h, d := range s.defunct {
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
			sold := player.Sell(g, h, s.acquiredBy)

			// First handle sells
			player.AddFunds(sold.Sell * d.PricePerStock)
			player.GiveStocks(h, -sold.Sell)
			g.AvailableStocks[h] += sold.Sell

			// Now handle trades - we assume sold.Trade%2 == 0
			player.GiveStocks(h, -sold.Trade)
			g.AvailableStocks[h] += sold.Trade
			player.GiveStocks(s.acquiredBy, sold.Trade/2)
			g.AvailableStocks[s.acquiredBy] -= sold.Trade / 2
		}
	}

	return NewStateBuy()
}
