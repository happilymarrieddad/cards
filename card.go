package cards

import "strconv"

var (
	PossibleCards = []Card{
		Card{"A"}, Card{"K"}, Card{"Q"}, Card{"J"}, Card{"10"}, Card{"9"}, Card{"8"},
		Card{"7"}, Card{"6"}, Card{"5"}, Card{"4"}, Card{"3"}, Card{"2"},
	}
)

type Card struct {
	val string
}

func (c *Card) Value() int {
	switch c.val {
	case "A":
		return 11
	case "K", "Q", "J", "10":
		return 10
	default:
		i, _ := strconv.Atoi(c.val)
		return i
	}
}
