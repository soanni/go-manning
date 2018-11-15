package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew", int(*l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	pew := laser(2)
	shout(&pew)
	// shout(pew)
}
