package main

import "fmt"

func main() {
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("%T\n", dwarfArray)
	fmt.Printf("%T\n", dwarfSlice)
	fmt.Printf("%T\n", dwarfs)
}
