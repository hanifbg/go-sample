package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%v %b %#X\n", x, x, x)
	y := 42 << 1
	fmt.Printf("%v %b %#X", y, y, y)
}
