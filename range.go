package main

import "fmt"

// range is use for iterating over data structurs
func main() {

	//slice iterating using range
	//nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// sum = sum + num

	//}
	// fmt.Println(sum)

	// map itreating using range
	// m := map[string]string{"fname": "Priti", "lname": "Aniket"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// for k := range m {
	// 	fmt.Println(k)
	// }

	//using String
	//unicode code point rune
	//i is stating byte of rune
	//
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
