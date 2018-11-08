package main

import (
	"fmt"
	"math/rand"
)

type RandomNumbers []int

func (arr RandomNumbers) appendNewRandomNumber() RandomNumbers {
	arr = append(arr, rand.Intn(10)+1)
	return arr
}

func (arr RandomNumbers) printLenAndCapacity() {
	fmt.Printf("Array is of %v - length, %v - capacity\n", len(arr), cap(arr))
}

func (arr RandomNumbers) printArray() {
	fmt.Println(arr)
}

func main() {
	// randomNumbers := make([]int, 0, 2)
	// numbers := []int{1, 2}
	numbers := make([]int, 0, 2)
	randomNumbers := RandomNumbers(numbers)
	randomNumbers.printLenAndCapacity()
	randomNumbers.printArray()
	for i := 0; i < 100; i++ {
		randomNumbers = randomNumbers.appendNewRandomNumber()
		randomNumbers.printLenAndCapacity()
		randomNumbers.printArray()
	}
}
