package main

import (
	"fmt"
	"math"
)

type square struct {
	sisi float64
}

type circle struct {
	jari float64
}

func (c circle) area() float64 {
	return math.Pi * c.jari * c.jari
}

func (s square) area() float64 {
	return s.sisi * s.sisi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		jari: 12.345,
	}

	s := square{
		sisi: 15,
	}

	info(c)
	info(s)
}
