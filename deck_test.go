package cards

import "testing"

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
