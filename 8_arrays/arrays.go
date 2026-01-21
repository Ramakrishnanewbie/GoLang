package main

import "fmt"

func main() {

	//Arrays are mostly used when the length is known and fixed
	// Slice are more commonly used in go as they are more flexible
	//Syntax: var name [length]type

	//default values of 0 are added for int arrays
	// var nums [4]int

	// //Array Length
	// fmt.Println(len(nums))

	// // Add elements at an index
	// nums[0] = 99
	// fmt.Println(nums)

	// var bools [10]bool
	// fmt.Println(len(bools))
	// fmt.Println(bools)

	// var strings [1]string
	// fmt.Println(len(strings))
	// fmt.Println(strings)

	// nums := [3]int{1, 2, 3}
	// fmt.Println(len(nums))
	// fmt.Println(nums)

	// 2D arrays
	nums := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(len(nums))
	fmt.Println(nums)
}
