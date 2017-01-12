package acquire

import (
	"math/rand"
	"testing"
)

func TestStateSellSplitsFundsCorrectly(t *testing.T) {
	cases := []struct {
		Total    int
		Split    int
		Expected int
	}{
		{1000, 2, 500},
		{1000, 3, 400},
		{5000, 5, 1000},
		{100, 2, 100},
	}

	for _, c := range cases {
		res := getFundSplit(c.Total, c.Split)

		if res != c.Expected {
			t.Errorf("%d / %d = %d but should be %d", c.Total, c.Split, res, c.Expected)
		}
	}
}

func TestStateSellGivesMajorityBonusesSimple(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	p3 := NewPlayerRandom(r)
	h := HotelLuxor

	worth := HotelWorth{
		// Set to 0 so randomly sold stock doesn't add funds
		PricePerStock:             0,
		MajorityHolderBonusFirst:  2000,
		MajorityHolderBonusSecond: 1000,
	}

	defunct := make(map[Hotel]HotelWorth)

	defunct[h] = worth

	p1.stocksOwned[h] = 2
	p2.stocksOwned[h] = 4
	p3.stocksOwned[h] = 3

	p1.funds = 0
	p2.funds = 0
	p3.funds = 0

	ps := []Player{p1, p2, p3}

	g := NewGame(r, ps)

	s := NewStateSell(defunct, HotelAmerican)

	s.Do(g)

	if p1.funds != 0 {
		t.Errorf("p1 should have $0, but has $%d", p1.funds)
	}

	if p2.funds != worth.MajorityHolderBonusFirst {
		t.Errorf("p2 should have $%d, but has $%d", worth.MajorityHolderBonusFirst, p2.funds)
	}

	if p3.funds != worth.MajorityHolderBonusSecond {
		t.Errorf("p3 should have $%d, but has $%d", worth.MajorityHolderBonusSecond, p3.funds)
	}
}

func TestStateSellGivesAllBonusToSingleHolder(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	p3 := NewPlayerRandom(r)
	h := HotelLuxor

	worth := HotelWorth{
		// Set to 0 so randomly sold stock doesn't add funds
		PricePerStock:             0,
		MajorityHolderBonusFirst:  2000,
		MajorityHolderBonusSecond: 1000,
	}

	defunct := make(map[Hotel]HotelWorth)

	defunct[h] = worth

	p1.stocksOwned[h] = 2

	p1.funds = 0
	p2.funds = 0
	p3.funds = 0

	ps := []Player{p1, p2, p3}

	g := NewGame(r, ps)

	s := NewStateSell(defunct, HotelAmerican)

	s.Do(g)

	expected := worth.MajorityHolderBonusFirst + worth.MajorityHolderBonusSecond

	if p1.funds != expected {
		t.Errorf("p1 should have $%d, but has $%d", expected, p1.funds)
	}

	if p2.funds != 0 {
		t.Errorf("p2 should have $%d, but has $%d", 0, p2.funds)
	}

	if p3.funds != 0 {
		t.Errorf("p3 should have $%d, but has $%d", 0, p3.funds)
	}
}

func TestStateSellSplitsBonusToTiedPlayers(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	p3 := NewPlayerRandom(r)
	h := HotelLuxor

	worth := HotelWorth{
		// Set to 0 so randomly sold stock doesn't add funds
		PricePerStock:             0,
		MajorityHolderBonusFirst:  2000,
		MajorityHolderBonusSecond: 1000,
	}

	defunct := make(map[Hotel]HotelWorth)

	defunct[h] = worth

	p1.stocksOwned[h] = 3
	p2.stocksOwned[h] = 3
	p3.stocksOwned[h] = 2

	p1.funds = 0
	p2.funds = 0
	p3.funds = 0

	ps := []Player{p1, p2, p3}

	g := NewGame(r, ps)

	s := NewStateSell(defunct, HotelAmerican)

	s.Do(g)

	expected := getFundSplit(worth.MajorityHolderBonusFirst+worth.MajorityHolderBonusSecond, 2)

	if p1.funds != expected {
		t.Errorf("p1 should have $%d, but has $%d", expected, p1.funds)
	}

	if p2.funds != expected {
		t.Errorf("p2 should have $%d, but has $%d", expected, p2.funds)
	}

	if p3.funds != 0 {
		t.Errorf("p3 should have $%d, but has $%d", 0, p3.funds)
	}
}

func TestStateSellSplitsSecondaryBonusToTiedPlayers(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	p3 := NewPlayerRandom(r)
	h := HotelLuxor

	worth := HotelWorth{
		// Set to 0 so randomly sold stock doesn't add funds
		PricePerStock:             0,
		MajorityHolderBonusFirst:  2000,
		MajorityHolderBonusSecond: 1000,
	}

	defunct := make(map[Hotel]HotelWorth)

	defunct[h] = worth

	p1.stocksOwned[h] = 3
	p2.stocksOwned[h] = 2
	p3.stocksOwned[h] = 2

	p1.funds = 0
	p2.funds = 0
	p3.funds = 0

	ps := []Player{p1, p2, p3}

	g := NewGame(r, ps)

	s := NewStateSell(defunct, HotelAmerican)

	s.Do(g)

	expected := getFundSplit(worth.MajorityHolderBonusSecond, 2)

	if p1.funds != worth.MajorityHolderBonusFirst {
		t.Errorf("p1 should have $%d, but has $%d", worth.MajorityHolderBonusFirst, p1.funds)
	}

	if p2.funds != expected {
		t.Errorf("p2 should have $%d, but has $%d", expected, p2.funds)
	}

	if p3.funds != expected {
		t.Errorf("p3 should have $%d, but has $%d", expected, p3.funds)
	}
}
