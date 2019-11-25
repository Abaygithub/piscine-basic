package piscine

func TrimAtoi(s string) int {
	sig := 1
	res := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			res = int(v-48) + res*10
		} else if v == '-' && res == 0 {
			sig = -1
		}
	}
	return res * sig
}
