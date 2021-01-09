package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("hello i am", p.first, p.last, "and i am", p.age, "old")
}

func main() {
	p1 := person{
		first: "first",
		last:  "last",
		age:   29,
	}

	p1.speak()
}
