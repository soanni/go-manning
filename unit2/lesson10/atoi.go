package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown, err := strconv.Atoi("10")
	if err != nil {
		// smth is wrong
	}
	fmt.Println(countdown)
}
