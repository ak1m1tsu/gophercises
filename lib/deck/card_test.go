package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Clubs
	// Jack of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expectedCard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expectedCard {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
	if cards[len(cards)-1].Suit != Joker {
		t.Error("Expected Joker as last card. Received:", cards[len(cards)-1])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expectedCard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expectedCard {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers. Received:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 ranks * 4 suits * 3 decks
	expectedLength := 13 * 4 * 3
	if len(cards) != expectedLength {
		t.Errorf("Expected %d cards. Received: %d.", expectedLength, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40 35 ...]
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s. Received: %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first card to be %s. Received: %s", second, cards[1])
	}
}
