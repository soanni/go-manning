package main

import "fmt"

func main() {
	message := "shalom"
	c := message[5]
	fmt.Printf("%c %[1]v %[1]T\n", c)
	// message[5] = 'd'
	// c = 'd'
}
