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
}
