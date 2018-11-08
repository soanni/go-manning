package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus", "Sedna")
	fmt.Println(dwarfs)
}
