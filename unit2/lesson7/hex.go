package main

import "fmt"

func main() {
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%X %X %X\n", red, green, blue)
	fmt.Printf("%02X%02X%02X\n", red, green, blue)
}
