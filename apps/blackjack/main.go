package main

import (
	"fmt"
	"strings"

	"github.com/romankravchuk/learn-go/lib/deck"
)

type Hand []deck.Card

func (hand Hand) String() string {
	strs := make([]string, len(hand))
	for i := range hand {
		strs[i] = hand[i].String()
	}
	return strings.Join(strs, ", ")
}

func (hand Hand) DealerString() string {
	return hand[0].String() + ", **HIDDEN**"
}

func (hand Hand) Score() int {
	minScore := hand.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, card := range hand {
		if card.Rank == deck.Ace {
			// Ace is currently worth 1, and we are changing it to be worth 11
			// 11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

func (hand Hand) MinScore() int {
	score := 0
	for _, card := range hand {
		score += min(int(card.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Shuffle(gameState GameState) GameState {
	ret := clone(gameState)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

func Deal(gameState GameState) GameState {
	ret := clone(gameState)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

func Hit(gameState GameState) GameState {
	ret := clone(gameState)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gameState GameState) GameState {
	ret := clone(gameState)
	ret.State++
	return ret
}

func EndHand(gameState GameState) GameState {
	ret := clone(gameState)
	playerScore, dealerScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", ret.Player, "\nScore:", playerScore)
	fmt.Println("Dealer:", ret.Dealer, "\nScore:", dealerScore)
	switch {
	case playerScore > 21:
		fmt.Println("You busted.")
	case dealerScore > 21:
		fmt.Println("Dealer busted")
	case playerScore > dealerScore:
		fmt.Println("You win!")
	case dealerScore > playerScore:
		fmt.Println("You lose")
	case playerScore == dealerScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gameState GameState
	gameState = Shuffle(gameState)
	gameState = Deal(gameState)

	var input string
	for gameState.State == StatePlayerTurn {
		fmt.Println("Player:", gameState.Player)
		fmt.Println("Dealer:", gameState.Dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gameState = Hit(gameState)
		case "s":
			gameState = Stand(gameState)
		default:
			fmt.Println("Invalid option:", input)
		}
	}

	for gameState.State == StateDealerTurn {
		if gameState.Dealer.Score() <= 16 || (gameState.Dealer.Score() == 17 && gameState.Dealer.MinScore() != 17) {
			gameState = Hit(gameState)
		} else {
			gameState = Stand(gameState)
		}
	}
	gameState = EndHand(gameState)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck []deck.Card
	State
	Player Hand
	Dealer Hand
}

func (gameState *GameState) CurrentPlayer() *Hand {
	switch gameState.State {
	case StatePlayerTurn:
		return &gameState.Player
	case StateDealerTurn:
		return &gameState.Dealer
	default:
		panic("It isn't currently any player's turn")
	}
}

func clone(gameState GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gameState.Deck)),
		State:  gameState.State,
		Player: make(Hand, len(gameState.Player)),
		Dealer: make(Hand, len(gameState.Dealer)),
	}
	copy(ret.Deck, gameState.Deck)
	copy(ret.Player, gameState.Player)
	copy(ret.Dealer, gameState.Dealer)
	return ret
}
