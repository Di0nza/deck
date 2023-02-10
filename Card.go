package deck

type Card struct {
	name  string
	suit  string
	value int
}

var names = [13]string{"A", "K", "Q", "V", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
var suits = [4]string{"spade", "diamond", "club", "heart"}
var values = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

func CreateDeck36() []Card {
	var cards []Card
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			card := Card{names[j], suits[i], values[j]}
			cards = append(cards, card)
		}
	}
	return cards
}

func CreateDeck52() []Card {
	var cards []Card
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			card := Card{names[j], suits[i], values[j]}
			cards = append(cards, card)
		}
	}
	return cards
}
