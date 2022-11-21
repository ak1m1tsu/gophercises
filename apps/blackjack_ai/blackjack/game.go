package blackjack

import (
	"errors"
	"fmt"

	"github.com/romankravchuk/learn-go/lib/deck"
)

type state int8

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

func New(opts Options) Game {
	game := Game{state: statePlayerTurn, balance: 0}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 100
	}
	if opts.BlackjackPayout == 0.0 {
		opts.BlackjackPayout = 1.5
	}
	game.numDecks = opts.Decks
	game.numHands = opts.Hands
	game.blackjackPayout = opts.BlackjackPayout
	return game
}

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type Game struct {
	// unexported fields
	numDecks        int
	numHands        int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []hand
	handIdx   int
	playerBet int
	balance   int

	dealer []deck.Card
}

func bet(game *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	if bet < 100 {
		panic("bet must be at least 100")
	}
	game.playerBet = bet
}

func deal(game *Game) {
	playerHand := make([]deck.Card, 0, 5)
	game.handIdx = 0
	game.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, game.deck = draw(game.deck)
		playerHand = append(playerHand, card)
		card, game.deck = draw(game.deck)
		game.dealer = append(game.dealer, card)
	}
	game.player = []hand{
		{
			cards: playerHand,
			bet:   game.playerBet,
		},
	}
	game.state = statePlayerTurn
}

func (game *Game) Play(ai AI) int {
	// CardsCount * GameDecksCount / DecksToReshuffleCount
	minLengthToReshuffle := 52 * game.numDecks / 3
	for i := 0; i < game.numHands; i++ {
		shuffled := false
		if len(game.deck) < minLengthToReshuffle {
			game.deck = deck.New(deck.Deck(game.numDecks), deck.Shuffle)
			shuffled = true
		}
		bet(game, ai, shuffled)
		deal(game)
		if Blackjack(game.dealer...) {
			endRound(game, ai)
			continue
		}
		for game.state == statePlayerTurn {
			hand := make([]deck.Card, len(*game.currentHand()))
			copy(hand, *game.currentHand())
			move := ai.Play(hand, game.dealer[0])
			err := move(game)
			switch err {
			case errBust:
				MoveStand(game)
			case nil:
				continue
			default:
				panic(err)
			}
		}
		for game.state == stateDealerTurn {
			dealerScore := Score(game.dealer...)
			if dealerScore < 16 || (dealerScore == 17 && Soft(game.dealer...)) {
				MoveHit(game)
			} else {
				MoveStand(game)
			}
		}
		endRound(game, ai)
	}
	return game.balance
}

var (
	errBust = errors.New("hand score exceeded 21")
)

type Move func(*Game) error

func MoveHit(game *Game) error {
	hand := game.currentHand()
	var card deck.Card
	card, game.deck = draw(game.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		return errBust
	}
	return nil
}

func MoveSplit(game *Game) error {
	cards := game.currentHand()
	if len(*cards) != 2 {
		return errors.New("you can only split with two cards in your hands")
	}
	if (*cards)[0].Rank != (*cards)[1].Rank {
		return errors.New("both cards must have the same rank to split")
	}
	game.player = append(game.player, hand{
		cards: []deck.Card{(*cards)[1]},
		bet:   game.player[game.handIdx].bet,
	})
	game.player[game.handIdx].cards = (*cards)[:1]
	return nil
}

func MoveDouble(game *Game) error {
	if len(*game.currentHand()) != 2 {
		return errors.New("can only double on a hand with 2 cards")
	}
	game.playerBet *= 2
	MoveHit(game)
	return MoveStand(game)
}

func (game *Game) currentHand() *[]deck.Card {
	switch game.state {
	case statePlayerTurn:
		return &game.player[game.handIdx].cards
	case stateDealerTurn:
		return &game.dealer
	default:
		panic("It isn't currently any player's turn")
	}
}

type hand struct {
	cards []deck.Card
	bet   int
}

func MoveStand(game *Game) error {
	if game.state == stateDealerTurn {
		game.state++
		return nil
	}
	if game.state == statePlayerTurn {
		game.handIdx++
		if game.handIdx >= len(game.player) {
			game.state++
		}
		return nil
	}
	return errors.New("invalid state")
}

func endRound(game *Game, ai AI) {
	dealerScore := Score(game.dealer...)
	dealerBlackjack := Blackjack(game.dealer...)
	allHands := make([][]deck.Card, len(game.player))
	for hi, hand := range game.player {
		allHands[hi] = hand.cards
		cards := hand.cards
		playerScore, playerBlackjack := Score(cards...), Blackjack(cards...)
		winnings := hand.bet
		switch {
		case (playerBlackjack && dealerBlackjack) || playerScore == dealerScore:
			winnings = 0
		case dealerBlackjack || playerScore > 21 || dealerScore > playerScore:
			winnings = -winnings
		case playerBlackjack:
			winnings = int(float64(winnings) * game.blackjackPayout)
		}
		game.balance += winnings
	}
	ai.Results(allHands, game.dealer)
	fmt.Println()
	game.player = nil
	game.dealer = nil
}

func draw(hand []deck.Card) (deck.Card, []deck.Card) {
	return hand[0], hand[1:]
}

// Returns the score of a hand
func Score(hand ...deck.Card) int {
	minScore := minScore(hand)
	if minScore > 11 {
		return minScore
	}
	for _, card := range hand {
		if card.Rank == deck.Ace {
			// Ace is currently worth 1,
			// and we are changing it to be worth 11
			// 11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

// Returns true if a hand is a blackjack
func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
}

func minScore(hand []deck.Card) int {
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

// Returns true if the score of a hand is a soft score -
// that is if an ace is being counted as 11 points.
func Soft(hand ...deck.Card) bool {
	return Score(hand...) != minScore(hand)
}
