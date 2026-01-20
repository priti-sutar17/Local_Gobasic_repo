package main

import (
	"fmt"
	"math"
)

/*If statemant
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(
		//pow(2, 2, 10),
		//pow(3, 2, 20),
	)

	//end user
	var x, n, lim float64

	fmt.Println("Enter base x:")
	fmt.Scan(&x)

	fmt.Println("Enter Power n:")
	fmt.Scan(&n)

	fmt.Println("Enter limit lim:")
	fmt.Scan(&lim)

	result := pow(x, n, lim)
	fmt.Println("Result:\n", result)

}*/

// If and else statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g>=%g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
