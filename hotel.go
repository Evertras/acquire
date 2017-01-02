package acquire

// Hotel represents the name of hotel chain
type Hotel int

const (
	// Tower hotels
	Tower Hotel = iota

	// Luxor hotels
	Luxor

	// American hotels
	American

	// Worldwide hotels
	Worldwide

	// Festival hotels
	Festival

	// Imperial hotels
	Imperial

	// Continental hotels
	Continental
)

var hotelInitials = [11]byte{
	'T',
	'L',
	'A',
	'W',
	'F',
	'I',
	'C',
}

// GetHotelInitial gets the initial for the given hotel chain
func GetHotelInitial(h Hotel) byte {
	return hotelInitials[h]
}

// HotelWorth contains information on how much a hotel's stock is worth as well
// as majority bonuses for that hotel.
type HotelWorth struct {
	PricePerStock             int
	MajorityHolderBonusFirst  int
	MajorityHolderBonusSecond int
}

var hotelWorths = [11]HotelWorth{
	HotelWorth{200, 2000, 1000},
	HotelWorth{300, 3000, 1500},
	HotelWorth{400, 4000, 2000},
	HotelWorth{500, 5000, 2500},
	HotelWorth{600, 6000, 3000},
	HotelWorth{700, 7000, 3500},
	HotelWorth{800, 8000, 4000},
	HotelWorth{900, 9000, 4500},
	HotelWorth{1000, 10000, 5000},
	HotelWorth{1100, 11000, 5500},
	HotelWorth{1200, 12000, 6000},
}

// GetWorth returns the hotel's worth given its current chain size
func GetWorth(hotel Hotel, chainSize int) HotelWorth {
	offset := 0

	if hotel == American || hotel == Worldwide || hotel == Festival {
		offset = 1
	} else if hotel == Imperial || hotel == Continental {
		offset = 2
	}

	var bucket int

	if chainSize == 2 {
		bucket = 0
	} else if chainSize == 3 {
		bucket = 1
	} else if chainSize == 4 {
		bucket = 2
	} else if chainSize == 5 {
		bucket = 3
	} else if chainSize <= 10 {
		bucket = 4
	} else if chainSize <= 20 {
		bucket = 5
	} else if chainSize <= 30 {
		bucket = 6
	} else if chainSize <= 40 {
		bucket = 7
	} else {
		bucket = 8
	}

	return hotelWorths[bucket+offset]
}
