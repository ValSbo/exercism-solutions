package blackjack

var bjRuleset = map[string]int{
	"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7,
	"eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	parsedValue := bjRuleset[card]
	if parsedValue != 0 {
		return parsedValue
	} else {
		return 0
	}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	var parsedDealerCard = ParseCard(dealerCard)

	var sumPlayerCards = ParseCard(card1) + ParseCard(card2)
	var outcomeDecision = ""

	switch {
	case card1 == "ace" && card2 == "ace":
		outcomeDecision = "P"
	case sumPlayerCards == 21:

		if dealerCard != "ace" && !IsFaceCard(dealerCard) && dealerCard != "ten" {
			outcomeDecision = "W"
		} else {
			outcomeDecision = "S"
		}
	case (sumPlayerCards >= 17 && sumPlayerCards <= 20):
		outcomeDecision = "S"
	case (sumPlayerCards >= 12 && sumPlayerCards <= 16):
		if parsedDealerCard < 7 {
			outcomeDecision = "S"
		} else {
			outcomeDecision = "H"
		}
	case sumPlayerCards <= 11:
		{
			outcomeDecision = "H"
		}

	}

	return outcomeDecision

}

func IsFaceCard(card string) bool {

	var isFace = false

	switch card {
	case "jack", "queen", "king":
		isFace = true
	default:
		isFace = false
	}

	return isFace
}
