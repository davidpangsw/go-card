package card

import "math/rand"

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

// Without replacement
// Caller should check the deck size before calling
func (d *Deck) Draw() *Card {
	index := rand.Intn(len(d.Cards))
	card := d.Cards[index]
	d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
	return card
}
