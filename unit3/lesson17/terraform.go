package main

import (
	"fmt"
)

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func main() {
	p := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	Planets(p).terraform()
	fmt.Println(p)
}
