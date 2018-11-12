package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
p.age++
}

func main() {
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	fmt.Printf("%T\n", timmy)
	fmt.Printf("%+v\n", timmy)
	birthday(timmy)
	fmt.Printf("%+v\n", timmy)
	// birthday(*timmy)
	// fmt.Printf("%+v\n", timmy)
}
