package piscine

func UltimateDivMod(a *int, b *int) {
	*a, *b = *a / *b, *a%*b
}
