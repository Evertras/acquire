package acquire

import (
	"math/rand"
	"testing"
)

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

	created := p.Create(g, 3, 3)

	if created != HotelAmerican && created != HotelLuxor {
		t.Errorf("Somehow created a hotel chain that wasn't available: %c", GetHotelInitial(created))
	}
}
