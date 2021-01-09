package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("anonymous func")
	}

	f()
}
