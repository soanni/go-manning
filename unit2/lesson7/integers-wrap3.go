package main

import "fmt"

func main() {
	var red uint8 =0
	fmt.Printf("%08b\n", red)
	red--
	fmt.Printf("%08b\n", red)
	fmt.Printf("%v\n", red)
	fmt.Println("==========")
	var number int8 = -128
	fmt.Printf("%08b\n", number)
	number--
	fmt.Printf("%08b\n", number)
	fmt.Printf("%v\n", number)
}
