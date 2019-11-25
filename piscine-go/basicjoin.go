package piscine

func BasicJoin(strs []string) string {
	var join string
	for _, i := range strs {
		join += i
	}
	return join
}
