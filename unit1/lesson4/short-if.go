package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if num := rand.Intn(3); num == 1 {
		fmt.Println("num = 1")
	} else if num == 2 {
		fmt.Println("num == 2")
	} else {
		fmt.Println("num == 3")
	}
}
