package acquire

import (
	"math/rand"
	"testing"
)

func TestStateEndGameCashesOutCorrectly(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)

	g := NewGame(r, []Player{p1, p2})

	majority := 5
	minority := 3

	p1.stocksOwned[HotelLuxor] = majority
	p2.stocksOwned[HotelLuxor] = minority

	g.AvailableStocks[HotelLuxor] -= majority + minority

	p1.funds = 0
	p2.funds = 0

	g.CurrentChainSizes[HotelLuxor] = 5

	worth := g.GetWorth(HotelLuxor)

	s := NewStateEndGame()

	s.Do(g)

	expectedMajority := worth.MajorityHolderBonusFirst + majority*worth.PricePerStock
	expectedMinority := worth.MajorityHolderBonusSecond + minority*worth.PricePerStock

	if p1.funds != expectedMajority {
		t.Errorf("p1 should have majority total of $%d, but has $%d", expectedMajority, p1.funds)
	}

	if p2.funds != expectedMinority {
		t.Errorf("p2 should have minority total of $%d, but has $%d", expectedMinority, p2.funds)
	}
}
