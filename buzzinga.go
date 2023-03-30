package piscine

func BuzZinga(number int) {
	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 { // If the number is divisible by 3 and 5, print `BuzZinga` followed by a newline
			printStr("BuzZinga")
		} else if i%3 == 0 { // If the number is divisible by 3, print `Buz` followed by a newline
			printStr("Buz")
		} else if i%5 == 0 { // If the number is divisible by 5, print `Zinga` followed by a newline
			printStr("Zinga")
		} else if i < 0 || (i%3 != 0 && i%5 != 0) { // If the number is negative or not divisible by 3 or 5, print `*` followed by a newline.
			printStr("*")
		}
	}
	if number == 0 { // if the number is 0, print `BuzZinga` followed by a newline.
		printStr("BuzZinga")
	}
}
