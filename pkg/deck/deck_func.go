package deckOfCards

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Cards []Card

const (
	Joker CardValue = iota
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
	Spade   SuitValue = 13
	Diamond SuitValue = 26
	Club    SuitValue = 39
	Heart   SuitValue = 52
	NoSuit  SuitValue = 0
)

var names = [14]string{"Joker", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func (c Card) String() string {
	var suit string
	if uint8(c.Name) == uint8(Joker) {
		return "Joker; "
	}
	switch c.Suit {
	case 13:
		suit = "spade"
	case 26:
		suit = "diamond"
	case 39:
		suit = "club"
	case 52:
		suit = "heart"
	}
	return fmt.Sprintf("%s of %s's; ", names[c.Name], suit)
}

func NewDeck() []Card {
	cards := make([]Card, 0, 52)
	suitI := 13
	for i := 0; i < 4; i++ {
		for j := 1; j < 14; j++ {
			card := Card{Name: CardValue(j), Suit: SuitValue(suitI)}
			cards = append(cards, card)
		}
		suitI += 13
	}
	return cards
}

func PrintDeck(d []Card) {
	for i := range d {
		fmt.Print(d[i])
	}
	fmt.Print("\n\n")
}

func ShuffleDeck(d []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d),
		func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func SortDeck(d []Card) {
	sort.Slice(d, func(i, j int) bool {
		return uint8(d[i].Suit)+uint8(d[i].Name) < uint8(d[j].Suit)+uint8(d[j].Name)
	})
}

func appendDeck(current []Card, appended []Card) []Card {
	for _, card := range appended {
		current = append(current, card)
	}
	return current
}

func CreateMultiplyDeck(decksAmount int) []Card {
	d := make([]Card, 0, decksAmount*52)
	for i := 0; i < decksAmount; i++ {
		d = appendDeck(d, NewDeck())
	}
	return d
}

func DeleteCard(d []Card, card CardValue) []Card {
	b := d[:0]
	for i, x := range d {
		if d[i].Name != card {
			b = append(b, x)
		}
	}
	return b
}

func AddJokers(d []Card, amount int) []Card {
	joker := Card{Name: CardValue(Joker), Suit: SuitValue(NoSuit)}
	for i := 0; i < amount; i++ {
		d = append(d, joker)
	}
	return d
}
