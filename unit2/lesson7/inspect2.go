package main

import "fmt"

func main() {
	a := 365.2425
	fmt.Printf("Type %T for %[1]v\n", a)
	b := "text"
	fmt.Printf("Type %T for %[1]v\n", b)
	c := 42
	fmt.Printf("Type %T for %[1]v\n", c)
	d := true
	fmt.Printf("Type %T for %[1]v\n", d)
}
