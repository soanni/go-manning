package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	messages := []string{
		"hello world",
		"a bad apple",
		"goodbye all",
		"hello world",
		"goodbye all",
		"one amazing feature",
		"two cats",
		"three dogs",
	}

	for _, v := range messages {
		downstream <- v
	}

	close(downstream)
}

func splitGopher(upstream chan string, downstream chan []string) {
	for item := range upstream {
			downstream <- strings.Fields(item)
	}
	close(downstream)
}

func printGopher(upstream chan []string) {
	for v := range upstream {
		for _, i := range v {
			fmt.Println(i)
		}
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan []string)
	go sourceGopher(c0)
	go splitGopher(c0, c1)
	// sourceGopher(c0)
	// filterGopher(c0, c1)
	printGopher(c1)
}
