package main

import "fmt"

// const name = "nitish"

// const (
// 	firstName = "nitish"
// 	lastName  = "singh"
// )

func main() {

	const (
		firstName = "nitish"
		lastName  = "singh"

		name = firstName + " " + lastName
	)

	fmt.Println(name)
}
