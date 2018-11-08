package main

import "fmt"

func main() {
	fmt.Println("one \ntwo")
	fmt.Println(`one \n two`)
	fmt.Println(`
one
two
three`)
fmt.Printf("%v is a %[1]T\n", "literal string")
fmt.Printf("%v is a %[1]T\n", `raw literal string`)
fmt.Println(`C:\go`)
}
