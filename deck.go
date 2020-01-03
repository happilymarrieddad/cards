package cards

import (
	"math/rand"
	"time"
)

// Deck - a deck of cards
type Deck struct {
	cards                []Card
	numberOfInitialCards int
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
	d.numberOfInitialCards = 0
	for i := 0; i < numOfDecks; i++ {

		numberOfPossibleCardsSetsInDeck := 4
		fullDeck := []Card{}
		for j := 0; j < numberOfPossibleCardsSetsInDeck; j++ {
			for k := 0; k < len(possibleCards); k++ {
				fullDeck = append(fullDeck, possibleCards[k])
			}
		}

		for j := 0; j < len(fullDeck); j++ {
			d.numberOfInitialCards++
			d.cards = append(d.cards, fullDeck[j])
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

func (d *Deck) GetCard() (x Card) {
	x, d.cards = d.cards[0], d.cards[1:]
	return
}

func (d *Deck) GetNumberOfCards() int {
	return len(d.cards)
}

func (d *Deck) GetCopyOfDeck() []Card {
	return d.cards
}

// PercentageLeftInDeck - return the percentage of the deck that is left
func (d *Deck) PercentageLeftInDeck() float32 {
	return (float32(len(d.cards)) / float32(d.numberOfInitialCards)) * 100
}
