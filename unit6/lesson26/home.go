package main

import "fmt"

func main() {
	samara := "Samara"

	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &samara
	fmt.Println(*home)
}
