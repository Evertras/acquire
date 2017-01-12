package acquire

// StateBuy is the state where a player can buy stocks
type StateBuy struct {
}

// NewStateBuy creates a new StateBuy instance
func NewStateBuy() StateBuy {
	return StateBuy{}
}

// Do buys stocks for the active player and goes to the draw state
func (s StateBuy) Do(g *Game) State {
	p := g.Players[g.CurrentPlayerIndex]
	toBuy := p.BuyStocks(g)

	for i := 0; i < len(toBuy); i++ {
		cost := g.GetWorth(toBuy[i]).PricePerStock
		h := toBuy[i]
		p.AddFunds(-cost)
		p.GiveStocks(h, 1)
		g.AvailableStocks[h]--
	}

	return NewStateDraw()
}
