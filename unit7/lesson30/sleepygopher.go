package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher()
	fmt.Printf("Count started ...\n")
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}
