package main

import "fmt"

func counter(number int) func() int {
	var count int = number

	return func() int {
		count++
		return count
	}
}

func main() {

	c := counter(2)

	fmt.Println(c())
}
