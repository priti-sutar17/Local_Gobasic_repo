package main

import (
	"fmt"
)

func main() {
	// //   simple switch

	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	//Multiple condition in Switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it is weekend")
	// default:
	// 	fmt.Println("It is workday")

	// }

	//type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a String")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
		}
	}
	whoAmI(true)

}
