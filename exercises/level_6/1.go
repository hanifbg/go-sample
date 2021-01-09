package main

import "fmt"

func main() {
	n := foo()
	i, s := bar()

	fmt.Println(n, i, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "hello"
}
