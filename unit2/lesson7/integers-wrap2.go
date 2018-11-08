package main

import "fmt"

func main() {
	var red uint8 = 255
	fmt.Printf("%08b\n", red)
	red += 2
	fmt.Printf("%08b\n", red)
	fmt.Printf("%v\n", red)
	fmt.Println("==========")
	var number int8 = 127
	fmt.Printf("%08b\n", number)
	number += 3
	fmt.Printf("%08b\n", number)
	fmt.Printf("%v\n", number)
}
