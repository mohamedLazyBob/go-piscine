package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	}
	return FindNextPrime(nb)
}
