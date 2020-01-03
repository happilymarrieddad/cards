package cards

// Hand - the hand of cards
type Hand []Card

// HasAce - does this hand have an ace
func (h *Hand) HasAce() bool {
	for _, c := range *h {
		if c.val == "A" {
			return true
		}
	}
	return false
}

// HasFace - does this hand have a face card
func (h *Hand) HasFace() bool {
	for _, c := range *h {
		if c.val == "K" || c.val == "Q" || c.val == "J" || c.val == "10" {
			return true
		}
	}
	return false
}

// Value - get the current value of the hand
func (h *Hand) Value() int {
	var val int
	var numAces int
	for _, c := range *h {
		val += c.Value()
		if c.val == "A" {
			numAces++
		}
	}

	// This corrects for Aces being used to count over 21
	for i := 0; i < numAces; i++ {
		if val > 21 {
			val -= 10
		}
	}

	return val
}

// Specific Blackjack Functions
func (h *Hand) HasBlackjack() bool {
	if len(*h) != 2 {
		return false
	}

	if h.HasAce() && h.HasFace() {
		return true
	}

	return false
}

func (h *Hand) Display() (display string) {
	for _, c := range *h {
		display += c.val
	}

	return
}
