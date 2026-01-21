package main

import "fmt"

// for is the only contruct for looping in go
func main() {
	// Simple for loop
	// var i int
	// for i = 0; i < 10; i++ {
	// 	fmt.Println("Value of i:", i)
	// }

	// while loop style
	for i := range 10 {
		fmt.Println("Value of ix:", i)
	}
}
