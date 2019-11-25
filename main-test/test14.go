package main

import "fmt"

func main() {
	len := 1
	str := "He he he"
	for _, j := range str {
		if j == ' ' || j == '\n' {
			len++
		}
	}
	fmt.Print("Dlina stroki ravna ")
	fmt.Println(len)
}
