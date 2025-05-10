package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Nine of Diamonds
	// King of Spades
	// Joker
}

func TestNew(t *testing.T) {
	deck := New()

	if len(deck) != 13*4 {
		t.Error("Wrong number of cards in deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(DefaultSort)
	expectedCard := Card{Suit: Spade, Rank: Ace}
	if deck[0] != expectedCard {
		t.Error("Expected 1st card to be Ace of Spades, got:", deck[0])
	}
}

func TestSort(t *testing.T) {
	deck := New(Sort(Less))
	expectedCard := Card{Suit: Spade, Rank: Ace}
	if deck[0] != expectedCard {
		t.Error("Expected 1st card to be Ace of Spades, got:", deck[0])
	}
}

func TestJokers(t *testing.T) {
	expectedJokerNum := 3
	deck := New(Jokers(expectedJokerNum))
	count := 0
	for _, card := range deck{
		if card.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Expected %v jokers, found %v", expectedJokerNum, count)
	}
}

func TestFilter(t *testing.T) {
	onlyTwoAndThree := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	deck := New(Filter(onlyTwoAndThree))
	for _, card := range deck {
		if card.Rank == Two || card.Rank == Three {
			t.Errorf("Card contained %v", card.Rank)
		}
	}
}

func TestDeck(t *testing.T) {
	expectedNumOfDecks := 3
	deck := New(Deck(expectedNumOfDecks))
	if len(deck) != 13 * 4 * expectedNumOfDecks {
		t.Errorf("Expected %v number of cards, found %v", expectedNumOfDecks * 13 * 4, len(deck))
	}
}