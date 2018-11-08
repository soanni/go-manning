package main

import "fmt"

func main() {
	c := 'a'
	c += 3
	fmt.Printf("%c %[1]v %[1]T", c)
}
