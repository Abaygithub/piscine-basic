package piscine

func Swap(a *int, b *int) {
	zero := *a
	one := *b
	*a = one
	*b = zero
}
