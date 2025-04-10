package main

import "fmt"

// pointers are used to pass by reference
func changeNumReference(num int) {
	num = 5
	fmt.Println("num in changeNum reference function", num)
}

// pointers are used to pass by value

func changeNumValue(num *int) {
	*num = 5
	fmt.Println("num in changeNum value function", *num)
}

func main() {

	num := 1

	changeNumReference(num)

	changeNumValue(&num)

	fmt.Println("num in main function", num)
}
