package main

import "fmt"

func main() {
	// data := 10
	// fmt.Print("Enter a Number :")
	// var input int
	// fmt.Scanln(&input)

	// if input%2 == 0 {
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("hey you are odd")
	// }

	// if x := 10; x%2 == 0 {
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("hey you are odd")
	// }
	// fmt.Println(x)

	data := 10
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 100
		fmt.Println(data)
	case 11:
		data = 100
		fmt.Println(data)
		fallthrough // execute the next case also
	case 15:
		data = 100
		fmt.Println(data)
		fallthrough
	default:
		data = 10909

	}
}
