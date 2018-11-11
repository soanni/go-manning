package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Printf("%T\n", address)
	fmt.Println(*address)

	fmt.Println(*&answer)
}
