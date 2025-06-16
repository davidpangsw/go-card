package card

import (
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	// An array of cards (Can be any cards, even repeating)
	Cards []*Card
}

// create new deck with standard 52 cards
func NewDeck() *Deck {
	return &Deck{Cards: AllCards[:]}
}

func (d *Deck) Size() int {
	return len(d.Cards)
}

func (d *Deck) ShuffleWithRand(r *rand.Rand) {
	inds := r.Perm(d.Size())
	for i, ind := range inds {
		temp := d.Cards[i]
		d.Cards[i] = d.Cards[ind]
		d.Cards[ind] = temp
	}
}

// use time.Now().UnixNano() as seed
func (d *Deck) Shuffle() {
	d.ShuffleWithRand(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func (d *Deck) String() string {
	var b strings.Builder
	b.WriteString("[")
	for _, c := range d.Cards {
		b.WriteString(c.String())
		b.WriteString(",")
	}
	b.WriteString("]")
	return b.String()
}

// Draw the top card Without replacement
// Caller should check the deck size before calling
func (d *Deck) Draw() *Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}
