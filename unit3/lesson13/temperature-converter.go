package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var c celsius = 32
	var k kelvin = 300
	var f fahrenheit = 300

	fmt.Println("32C ->", c.kelvin(), "K")
	fmt.Println("300K->", k.celsius(), "C")
	fmt.Println("300F->", f.celsius(), "C")
}
