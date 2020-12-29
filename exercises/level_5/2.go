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

	m := map[string]person{
		p1.lastName: p1,

		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, v := range v.favoriteFlavor {
			fmt.Println(i, v)
		}
	}
}
