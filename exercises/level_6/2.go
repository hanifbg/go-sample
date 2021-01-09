package main

import "fmt"

func main() {
	xi := []int{11, 12, 13, 14, 15}
	xii := foo(xi...)

	yi := []int{1, 2, 3, 4, 5}
	yii := bar(yi)

	fmt.Println(xii)
	fmt.Println(yii)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
