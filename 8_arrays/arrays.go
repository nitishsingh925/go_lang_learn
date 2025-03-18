package main

import "fmt"

// number sequence of specific length
func main() {

	// var intergers [5]int
	// intergers[0] = 1
	// fmt.Println(intergers)

	// var boolsValues [5]bool
	// boolsValues[0] = true
	// fmt.Println(boolsValues)

	// var stringsValues [5]string
	// stringsValues[0] = "nitish"
	// fmt.Println(stringsValues)

	// var floatValues [5]float32
	// floatValues[0] = 3.14
	// fmt.Println(floatValues)

	// var complexValues [5]complex64
	// complexValues[0] = 1 + 2i
	// fmt.Println(complexValues)

	// 1 dimentional array
	// var intergers1 = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(intergers1)

	// 2 dimentional array
	// var intergers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(intergers2)

	// dimentional array
	var intergers3 = [2][3][4]int{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, {{13, 14, 15, 16}, {17, 18, 19, 20}, {21, 22, 23, 24}}}
	fmt.Println(intergers3)

}
