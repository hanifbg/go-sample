package main

import "fmt"

func main() {

	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	sm := [][]string{s1, s2}
	fmt.Println(sm)

	for i, v := range sm {
		for j, val := range v {
			fmt.Printf("slice : %d, index : %d, value : %s\n", i, j, val)
		}
	}
}
