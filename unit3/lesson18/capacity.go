package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumbers := make([]int, 0, 2)
	fmt.Printf("randomNumbers is of %v - length, %v - capacity\n", len(randomNumbers), cap(randomNumbers))
	for i := 0; i < 100; i++ {
		newNumber := rand.Intn(10) + 1
		randomNumbers = append(randomNumbers, newNumber)
		fmt.Printf("randomNumbers is of %v - length, %v - capacity\n", len(randomNumbers), cap(randomNumbers))
	}
}
