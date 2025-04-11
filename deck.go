package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func (d Deck) print() {
	for i, j := range d {
		fmt.Println(i, j)
	}
}

func newDeck() Deck {
	cards := Deck{}

	types := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	vals := []string{"Ace", "Two", "Three", "Four"}

	for _, val := range vals {
		for _, tp := range types {
			cards = append(cards, val+" of "+tp)
		}
	}

	return cards
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	x := string(bs)

	y := strings.Split(x, ",")

	return Deck(y)

}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for idx := range d {
		newPos := r.Intn(len(d) - 1)

		d[idx], d[newPos] = d[newPos], d[idx]
	}
}
