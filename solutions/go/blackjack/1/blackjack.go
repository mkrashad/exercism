package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch {
        case card == "ace":
        	return 11
        case card == "jack" || card == "queen" || card == "king" || card == "ten":
        	return 10
		case card == "two": 
        	return 2
        case card == "three": 
        	return 3
        case card == "four": 
        	return 4
        case card == "five": 
        	return 5
        case card == "six": 
        	return 6
        case card == "seven": 
        	return 7
        case card == "eight": 
            return 8
        case card == "nine": 
        	return 9
    	default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerTotal := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case playerTotal == 21 && dealerValue < 10:
		return "W"
	case playerTotal == 21:
		return "S"
	case playerTotal >= 17:
		return "S"
	case playerTotal <= 11:
		return "H"
	case dealerValue >= 7:
		return "H"
	default:
		return "S"
	}
}
