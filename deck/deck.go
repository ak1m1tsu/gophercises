package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Shuffle the Deck
func (deck Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range deck {
		newPosition := r.Intn(len(deck) - 1)

		deck[i], deck[newPosition] = deck[newPosition], deck[i]
	}
}

// Saves a deck to the file
func (deck Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(
		filename,
		[]byte(deck.ToString()),
		0666,
	)
}

func (deck Deck) ToString() string {
	return strings.Join([]string(deck), sep)
}

// Returns a new instance of Deck type
func New() Deck {
	cardSuits := getCardSuits()
	cardValues := getCardValues()
	return generateDeck(cardSuits, cardValues)
}

// Read file and returns a instance of 'Deck'
func NewFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR |", err)
		os.Exit(1)
	}
	data := strings.Split(string(bs), sep)
	return Deck(data)
}

// Generates a new instance of Deck with card suits and values
func generateDeck(cardSuits, cardValues []string) Deck {
	cards := Deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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
