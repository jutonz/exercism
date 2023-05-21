package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0

	switch card {
	case "ace":
		value = 11
	case "king", "queen", "jack", "ten":
		value = 10
	case "nine":
		value = 9
	case "eight":
		value = 8
	case "seven":
		value = 7
	case "six":
		value = 6
	case "five":
		value = 5
	case "four":
		value = 4
	case "three":
		value = 3
	case "two":
		value = 2
	}

	return value
}

const SPLIT = "P"
const WIN = "W"
const STAND = "S"
const HIT = "H"

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return SPLIT
	}

	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	sum := card1Value + card2Value
	dealerCardValue := ParseCard(dealerCard)
	var action string

	switch {
	case sum == 21 && dealerCardValue < 10:
		action = WIN
	case sum == 21 && dealerCardValue >= 10:
		action = STAND
	case sum <= 20 && sum >= 17:
		action = STAND
	case sum <= 16 && sum >= 12 && dealerCardValue >= 7:
		action = HIT
	case sum <= 16 && sum >= 12 && dealerCardValue < 7:
		action = STAND
	default:
		action = HIT
	}

	return action
}
