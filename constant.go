package main

import "fmt"

// there ar 2 type of constant Typed and Untyped
// const PI = 3.14
// const A int = 1 \\typed
// const B = 1. \\untyped
// multiple type constant
const (
	A int = 1
	B     = 3.14
	C     = "Hii"
)

func main() {
	//fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
