package main

import (
	"deck"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var names = [13]string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
var suits = [4]string{"spade", "diamond", "club", "heart"}
var values = [13]int{1, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
var suitValues = [4]int{1, 2, 3, 4}

func CreateDeck36() []deck.Card {
	cards := make([]deck.Card, 0, 36)
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			card := deck.Card{Name: names[j], Suit: suits[i], Value: values[j], SuitValue: suitValues[i]}
			cards = append(cards, card)
		}
	}
	return cards
}

func CreateDeck52() []deck.Card {
	cards := make([]deck.Card, 0, 52)
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			card := deck.Card{Name: names[j], Suit: suits[i], Value: values[j], SuitValue: suitValues[i]}
			cards = append(cards, card)
		}
	}
	return cards
}

func PrintDeck(d []deck.Card) {
	border := len(d) / 4
	for i := 0; i < len(d); i++ {
		fmt.Print(d[i].Name + " " + d[i].Suit + "; ")
		if float32((i+1)%border) == 0 {
			fmt.Println()
		}
	}
	fmt.Print("\n\n")
}

func ShuffleDeck(d []deck.Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d),
		func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func SortDeck(d []deck.Card) {
	sort.Slice(d, func(i, j int) bool {
		return d[i].SuitValue < d[j].SuitValue
	})
	counter := 0
	border := len(d) / 4
	for i := 0; i < 4; i++ {
		for j := counter; j < border; j++ {
			for k := j; k < border; k++ {
				if d[j].Value > d[k].Value {
					d[j], d[k] = d[k], d[j]
				}
			}
		}
		counter = border
		border += len(d) / 4
	}
}

func main() {
	newDeck := CreateDeck36()
	PrintDeck(newDeck)
	ShuffleDeck(newDeck)
	PrintDeck(newDeck)
	SortDeck(newDeck)
	PrintDeck(newDeck)
}
