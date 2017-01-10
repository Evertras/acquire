package acquire

// State is a single phase of decision-making within the Game that can affect
// the Game's state and returns the next State that the game is now in
type State interface {
	Do(g *Game) State
}
