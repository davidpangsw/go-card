package card

// Suit represents a card suit in a standard deck of playing cards.
type Suit struct {
	value         int    `json:"value"`          // Numeric value (1 for Clubs, 2 for Diamonds, 3 for Hearts, 4 for Spades)
	symbol        string `json:"symbol"`         // Short symbol (e.g., "C", "D", "H", "S")
	displaySymbol string `json:"display_symbol"` // Unicode symbol (e.g., "♣", "♦", "♥", "♠")
}

// Getter methods for Suit
func (s *Suit) Value() int {
	return s.value
}

func (s *Suit) Symbol() string {
	return s.symbol
}

func (s *Suit) DisplaySymbol() string {
	return s.displaySymbol
}

func SuitOfValue(value int) *Suit {
	return AllSuits[value-1]
}

// allSuits defines all possible suits in a standard deck.
var AllSuits = [4]*Suit{
	{value: 1, symbol: "C", displaySymbol: "♣"},
	{value: 2, symbol: "D", displaySymbol: "♦"},
	{value: 3, symbol: "H", displaySymbol: "♥"},
	{value: 4, symbol: "S", displaySymbol: "♠"},
}
