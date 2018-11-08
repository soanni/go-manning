package main

import "fmt"

func main() {
	f := func(message string) {
		fmt.Println(message)
	}
	fmt.Printf("f type is of %T\n", f)
	f("Go-go-go ...")
}
