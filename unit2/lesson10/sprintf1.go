package main

import "fmt"

func main() {
	countdown := 9
	countdown = 8
	str := fmt.Sprintf("Launch in T minus %v seconds", countdown)
	fmt.Println(str)
}
