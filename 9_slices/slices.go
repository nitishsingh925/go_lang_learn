package main

import "fmt"

// sluce -> dynamic
// most used construct in go
//   - useful methods
func main() {

	// uninitialized slice is nil
	// var nums []int
	// fmt.Println("this is a nil slice", nums == nil)
	// fmt.Println("length:", len(nums))
	// fmt.Println("capacity:", cap(nums))
	// fmt.Println("value:", nums)

	// initialized slice
	// var nums = make([]int, 2, 5)
	// nums[0] = 1
	// nums[1] = 2
	// nums := []int{}

	// now this not work because length is 2
	// nums[2] = 3
	// nums[3] = 4

	// nums = append(nums, 3)
	// nums = append(nums, 4, 5, 6)
	// nums = append(nums, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	// fmt.Println("this is a nil slice", nums == nil)
	// fmt.Println("length:", len(nums))
	// fmt.Println("capacity:", cap(nums))
	// fmt.Println("value:", nums)

	// copy function

	// var copyNum1 = make([]int, 0, 5)

	// copyNum1 = append(copyNum1, 2)
	// copyNum1 = append(copyNum1, 3)

	// copyNum2 := make([]int, len(copyNum1))

	// copy(copyNum2, copyNum1)

	// fmt.Println(copyNum1, copyNum2)

	// slice operator

	// nums := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(nums[1:3])
	// fmt.Println(nums[:3])
	// fmt.Println(nums[1:])

	// slices

	// num1 := []int{1, 2, 3}
	// num2 := []int{1, 2, 4}

	// fmt.Println(slices.Equal(num1, num2))

	// 2 dimentional slice
	nums := [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(nums)

}
