package cards

import "fmt"

// Create a new type of 'Deck'
// which is a slice of strings
type Deck []string

// Prints each card of deck in one line
func (d Deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
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

