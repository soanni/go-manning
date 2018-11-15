package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c:
			fmt.Println("gopher", gopherId, "has finished sleeping")
		case <-timeout:
			fmt.Println("run out of time")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore ...")
	c <- id
}
