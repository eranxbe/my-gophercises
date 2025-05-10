//go:generate stringer -type=Suit,Rank
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

var suits = [...]Suit{Spade, Diamond,Club,Heart}

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

func New(opts ...func([]Card) []Card) []Card {
	var deck []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		} 
	}
	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck
}

func Less(deck []Card) func(i, j int) bool { // implement the 'less' function of sort package
	return func(i, j int) bool {
		return absRank(deck[i]) < absRank(deck[j])
	}
}

func DefaultSort(deck []Card) []Card { // sorts by ace:spades, two:spades...
	sort.Slice(deck, Less(deck))
	return deck
}

func Sort(less func(deck []Card) func(i, j int) bool) func([]Card) []Card {
	return func(deck []Card) []Card {
		sort.Slice(deck, less(deck))
		return deck
	}
}

func absRank(c Card) int { // get a distinct int value for each card, in ranges of the suit
	return int(maxRank) * int(c.Suit) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(deck []Card) []Card {
	ret := make([]Card, len(deck))
	perm := shuffleRand.Perm(len(deck))
	for i, j := range perm {
		ret[i] = deck[j]
	}
	return ret
}

func Jokers(n int) func([]Card) []Card {
	return func(deck []Card) []Card {
		for i := range n {
			deck = append(deck, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return deck
	}
}

func Filter(f func(c Card) bool) func([]Card) []Card { // build a new slice and fill it only with the relevant cards to the filter
	return func(deck []Card) []Card {
		var ret []Card
		for _, c := range deck {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	} 
}

func Deck(n int) func([]Card) []Card {
	return func(deck []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, deck...)
		}
		return ret
	}
}