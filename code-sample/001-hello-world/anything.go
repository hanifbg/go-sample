package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	fmt.Println("something else")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("In foo")
}
