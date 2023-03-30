package piscine

func Lcm(first, second int) int {
	if first == 0 || second == 0 {
		return 0
	}
	return first / gcd(second, first%second) * second
}

func gcd(nb1, nb2 int) int {
	for nb1 != nb2 {
		if nb1 > nb2 {
			nb1 -= nb2
		} else {
			nb2 -= nb1
		}
	}
	return nb1
}
