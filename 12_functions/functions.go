package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }
// func add(a, b int) int {
// 	return a + b
// }

// func getData() (string, int, bool, float64) {
// 	return "nitish", 25, true, 3.14
// }

func processData() func(a int) int {
	return func(a int) int {
		return a * 1
	}
}
func main() {
	// fmt.Println(add(1, 2))

	// fmt.Println(getData())

	// fn := processData()
	// fmt.Println(fn(5))

	fmt.Println(processData()(5))

}
