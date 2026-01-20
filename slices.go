package main

// slice-> dynamic
// most used construct in go
// useful methods
func main() {
	// unintialize slice is nil
	// var num []int
	// fmt.Println(num == nil)
	// fmt.Println(len(num))

	//1.make is function
	// var nums = make([]int, 2  , 5) //making slice using make function
	// fmt.Println(nums == nil)

	//2.
	// nums := []int{}

	//cap means capacity ->maximum numbers of elements can fit and it is a function
	// fmt.Println(cap(nums))

	//append is a method
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums[0] = 3
	// nums[4] = 2
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	//using copy function we can copy the slice
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)

	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	//slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[1:])

	//slices package
	// var num1 = []int{1, 2, 3}
	// var num2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(num1, num2))

	// //2D slices
	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

}
