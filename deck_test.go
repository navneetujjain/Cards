package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First Card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Last Card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveLoadFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	x := newDeckFromFile("_decktesting")

	if len(x) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(x))
	}

	os.Remove("_decktesting")
}
