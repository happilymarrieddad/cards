package cards

import (
	"math/rand"
	"time"
)

// Deck - a deck of cards
type Deck struct {
	cards []Card
}

// Build - build the deck with a series of cards
func (d *Deck) Build(numOfDecks int, possibleCards ...Card) {
	if numOfDecks < 1 {
		numOfDecks = 1
	}

	if len(possibleCards) == 0 {
		possibleCards = PossibleCards
	}

	// Empty the current deck
	d.cards = []Card{}

	// Actually build the deck
	for i := 0; i < numOfDecks; i++ {
		for j := 0; j < len(possibleCards); j++ {
			d.cards = append(d.cards, possibleCards[j])
		}
	}

	d.Shuffle()
}

// Shuffle - shuffle
func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(d.cards); n > 0; n-- {
		randIndex := r.Intn(n)
		d.cards[n-1], d.cards[randIndex] = d.cards[randIndex], d.cards[n-1]
	}
}
