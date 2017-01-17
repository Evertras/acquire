package acquire

import (
	"math/rand"
	"testing"
)

func TestStateEndTurnPassesToNextPlayer(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	players := []Player{p1, p2}
	g := NewGame(r, players)
	s := NewStateEndTurn()

	s.Do(g)

	if g.CurrentPlayerIndex != 1 {
		t.Errorf("Should have passed to next player, but have index %d", g.CurrentPlayerIndex)
	}
}

func TestStateEndTurnResetsToFirstPlayer(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	players := []Player{p1, p2}
	g := NewGame(r, players)
	s := NewStateEndTurn()
	g.CurrentPlayerIndex = len(players) - 1

	s.Do(g)

	if g.CurrentPlayerIndex != 0 {
		t.Errorf("Should have reset to index 0, but have index %d", g.CurrentPlayerIndex)
	}
}

func TestStateEndTurnEndsGameIfTwoChainsAreAt11(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	players := []Player{p1, p2}

	g := NewGame(r, players)

	g.CurrentChainSizes[HotelLuxor] = 11
	g.CurrentChainSizes[HotelAmerican] = 11

	s := NewStateEndTurn()

	next := s.Do(g)

	if next != nil {
		t.Error("Should have ended the game")
	}
}

func TestStateEndTurnContinuesGameWhenOneChainStillSmall(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	p1 := NewPlayerRandom(r)
	p2 := NewPlayerRandom(r)
	players := []Player{p1, p2}

	g := NewGame(r, players)

	g.CurrentChainSizes[HotelLuxor] = 10
	g.CurrentChainSizes[HotelAmerican] = 11

	s := NewStateEndTurn()

	next := s.Do(g)

	if next == nil {
		t.Error("Should not have ended the game")
	}
}
