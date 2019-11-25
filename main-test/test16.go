package main

import "fmt"

func main() {
	// var numbers [5]int = [5]int{1, 2}
	// fmt.Println(numbers)
	// fmt.Println(numbers[0])
	// fmt.Println(numbers[1])

	// fmt.Println(numbers[3])
	// numbers[3] = 8712412
	// fmt.Println(numbers[3])
	// fmt.Println(numbers)

	// colors := [3]int{2: 4, 0: 2, 1: 3}
	// fmt.Println(colors[2])

	// a := "solo"
	// switch a {
	// case "solo":
	// 	fmt.Println("9")
	// case "polo":
	// 	fmt.Println("8")
	// case "molo":
	// 	fmt.Println("7")
	// }

	// a := 7
	// switch a + 2 {
	// case 9:
	// 	fmt.Println("9")
	// case 8:
	// 	fmt.Println("8")
	// case 7:
	// 	fmt.Println("7")
	// }

	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i * i)
	// }

	// var i = 1
	// for i < 10 {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		fmt.Print(i*j, "\t")
	// 	}
	// 	fmt.Println()
	// }

	// var users = [3]string{"Tom", "Alice", "Kate"}
	// for index, value := range users {
	// 	fmt.Println(index, value)
	// }

	// var users = [3]string{"Tom", "Alice", "Kate"}
	// for i := 0; i < len(users); i++ {
	// 	fmt.Println(users[i])
	// }

	// var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	// var sum = 0

	// for _, value := range numbers {
	// 	if value < 0 {
	// 		continue // переходим к следующей итерации
	// 	}
	// 	sum += value // sum = sum + value
	// }
	// fmt.Println("Sum:", sum) // Sum: 27

	// users := []string{"Tom", "Alice", "Kate"}
	// users = append(users, "Bob")

	// for _, value := range users {
	// 	fmt.Println(value)
	// }

	// var users []string = []string{"Tom", "Alice", "Kate"}
	// fmt.Println(users[2]) // Kate
	// users[2] = "Katherine"

	// for _, value := range users {
	// 	fmt.Println(value)
	// }

	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // с 3-го по 6-й
	users2 := initialUsers[:4]  // с 1-го по 4-й
	users3 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
	fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
	fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]
}
