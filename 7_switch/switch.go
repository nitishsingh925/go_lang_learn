package main

func main() {

	//  simple Switch

	// i := 3
	// switch i {
	// case 1:
	// 	println("one")
	// case 2:
	// 	println("two")
	// case 3:
	// 	println("three")
	// case 4:
	// 	println("four")
	// case 5:
	// 	println("five")
	// default:
	// 	println("default")
	// }

	// multiple conditions Switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("Weekend")
	// default:
	// 	println("Weekday")
	// }

	// type Switch

	whoAmI := func(i interface{}) {
		switch i := i.(type) {
		case int:
			println("int", i)
		case string:
			println("string", i)
		case bool:
			println("bool", i)
		default:
			println("unknown", i)
		}
	}

	whoAmI(1)
	whoAmI("nitish")
	whoAmI(true)
	whoAmI(3.14)

}
