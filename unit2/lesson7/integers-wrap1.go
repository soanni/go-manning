package main

import "fmt"

func main() {
	var red uint8 = 255
	fmt.Printf("%08b\n", red)
	red++
	fmt.Printf("%08b\n", red)
	fmt.Println("==========")
	var number int8 = 127
	fmt.Printf("%08b\n", number)
	number++
	fmt.Printf("%08b\n", number)
}
