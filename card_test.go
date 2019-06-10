package cards

import (
	"strconv"
	"testing"
)

func Test_Value_A5(t *testing.T) {
	var hand Hand

	hand = append(hand, Card{val: "A"})
	hand = append(hand, Card{val: "5"})

	val := hand.Value()

	if val != 16 {
		t.Error("A and 5 should return 16. Returned: " + strconv.Itoa(val))
	}
}

func Test_Value_TwoAces(t *testing.T) {
	var hand Hand

	hand = append(hand, Card{val: "A"})
	hand = append(hand, Card{val: "A"})
	hand = append(hand, Card{val: "5"})

	val := hand.Value()

	if val != 17 {
		t.Error("A,A and 5 should return 17. Returned: " + strconv.Itoa(val))
	}
}

func Test_Value_TwoAcesAndKing(t *testing.T) {
	var hand Hand

	hand = append(hand, Card{val: "A"})
	hand = append(hand, Card{val: "A"})
	hand = append(hand, Card{val: "K"})
	hand = append(hand, Card{val: "5"})

	val := hand.Value()

	if val != 17 {
		t.Error("A,A,K and 5 should return 17. Returned: " + strconv.Itoa(val))
	}
}

func Test_Value_TwoAcesAndTwoKing(t *testing.T) {
	var hand Hand

	hand = append(hand, Card{val: "A"})
	hand = append(hand, Card{val: "A"})
	hand = append(hand, Card{val: "K"})
	hand = append(hand, Card{val: "K"})

	val := hand.Value()

	if val != 22 {
		t.Error("A,A,K,K should return 22. Returned: " + strconv.Itoa(val))
	}
}
