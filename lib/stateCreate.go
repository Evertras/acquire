package acquire

// StateCreate is the state that triggers when a new hotel chain needs to be
// chosen and then created
type StateCreate struct {
	TriggeringPiece Piece
	ActivePlayer    *Player
}

// NewStateCreate creates a new StateCreate instance
func NewStateCreate(activePlayer *Player, triggeringPiece Piece) StateCreate {
	return StateCreate{triggeringPiece, activePlayer}
}

// Do creates a new hotel chain and proceeds with the next state
func (s StateCreate) Do(g *Game) State {
	selected := (*s.ActivePlayer).Create(g, s.TriggeringPiece)
	g.CurrentChainSizes[selected] = g.Board.Fill(s.TriggeringPiece, selected)

	if g.AvailableStocks[selected] > 0 {
		g.AvailableStocks[selected]--
		(*s.ActivePlayer).GiveStocks(selected, 1)
	}

	l := len(g.AvailableChains)

	if l == 1 || g.AvailableChains[l-1] == selected {
		g.AvailableChains = g.AvailableChains[:l-1]
	} else {
		for i := 0; i < l-1; i++ {
			if g.AvailableChains[i] == selected {
				g.AvailableChains[i] = g.AvailableChains[l-1]
				g.AvailableChains = g.AvailableChains[:l-1]
				break
			}
		}
	}

	return NewStateBuy()
}
