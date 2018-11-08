package main

import "fmt"

func main() {
	age := float64(32)
	marsDays := 687.0
	earthDays := 365.2425
	fmt.Println("I am", age*earthDays/marsDays, "years old on Mars.")
	fmt.Println("I am", int(age*earthDays/marsDays), "years old on Mars.")
}
