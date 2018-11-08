package main

import "fmt"

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// dwarfs = append(dwarfs, "Orcus", "Sedna")
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])
	dwarfs1 := append(dwarfs, "Orcus", "Sedna")
	dump("dwarfs1", dwarfs)
	dump("dwarfs1[1:2]", dwarfs[1:2])
	dwarfs1[0] = "Earth"
	fmt.Println(dwarfs)
	fmt.Println(dwarfs1)
}
