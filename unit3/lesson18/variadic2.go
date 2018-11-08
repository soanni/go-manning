package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	planets := []string{"Earth", "Mars", "Jupiter"}
	fmt.Println(terraform("Old", planets...))
}
