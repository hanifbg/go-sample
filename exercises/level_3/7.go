package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("When %v devided by 2, the remain is %v\n", i, i%2)
		} else if i%3 == 0 {
			fmt.Printf("When %v devided by 3, the remain is %v\n", i, i%3)
		} else if i%4 == 0 {
			fmt.Printf("When %v devided by 4, the remain is %v\n", i, i%4)
		}
	}
}
