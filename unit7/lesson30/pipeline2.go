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
	}

	for _, v := range messages {
		downstream <- v
	}

	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(strings.ToUpper(v))
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	// sourceGopher(c0)
	// filterGopher(c0, c1)
	printGopher(c1)
}
