package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous func")
	}()

	fmt.Println("done")
}
