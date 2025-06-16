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

func RandomCard() *Card {
	return AllCards[rand.Intn(52)]
}

// n must be < 52
func RandomCards(n int) []*Card {
	inds := rand.Perm(52)[:n]
	result := make([]*Card, n)
	for i := range n {
		result[i] = AllCards[inds[i]]
	}
	return result
}
