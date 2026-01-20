package main

import "fmt"

//var a int
//var b int = 2
//var c = 3

func main() {
	//1.declararation with inital values
	/*var stud1 string = "priti"
	var stud2 = "Aniket"
	X := 2

	fmt.Println(stud1)
	fmt.Println(stud2)
	fmt.Println(X)*/

	/* 2.declation without intial values
	var a string
	var b int
	var c bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)*/

	// 3. Value assignment after declaration
	/*var stud1 string
	stud1 = "Priti"
	fmt.Println(stud1)*/

	//4 outside of function with var keyword
	/*a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)*/

	// 5. Multipe variable on same line
	/*var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)*/

	//6.if Type is not specify then diff type variable declarenon the same line

	/*var a, b = 6, "hello"
	c, d := 7, "world!"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)*/

	// 7.variable declation in a block

	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
