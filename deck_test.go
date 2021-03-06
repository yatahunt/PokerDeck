package PokerDeck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const seedTest int64 = 1234567890

func TestNewDeck(t *testing.T) {
	deck1 := NewDeck(seedTest)
	deck2 := NewDeck(seedTest)
	assert.Len(t, deck1.cards, 52)
	assert.Len(t, deck2.cards, 52)

	same := true
	for i := range deck1.cards {
		same = same && (deck1.cards[i] == deck2.cards[i])
	}
	assert.False(t, same)
}

func TestDraw(t *testing.T) {
	deck := NewDeck(seedTest)

	cards := deck.Draw(5)
	assert.Len(t, cards, 5)
	assert.False(t, deck.Empty())

	deck.Draw(52 - 5)
	assert.True(t, deck.Empty())
}

func TestEmpty(t *testing.T) {
	deck := NewDeck(seedTest)
	assert.False(t, deck.Empty())

	deck.Draw(51)
	assert.False(t, deck.Empty())

	deck.Draw(1)
	assert.True(t, deck.Empty())
}
