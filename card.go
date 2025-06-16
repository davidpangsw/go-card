package card

import "math/rand"

type Card struct {
	value int // 0-12, 13-25, 26-38, 39-52 (No Joker)
	rank  *Rank
	suit  *Suit
}

func (c *Card) Rank() *Rank {
	return c.rank
}

func (c *Card) Suit() *Suit {
	return c.suit
}

// Return string representation of a card
// e.g. "H2" for Heart 2, "ST" for Spade 10, "DA" for Diamond A
func (c *Card) String() string {
	return c.suit.symbol + c.rank.symbol
}

var AllCards = [52]*Card{}

func init() {
	for value := range 52 {
		card := &Card{value, AllRanks[value%13], AllSuits[value/13]}
		AllCards[value] = card
	}
}

// Randomly choose 1 card from a standard 52-deck
func RandomCard() *Card {
	return AllCards[rand.Intn(52)]
}

// Randomly choose n different cards from a standard 52-deck
// n must be < 52
func RandomCards(n int) []*Card {
	inds := rand.Perm(52)[:n]
	result := make([]*Card, n)
	for i := range n {
		result[i] = AllCards[inds[i]]
	}
	return result
}
