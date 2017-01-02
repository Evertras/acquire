package acquire

// State is the state of the game that requires a specific decision from a Player
type State int

const (
	// StatePlayTile occurs when a tile needs to be played
	StatePlayTile int = iota

	// StateBuy occurs when stocks can be bought
	StateBuy

	// StateMerge Occurs when a hotel merge begins, possibly requiring a choice of which
	// chain will acquire which if multiple chains are equal size
	StateMerge

	// StateCreate occurs when a new hotel chain is created and the particular
	// chain must be chosen from available chains
	StateCreate

	// StateSell occurs when a merge is complete and each player that holds stock
	// must choose to sell, trade, or hold their shares
	StateSell

	// StateEnd occurs when the game is over
	StateEnd
)
