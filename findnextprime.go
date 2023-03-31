package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	}
	nb++
	return FindNextPrime(nb)
}
