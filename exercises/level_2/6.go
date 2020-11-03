package main

import "fmt"

const (
	w = 2020
	x = 2020 + iota
	y = 2020 + iota
	z = 2020 + iota
)

func main() {
	fmt.Println(w, x, y, z)
}
