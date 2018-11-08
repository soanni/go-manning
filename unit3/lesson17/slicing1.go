package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	giants := planets[4:8]
	gasGiants := giants[0:2]
	iceGiants := giants[2:4]

	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets, iceGiantsMarkII, iceGiants, gasGiants)
}
