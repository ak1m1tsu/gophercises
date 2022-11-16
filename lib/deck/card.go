package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// Represents an option for func New
type NewOption func([]Card) []Card

// Returns a slice of Cards
func New(opts ...NewOption) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{suit, rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// Option that sort deck
//
// The first card is always "Ace of Spades"
//
// The last card is "King of Hearts"
//
// If Joker option is used then a "Joker"
// is a last card
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Option that sort deck
func Sort(less func(cards []Card) func(i, j int) bool) NewOption {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Sorts deck
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return orderNumber(cards[i]) < orderNumber(cards[j])
	}
}

func orderNumber(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Option that shuffle deck
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	for i, j := range shuffleRand.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}

// Options that adds a N "Joker" to deck
func Jokers(n int) NewOption {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}

// Option that filter out cards in deck
func Filter(filter func(card Card) bool) NewOption {
	return func(cards []Card) []Card {
		var ret []Card
		for _, card := range cards {
			if !filter(card) {
				ret = append(ret, card)
			}
		}
		return ret
	}
}

// Option that return a N times deck copy
func Deck(n int) NewOption {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
