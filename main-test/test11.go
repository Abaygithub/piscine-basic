package main

import "fmt"

func ToUpper(value string) string {
	var result string
	for _, letter := range value {
		if letter >= 'a' && letter <= 'z' {
			letter = letter - 32
			result = result + string(letter)
		} else {
			result = result + string(letter)
		}
	}
	return result

}

func ToLower(value string) string {
	var result string
	for _, letter := range value {
		if letter >= 'A' && letter <= 'Z' {
			letter = letter + 32
			result = result + string(letter)
		} else {
			result = result + string(letter)
		}
	}
	return result
}

func Capitalize(s string) string {
	newWord := true
	for _,value := range s{
	if newWord{
		func ToUpper(value string)string{}
	} else !newWord{
		func ToLower(value string)string{}
	}
		if !('a' <= value && value <= 'z') && !('A' <= value && value <= 'Z') && !('1' <= value && value <= '9') {
			newWord = true
		} else {
			newWord = false
		}
		return s


}



	// newWord := true
	// myRune := []rune(s)
	// for index, value := range myRune {
	// 	if ('a' <= value && value <= 'z') && newWord {
	// 		value = value - 32
	// 		myRune[index] = value
	// 		newWord = false
	// 	} else if ('A' <= value && value <= 'Z') && !newWord {
	// 		value = value + 32
	// 		myRune[index] = value
	// 	}
	// 	if !('a' <= value && value <= 'z') && !('A' <= value && value <= 'Z') && !('1' <= value && value <= '9') {
	// 		newWord = true
	// 	} else {
	// 		newWord = false
	// 	}
	// }
	// return string(myRune)
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
