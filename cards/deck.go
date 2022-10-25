package cards

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const sep = ";"

// Create a new type of 'Deck'
// which is a slice of strings
type Deck []string

// Prints each card of deck in one line
func (deck Deck) Print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}

func (deck Deck) ToString() string {
	return strings.Join([]string(deck), sep)
}

// Returns a new instance of Deck type
func NewDeck() Deck {
	cardSuits := getCardSuits()
	cardValues := getCardValues()
	return generateDeck(cardSuits, cardValues)
}

// Generates a new instance of Deck with card suits and values
func generateDeck(cardSuits, cardValues []string) Deck {
	cards := Deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// Returns a cards suits
func getCardSuits() []string {
	return []string{"Spades", "Hearts", "Diamonds", "Clubs"}
}

// Returns a cards values
func getCardValues() []string {
	return []string{"Ace", "Two", "Three", "Four"}
}

// Returns a cards for hand and remaining cards
func Deal(deck Deck, handSize int) (Deck, Deck) {
	return deck[:handSize], deck[handSize:]
}

// Saves a deck to the file
func (deck Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(
		filename,
		[]byte(deck.ToString()),
		0666,
	)
}
