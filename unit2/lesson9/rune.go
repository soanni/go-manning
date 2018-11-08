package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	var smile rune = 128515

	fmt.Printf("%v is of type %[1]T\n", pi)
	fmt.Printf("%v is of type %[1]T\n", bang)
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)
	fmt.Printf("%c\n", smile)
}
