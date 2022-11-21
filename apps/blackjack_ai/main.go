package main

import (
	"flag"
	"fmt"

	"github.com/romankravchuk/learn-go/apps/blackjack_ai/blackjack"
)

var (
	numDecks        int
	numHands        int
	blackjackPayout float64
)

func init() {
	flag.IntVar(&numDecks, "decks", 3, "The number of decks.")
	flag.IntVar(&numHands, "hands", 2, "The number of hands.")
	flag.Float64Var(&blackjackPayout, "payout", 1.5, "The blackjack payout.")
}

func main() {
	flag.Parse()
	opts := blackjack.Options{
		Decks:           numDecks,
		Hands:           numHands,
		BlackjackPayout: blackjackPayout,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
