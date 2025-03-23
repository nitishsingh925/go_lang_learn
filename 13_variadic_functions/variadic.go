package main

import "fmt"

// func sum(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

func anything(anythingValue ...interface{}) {

	fmt.Println(anythingValue...)
}
func main() {

	// fmt.Println(sum(1, 2, 3, 4, 5))

	anything("nitish", 12, 24.44, true)

}
