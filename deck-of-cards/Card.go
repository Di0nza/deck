package deck_of_cards

type CardValue uint8
type SuitValue uint8

type Card struct {
	Name CardValue
	Suit SuitValue
}
