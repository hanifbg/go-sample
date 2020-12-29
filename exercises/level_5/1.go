package main

import "fmt"

type person struct {
	firstName      string
	lastName       string
	favoriteFlavor []string
}

func main() {
	p1 := person{
		firstName:      "first",
		lastName:       "last",
		favoriteFlavor: []string{"vanila", "chocolate"},
	}

	p2 := person{
		firstName:      "depan",
		lastName:       "belakang",
		favoriteFlavor: []string{"coklat", "strawberry"},
	}
	fmt.Println(p1.firstName, p1.lastName, p1.favoriteFlavor)
	fmt.Println(p2.firstName, p2.lastName, p2.favoriteFlavor)

	for _, v := range p1.favoriteFlavor {
		fmt.Println(v)
	}

	for _, v := range p2.favoriteFlavor {
		fmt.Println(v)
	}
}
