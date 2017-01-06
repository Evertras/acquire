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
	toBuy := g.Players[g.CurrentPlayerIndex].BuyStocks(g)
	totalCost := 0

	for i := 0; i < len(toBuy); i++ {
		totalCost += g.GetWorth(toBuy[i]).PricePerStock
	}

	return NewStateDraw()
}
