package main

import "fmt"

func main() {
	s := struct {
		firstName   string
		friends     map[string]int
		favIceCream []string
	}{
		firstName: "bambank",
		friends: map[string]int{
			"zonk": 1,
			"tonk": 2,
		},
		favIceCream: []string{
			"water",
			"orange",
		},
	}
	fmt.Println(s)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favIceCream {
		fmt.Println(i, val)
	}
}
