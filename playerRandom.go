package acquire

import "math/rand"

// PlayerRandom just randomly chooses anything it can
type PlayerRandom struct {
	r           *rand.Rand
	funds       int
	stocksOwned [HotelCount]int
}

// NewPlayerRandom creates a new PlayerRandom that selects all choices at random
func NewPlayerRandom(r *rand.Rand) PlayerRandom {
	return PlayerRandom{r: r, funds: StartingMoney}
}

// GetFunds gets the current funds of the player
func (p *PlayerRandom) GetFunds() int {
	return p.funds
}

// AddFunds adds the given amount to the player's funds
func (p *PlayerRandom) AddFunds(funds int) {
	p.funds += funds
}

// GetStocks returns the current counts of what stocks the player owns
func (p *PlayerRandom) GetStocks() [HotelCount]int {
	return p.stocksOwned
}

// BuyStocks picks a random set of available stocks and buys as many as possible
func (p *PlayerRandom) BuyStocks(g *Game) []Hotel {
	bought := []Hotel{}

	for i := 0; i < BuyStocksPerTurn; i++ {
		available := []Hotel{}
		var prices [HotelCount]int

		for h, s := range g.AvailableStocks {
			if s > 0 {
				prices[h] = g.GetWorth(Hotel(h)).PricePerStock

				if prices[h] <= p.funds {
					available = append(available, Hotel(h))
				}
			}
		}

		l := len(available)

		if l > 0 {
			choice := available[p.r.Intn(l)]

			p.funds -= prices[choice]
			p.stocksOwned[choice]++

			bought = append(bought, choice)
		}
	}

	return bought
}

// Merge randomly picks among equally sized hotel chains to decide which will
// acquire the others
func (p *PlayerRandom) Merge(g *Game, choices []Hotel) Hotel {
	return choices[p.r.Intn(len(choices))]
}

// Create randomly picks an available hotel chain and creates it
func (p *PlayerRandom) Create(g *Game, rowPlayed int, colPlayed int) Hotel {
	// assume there is at least one available to have gotten here
	return g.AvailableChains[p.r.Intn(len(g.AvailableChains))]
}
