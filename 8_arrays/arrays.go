package main

import "fmt"

func main() {
	//default values of 0 are added for int arrays
	var nums [4]int

	//Array Length
	fmt.Println(len(nums))

	// Add elements at an index
	nums[0] = 99
	fmt.Println(nums)

	var bools [10]bool
	fmt.Println(len(bools))
	fmt.Println(bools)

	var strings [1]string
	fmt.Println(len(strings))
	fmt.Println(strings)
}
