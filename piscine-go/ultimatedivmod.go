package piscine

func UltimateDivMod(a *int, b *int) {
	var first int
	var second int
	first = *a / *b
	second = *a % *b
	*a = first
	*b = second
}
