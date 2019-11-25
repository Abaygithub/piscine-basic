package main

import "fmt"

// func StrLen(str string) int {
// 	r := []rune(str)
// 	ln := 0
// 	for l := range r {
// 		ln = l
// 	}
// 	ln = ln + 1
// 	return ln
// }

func main() {
	str := "Hello World!"	
	s := []string{str}
	r := []rune(str)
	
	// l := []len{str}
	fmt.Println(s)
	fmt.Println(r)
	fmt.Println(len(s))
	fmt.Println(len(r))

}