package main

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

func main() {
	var t talker

	t = martian{}
	fmt.Println(t.talk())
	shout(t)

	t = laser(3)
	fmt.Println(t.talk())
	shout(t)

	shout(starship{laser(3)})
}
