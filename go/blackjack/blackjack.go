package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	var parsedValue = 0
	switch card {
	case "ace":
		parsedValue = 11
	case "two":
		parsedValue = 2
	case "three":
		parsedValue = 3
	case "four":
		parsedValue = 4
	case "five":
		parsedValue = 5
	case "six":
		parsedValue = 6
	case "seven":
		parsedValue = 7
	case "eight":
		parsedValue = 8
	case "nine":
		parsedValue = 9
	case "ten", "jack", "queen", "king":
		parsedValue = 10

	default:
		parsedValue = 0
	}
	return parsedValue
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
	default:
		outcomeDecision = "H"

	}

	return outcomeDecision

}

func IsFaceCard(card string) bool {

	var isFace = false

	switch card {
	case "jack":
		isFace = true
	case "queen":
		isFace = true
	case "king":
		isFace = true
	default:
		isFace = false
	}

	return isFace
}
