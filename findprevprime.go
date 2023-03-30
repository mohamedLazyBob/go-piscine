package piscine

func FindPrevPrime(nb int) int {
	if nb <= 2 {
		return 0
	}
	nb--
	if IsPrime(nb) {
		return nb
	}
	return FindPrevPrime(nb)
}
