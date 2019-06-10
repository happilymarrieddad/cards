package cards

import "testing"

func Test_Hand_HasAce_Success(t *testing.T) {
	hand := Hand{
		Card{"A"},
	}

	if !hand.HasAce() {
		t.Error("A hand with an ace should have one")
	}
}

func Test_Hand_HasAce_Failure(t *testing.T) {
	hand := Hand{
		Card{"K"},
	}

	if hand.HasAce() {
		t.Error("This hand does not have an ace and shouldn't show that it should")
	}
}

func Test_Hand_HasFace_Success(t *testing.T) {
	hand := Hand{
		Card{"K"},
	}

	if !hand.HasFace() {
		t.Error("A hand with a face should have one")
	}
}

func Test_Hand_HasFace_Failure(t *testing.T) {
	hand := Hand{
		Card{"2"},
	}

	if hand.HasFace() {
		t.Error("This hand does not have a face and shouldn't show that it should")
	}
}

func Test_Hand_HasBlackjack_Success(t *testing.T) {
	hand := Hand{
		Card{"A"},
		Card{"J"},
	}

	if !hand.HasBlackjack() {
		t.Error("This hand does have a blackjack")
	}
}

func Test_Hand_HasBlackjack_Failure_MoreThan2Cards(t *testing.T) {
	hand := Hand{
		Card{"A"},
		Card{"J"},
		Card{"J"},
	}

	if hand.HasBlackjack() {
		t.Error("This hand does have a blackjack but it has more than 2 cards")
	}
}

func Test_Hand_HasBlackjack_Failure(t *testing.T) {
	hand := Hand{
		Card{"3"},
		Card{"5"},
	}

	if hand.HasBlackjack() {
		t.Error("This hand does not have a blackjack")
	}
}
