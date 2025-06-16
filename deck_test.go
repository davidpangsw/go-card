package card

import (
	"math/rand"
	"testing"
)

func TestShuffle(t *testing.T) {
	d := NewDeck()
	d.ShuffleWithRand(rand.New(rand.NewSource(0)))

	// check if the cards are 52 distinct cards
	if len(d.Cards) != 52 {
		t.Fatal("Deck is not 52-card.", len(d.Cards), d)
	}
	cards := make(map[string]bool)
	for i := range 52 {
		s := d.Cards[i].String()
		if _, exists := cards[s]; exists {
			t.Fatal("Repeated card", s, d)
		}
		cards[s] = true
	}
}
