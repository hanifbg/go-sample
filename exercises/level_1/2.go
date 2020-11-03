package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	x = 2
	y = "Hello"
	z = true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
