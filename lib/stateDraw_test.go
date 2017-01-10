package acquire

import (
	"math/rand"
	"testing"
)

func TestStateDrawGivesSinglePiece(t *testing.T) {
	s := NewStateDraw()
	r := rand.New(rand.NewSource(0))
	p := NewPlayerRandom(r)
	g := NewGame(r, []Player{p})

	heldBefore := len(p.piecesHeld)
	s.Do(g)
	heldAfter := len(p.piecesHeld)

	if heldAfter != heldBefore+1 {
		t.Errorf("Should have %d pieces but have %d", heldBefore+1, heldAfter)
	}
}
