package cards

import (
	"testing"
)

func Test_Deck_Build(t *testing.T) {
	var d Deck

	d.Build(1)

	d.Shuffle()
}
func Test_Deck_BuildMoreThan1(t *testing.T) {
	var d Deck

	d.Build(5)

	d.Shuffle()
}
func Test_Deck_Build_Test_No_Decks(t *testing.T) {
	var d Deck

	d.Build(0)
}

func Test_Deck_PercentageLeft(t *testing.T) {
	var d Deck

	d.Build(1)
	d.Shuffle()

	deck := d.GetCopyOfDeck()

	numberOf2s := 0
	numberOf3s := 0
	numberOf4s := 0
	numberOf5s := 0
	numberOf6s := 0
	numberOf7s := 0
	numberOf8s := 0
	numberOf9s := 0
	numberOf10s := 0
	numberOfJs := 0
	numberOfQs := 0
	numberOfKs := 0
	numberOfAs := 0

	for _, c := range deck {
		switch c.val {
		case "2":
			numberOf2s++
		case "3":
			numberOf3s++
		case "4":
			numberOf4s++
		case "5":
			numberOf5s++
		case "6":
			numberOf6s++
		case "7":
			numberOf7s++
		case "8":
			numberOf8s++
		case "9":
			numberOf9s++
		case "10":
			numberOf10s++
		case "J":
			numberOfJs++
		case "Q":
			numberOfQs++
		case "K":
			numberOfKs++
		case "A":
			numberOfAs++
		}
	}

	if numberOf2s != 4 {
		t.Error("numberOf2s not equal to 4")
		return
	}
	if numberOf3s != 4 {
		t.Error("numberOf3s not equal to 4")
		return
	}
	if numberOf4s != 4 {
		t.Error("numberOf4s not equal to 4")
		return
	}
	if numberOf5s != 4 {
		t.Error("numberOf5s not equal to 4")
		return
	}
	if numberOf6s != 4 {
		t.Error("numberOf6s not equal to 4")
		return
	}
	if numberOf7s != 4 {
		t.Error("numberOf7s not equal to 4")
		return
	}
	if numberOf8s != 4 {
		t.Error("numberOf8s not equal to 4")
		return
	}
	if numberOf9s != 4 {
		t.Error("numberOf9s not equal to 4")
		return
	}
	if numberOf10s != 4 {
		t.Error("numberOf10s not equal to 4")
		return
	}
	if numberOfJs != 4 {
		t.Error("numberOfJs not equal to 4")
		return
	}
	if numberOfQs != 4 {
		t.Error("numberOfQs not equal to 4")
		return
	}
	if numberOfKs != 4 {
		t.Error("numberOfKs not equal to 4")
		return
	}
	if numberOfAs != 4 {
		t.Error("numberOfAs not equal to 4")
		return
	}

	// Ignore the returned card
	d.GetCard()

	if float32(98.07692) != d.PercentageLeftInDeck() {
		t.Error("bad percentage left")
		return
	}

	if 51 != d.GetNumberOfCards() {
		t.Error("bad number of cards left")
		return
	}
}
