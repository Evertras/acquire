package acquire

import "math/rand"

// PlayerRandom just randomly chooses anything it can
type PlayerRandom struct {
	r           *rand.Rand
	funds       int
	stocksOwned [HotelSize]int
	piecesHeld  []Piece
}

// NewPlayerRandom creates a new PlayerRandom that selects all choices at random
func NewPlayerRandom(r *rand.Rand) *PlayerRandom {
	return &PlayerRandom{r: r, funds: StartingMoney, piecesHeld: []Piece{}}
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
func (p *PlayerRandom) GetStocks() [HotelSize]int {
	return p.stocksOwned
}

// BuyStocks picks a random set of available stocks and buys as many as possible
func (p *PlayerRandom) BuyStocks(g *Game) []Hotel {
	bought := []Hotel{}
	startingAvailable := make([]int, HotelSize)

	copy(startingAvailable, g.AvailableStocks[:])

	for i := 0; i < BuyStocksPerTurn; i++ {
		available := []Hotel{}
		var prices [HotelSize]int

		for h, s := range startingAvailable {
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

			startingAvailable[choice]--

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
func (p *PlayerRandom) Create(g *Game, triggeringPiece Piece) Hotel {
	// assume there is at least one available to have gotten here
	return g.AvailableChains[p.r.Intn(len(g.AvailableChains))]
}

// Draw will draw a piece from the piece bag and hold it
func (p *PlayerRandom) Draw(g *Game) {
	p.piecesHeld = append(p.piecesHeld, g.PieceBag.Draw())
}

// PlayTile selects a random held tile to play
func (p *PlayerRandom) PlayTile(g *Game) Piece {
	l := len(p.piecesHeld)
	i := p.r.Intn(l)
	choice := p.piecesHeld[i]
	p.piecesHeld[i] = p.piecesHeld[l-1]
	p.piecesHeld = p.piecesHeld[:l-1]
	return choice
}

// Sell randomly chooses how much stock to hold or sell, will never trade
func (p *PlayerRandom) Sell(g *Game, defunct Hotel, acquiredBy Hotel) SellInfo {
	s := SellInfo{}

	owned := p.stocksOwned[defunct]

	for i := 0; i < owned; i++ {
		choice := p.r.Intn(2)

		switch choice {
		case 0:
			s.Hold++
		case 1:
			s.Sell++
		}
	}

	return s
}

// GiveStocks gives the specified count of stocks to the player
func (p *PlayerRandom) GiveStocks(h Hotel, count int) {
	p.stocksOwned[h] += count
}
