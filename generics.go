package main

import "fmt"

/* reduse repetion of coding logic so we use generics*/

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func main() {
	// vals := []bool{true, false, true}
	names := []string{"priti", "aniket"}
	printSlice(names)
	// printStringSlice(names)

}
