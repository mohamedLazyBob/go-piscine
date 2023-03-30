package piscine

func FindNextPrime(nb int) int {
	nb++
	if IsPrime(nb) {
		return nb
	}
	return FindNextPrime(nb)
}
