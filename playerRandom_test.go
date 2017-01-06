package acquire

import (
	"math/rand"
	"testing"
)

func TestPlayerRandomFitsInterface(t *testing.T) {
	var p Player
	p = NewPlayerRandom(rand.New(rand.NewSource(0)))
	p.GetFunds()
}

func TestPlayerRandomStartsWithCorrectFunds(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))

	funds := p.GetFunds()

	if funds != StartingMoney {
		t.Errorf("Random player is cheating!  Started with %d but expected %d", funds, StartingMoney)
	}
}

func TestPlayerRandomAddsFunds(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))

	startingFunds := p.GetFunds()
	toAdd := 500

	p.AddFunds(toAdd)

	afterFunds := p.GetFunds()
	expectedFunds := startingFunds + toAdd

	if afterFunds != expectedFunds {
		t.Errorf("Funds not added correctly; got %d but expected %d (%d + %d)", afterFunds, expectedFunds, startingFunds, toAdd)
	}
}

func TestPlayerRandomStartsWithNoStocksOwned(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))

	stocks := p.GetStocks()

	for h, s := range stocks {
		if s != 0 {
			t.Errorf("Started with %d stocks in %v", s, h)
		}
	}
}

func TestPlayerRandomBuyStock(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))
	g := &Game{}

	hotel := HotelLuxor
	expectedPricePer := 200
	startingStockCount := 5
	ownedBefore := p.stocksOwned[hotel]

	g.AvailableStocks[hotel] = startingStockCount
	g.CurrentChainSizes[hotel] = 2

	bought := p.BuyStocks(g)

	ownedAfter := p.stocksOwned[hotel]

	if len(bought) != BuyStocksPerTurn {
		t.Errorf("Expected to buy %d, instead bought %d", BuyStocksPerTurn, len(bought))
	}

	expectedFunds := StartingMoney - expectedPricePer*BuyStocksPerTurn

	if p.funds != expectedFunds {
		t.Errorf("Expected funds of %d, but had %d after purchase", expectedFunds, p.funds)
	}

	if ownedBefore != 0 {
		t.Errorf("Started with %d stocks, should have zero", ownedBefore)
	}

	if ownedAfter != len(bought) {
		t.Errorf("Expected to own %d stocks but instead own %d", len(bought), ownedAfter)
	}
}

func TestPlayerRandomBuyStockRespectsFunds(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))
	g := &Game{}

	hotel := HotelLuxor
	expectedPricePer := 200
	startingStockCount := 5

	g.AvailableStocks[hotel] = startingStockCount
	g.CurrentChainSizes[hotel] = 2

	p.funds = expectedPricePer * 2

	bought := p.BuyStocks(g)

	if len(bought) != 2 {
		t.Errorf("Expected to buy 2 stocks, but instead bought %d", len(bought))
	}

	expectedFunds := 0

	if p.funds != expectedFunds {
		t.Errorf("Expected player to have %d funds, but instead has %d", expectedFunds, p.funds)
	}
}

func TestPlayerRandomBuyStockRespectsAvailableCount(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))
	g := &Game{}

	hotel := HotelLuxor
	startingStockCount := 2

	g.AvailableStocks[hotel] = startingStockCount
	g.CurrentChainSizes[hotel] = 2

	bought := p.BuyStocks(g)

	if len(bought) != startingStockCount {
		t.Errorf("Expected to buy %d stocks (starting available count), but instead bought %d", startingStockCount, len(bought))
	}
}

func TestPlayerRandomBuyStockPassesIfNoneAvailable(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))
	g := &Game{}

	bought := p.BuyStocks(g)

	if len(bought) != 0 {
		t.Errorf("Somehow bought %d stocks when none were available", len(bought))
	}

	if p.funds != StartingMoney {
		t.Errorf("Somehow spent %d money when no stock was available", StartingMoney-p.funds)
	}
}

func TestPlayerRandomMerge(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))
	g := &Game{}
	hotels := []Hotel{HotelAmerican, HotelLuxor}

	choice := p.Merge(g, hotels)

	// Don't care which, it's random
	if choice != HotelAmerican && choice != HotelLuxor {
		t.Errorf("Somehow chose the hotel %d that wasn't in the list to merge", choice)
	}
}

func TestPlayerRandomCreate(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))
	g := &Game{}

	g.AvailableChains = []Hotel{HotelAmerican, HotelLuxor}

	created := p.Create(g, Piece{3, 3})

	if created != HotelAmerican && created != HotelLuxor {
		t.Errorf("Somehow created a hotel chain that wasn't available: %c", GetHotelInitial(created))
	}
}

func TestPlayerRandomGiveStocks(t *testing.T) {
	p := NewPlayerRandom(rand.New(rand.NewSource(0)))

	p.GiveStocks(HotelLuxor, 2)

	if p.stocksOwned[HotelLuxor] != 2 {
		t.Errorf("Should have 2 stocks, but have %d", p.stocksOwned[HotelLuxor])
	}
}

func TestPlayerRandomDrawsUniqueTiles(t *testing.T) {
	r := rand.New(rand.NewSource(0))

	// Churn through a ton of cases with the same RNG generator to try as many
	// possible cases as sanely possible via brute force
	for i := 0; i < 100000; i++ {
		p := NewPlayerRandom(r)
		g := &Game{
			PieceBag: NewPieceCollection(r),
		}

		startingCount := BoardWidth * BoardHeight

		if len(g.PieceBag.Pieces) != startingCount {
			t.Errorf("Started with %d pieces, expected %d", len(g.PieceBag.Pieces), startingCount)
		}

		drawCount := 10

		for j := 0; j < drawCount; j++ {
			p.Draw(g)
		}

		if len(p.piecesHeld) != drawCount {
			t.Errorf("Wanted to be holding %d but instead holding %d", drawCount, len(p.piecesHeld))
		}

		for i1, p1 := range p.piecesHeld {
			for i2, p2 := range p.piecesHeld {
				if i1 != i2 && p2 == p1 {
					t.Errorf("Found duplicate tile being held")
				}
			}
		}
	}
}

func TestPlayerRandomSell(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p := NewPlayerRandom(r)
	g := NewGame(r, []Player{p})

	defunct := HotelLuxor
	acquiredBy := HotelAmerican
	totalOwned := 1000000

	p.stocksOwned[defunct] = totalOwned

	sold := p.Sell(g, defunct, acquiredBy)

	if sold.Hold+sold.Sell != totalOwned {
		t.Errorf("Should have held/sold %d but held/sold %d", totalOwned, sold.Hold+sold.Sell)
	}

	if sold.Hold == 0 {
		t.Error("Should have at least one stock held, but have 0")
	}

	if sold.Sell == 0 {
		t.Error("Should have at least one stock sold, but have 0")
	}
}
