package main

import "fmt"

func main() {
	type location struct {
		name string
		lat  float64
		long float64
	}

	locations := []location{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
	}

	fmt.Printf("%+v\n", locations)
}
