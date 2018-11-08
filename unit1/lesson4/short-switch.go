package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(20); num {
	case 1:
		fmt.Println("num = 1")
	case 2:
		fmt.Println("num == 2")
	default:
		fmt.Println("num == 3")
	}
}
