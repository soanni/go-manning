package main

import "fmt"

func main() {
	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden)

	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden)

	fmt.Println(administrator == major)

	lightfood := "Robert M. Lightfood Jr."
	administrator = &lightfood
	fmt.Println(administrator == major)

	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(*major)

}
