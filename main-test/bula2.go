package main

import "fmt"

func main() {
	b := []int32{3,4,5,6,7}
	for key, value := range b {
		fmt.Println("index is",key,",  value is",value)
		//fmt.Println(value)
	}
}