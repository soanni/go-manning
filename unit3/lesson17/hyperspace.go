package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Venus ", "Earth ", " Mars "}
        
	dwarfArray := [...]string{" Ceres  ", " Pluto    ", " Haumea  ", " Makemake", " Eris    "}
	dwarfSlice := dwarfArray[:]
	
	hyperspace(planets)
	hyperspace(dwarfSlice)

	fmt.Println(strings.Join(planets, ""))
	fmt.Println(planets)
	fmt.Println(dwarfSlice)
	fmt.Println(dwarfArray)
}
