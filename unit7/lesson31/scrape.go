package main

import (
	"fmt"
	"sync"
)

// Visited tracks whetner web pages have been visited.
// Its methods may be used concurrently from multiple goroutines.
type Visited struct {
	// mu guards the visited map
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	visited := Visited{visited: map[string]int{"http://crossover.com": 1}}
	go visited.VisitLink("http://crossover.com")
	go visited.VisitLink("http://crossover.com")
	fmt.Printf("%+v\n", visited)
}
