package main

import "fmt"

func main() {
	switch favSport := "Soccer"; favSport {
	case "Basket Ball":
		fmt.Println("You like Basket")
	case "Soccer":
		fmt.Println("You like Soccer")
	default:
		fmt.Println("this is default")
	}
}
