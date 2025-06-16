package card

// Rank represents a card rank in a standard deck of playing cards.
type Rank struct {
	value  int    `json:"value"`  // Numeric value (1 for Ace, 2-10, 11 for Jack, 12 for Queen, 13 for King)
	symbol string `json:"symbol"` // Display symbol (e.g., "A", "2", ..., "J", "Q", "K")
}

// allRanks defines all possible ranks in a standard deck.
var AllRanks = [13]*Rank{
	{value: 1, symbol: "A"},
	{value: 2, symbol: "2"},
	{value: 3, symbol: "3"},
	{value: 4, symbol: "4"},
	{value: 5, symbol: "5"},
	{value: 6, symbol: "6"},
	{value: 7, symbol: "7"},
	{value: 8, symbol: "8"},
	{value: 9, symbol: "9"},
	{value: 10, symbol: "T"},
	{value: 11, symbol: "J"},
	{value: 12, symbol: "Q"},
	{value: 13, symbol: "K"},
}

// Getter methods for Rank
func (r *Rank) Value() int {
	return r.value
}

func (r *Rank) Symbol() string {
	return r.symbol
}

func RankOfValue(value int) *Rank {
	return AllRanks[value-1]
}
