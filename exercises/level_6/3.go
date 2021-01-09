package main

import "fmt"

func main() {
	defer foo() //run after this func end
	fmt.Println("hellloo")
}

func foo() {
	defer func() { //run after this func end
		fmt.Println("defeeer")
	}()

	fmt.Println("the foo")
}
