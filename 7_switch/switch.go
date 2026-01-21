package main

import (
	"fmt"
)

func main() {

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("It's the weekend")
	// case time.Monday:
	// 	println("It's Monday")
	// default:
	// 	println("It's a weekday")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			if t < 5 {
				fmt.Println("it is less than 5")
			} else if t > 5 && t < 10 {
				fmt.Println("it is between 5 and 10")
			} else {
				fmt.Println("it is greater than 10")
			}
		default:
			fmt.Println("it is greater than 10")
		}
	}

	whoAmI(7)
}
