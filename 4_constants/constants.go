package main

import "fmt"

const age int = 30

func main() {
	const name string = "bot"
	fmt.Println(name)
	fmt.Println(age)
	const (
		port int    = 8008
		host string = "localhost"
	)
	fmt.Println(port)
	fmt.Println(host)
}
