package acquire

import "testing"

func TestInitials(t *testing.T) {
	cases := []struct {
		h Hotel
		i byte
	}{
		{Tower, 'T'},
		{Luxor, 'L'},
		{American, 'A'},
		{Worldwide, 'W'},
	}

	for _, c := range cases {
		i := GetHotelInitial(c.h)

		if i != c.i {
			t.Errorf("Incorrect initials: Testing %v, got %v", c, i)
		}
	}
}

type worthCase struct {
	h Hotel
	s int
	w HotelWorth
}

func TestWorth(t *testing.T) {
	cases := []worthCase{}

	tier1 := []Hotel{Tower, Luxor}
	tier2 := []Hotel{American, Worldwide, Festival}
	tier3 := []Hotel{Imperial, Continental}

	for s := 2; s < 45; s++ {
		var offset int

		if s < 6 {
			offset = s - 2
		} else if s <= 10 {
			offset = 4
		} else if s <= 20 {
			offset = 5
		} else if s <= 30 {
			offset = 6
		} else if s <= 40 {
			offset = 7
		} else {
			offset = 8
		}

		t1Worth := HotelWorth{offset*100 + 200, offset*1000 + 2000, offset*500 + 1000}
		t2Worth := HotelWorth{t1Worth.PricePerStock + 100, t1Worth.MajorityHolderBonusFirst + 1000, t1Worth.MajorityHolderBonusSecond + 500}
		t3Worth := HotelWorth{t1Worth.PricePerStock + 200, t1Worth.MajorityHolderBonusFirst + 2000, t1Worth.MajorityHolderBonusSecond + 1000}

		for _, h := range tier1 {
			cases = append(cases, worthCase{h, s, t1Worth})
		}

		for _, h := range tier2 {
			cases = append(cases, worthCase{h, s, t2Worth})
		}

		for _, h := range tier3 {
			cases = append(cases, worthCase{h, s, t3Worth})
		}
	}

	for _, c := range cases {
		worth := GetWorth(c.h, c.s)

		if worth.PricePerStock != c.w.PricePerStock {
			t.Errorf("Worths not equal: Testing against %v and testing against %v", c, worth)
		}
	}
}
