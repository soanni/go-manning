package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	bradbury := location{-1.9462, 354.4734}
	curiosity := bradbury
	curiosity.long += 0.0106

	fmt.Printf("%+v\n", bradbury)
	fmt.Printf("%+v\n", curiosity)
}
