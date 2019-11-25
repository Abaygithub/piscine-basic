package main

import "fmt"

func main(){
	//a := []string{"Hello","world"}
	b := []int32{3,4,5,6,7}
	//arr := make([]string,5)
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
}