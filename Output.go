package main

import "fmt"

func main() {

	//1.Print() output

	/*var i, j string = "priti", "aniket"
	fmt.Print(i)
	fmt.Print(j)*/

	//printing multiple variables
	/*var i, j string = "hello", "world"
	fmt.Print(i, "\n", j)*/

	//space betweent 2 arguments
	/*var i, j string = "hello", "world"
	var p, q = 10, 20
	fmt.Print(i, " ", j)
	fmt.Print(p, " ", q)*/

	/*2.Println() output
	var i, j string = "hello", "world"
	fmt.Println(i, j)*/

	//Printf() output
	var i string = "hello"
	var j int = 15
	fmt.Printf("i has value:%v and type:%T", i, i)
	fmt.Printf("\nj has value:%v and type:%T", j, j)
}
