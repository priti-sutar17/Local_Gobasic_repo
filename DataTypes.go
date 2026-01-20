package main

import "fmt"

func main() {
	//Boolean data type
	/*var b1 bool = true //type declared
	var b2 = true      //undclare type
	var b3 bool        //type without value
	b4 := true         //untype with value

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)*/

	//Numeric data type
	//1.Integer data type
	//i.Signed integer
	/*var x int = 500
	var y int = 7800
	fmt.Printf("Type:%T, Value:%v", x, x)
	fmt.Printf("\nType:%T, Value:%v", y, y)*/

	//ii.unsigend t=integer
	/*var x uint = 800
	var y uint = 9000
	fmt.Printf("Type:%T, Value:%v", x, x)
	fmt.Printf("\nType:%T, Value:%v", y, y)*/

	//2.Float data type
	var x float32 = 123.78
	var y float32 = 3.4e+38
	var z float64 = 1.7e+308
	fmt.Printf("Type:%T, Value:%v", x, x)
	fmt.Printf("\nType:%T, Value:%v", y, y)
	fmt.Printf("\nType:%T, Value:%v", z, z)
}
