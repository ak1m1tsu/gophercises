package deck_test

import (
	"os"
	"testing"

	"github.com/romankravchuk/learn-go/deck"
)

func TestNew(t *testing.T) {
	deck := deck.New()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Exptected first card is 'Ace of Spades', but got '%v'", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Exptected first card is 'Four of Clubs', but got '%v'", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewFromFile(t *testing.T) {
	const file = "_decktesting"
	os.Remove(file)

	d := deck.New()
	d.SaveToFile(file)

	loadedDeck := deck.NewFromFile(file)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove(file)
}
