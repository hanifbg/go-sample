package main

import "fmt"

func main() {
	switch {
	case (9 == 0):
		fmt.Println("this is false")
	case (8 == 8):
		fmt.Println("this is true")
	default:
		fmt.Println("this is default")
	}
}
