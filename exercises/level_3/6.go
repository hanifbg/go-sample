package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("When %v devided by 2, the remain is %v\n", i, i%2)
		}
	}
}
