package main

import "fmt"

type person struct {
	firstName string
}

func changeMe(p *person) {
	p.firstName = "apa"
}

func main() {
	p1 := person{
		firstName: "hanifbg",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
