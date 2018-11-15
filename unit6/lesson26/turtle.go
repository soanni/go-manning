package main

import "fmt"

type location struct {
	x, y int
}

func (l *location) up() {
	l.y--
}

func (l *location) down() {
	l.y++
}

func (l *location) right() {
	l.x++
}

func (l *location) left() {
	l.x--
}

type turtle struct {
	location
}

// func (t *turtle) up() {
// 	t.upY()
// }

// func (t *turtle) down() {
// 	t.downY()
// }

// func (t *turtle) left() {
// 	t.leftX()
// }

// func (t *turtle) right() {
// 	t.rightX()
// }

func (t turtle) String() string {
	return fmt.Sprintf("x: %v, y: %v", t.x, t.y)
}

func main() {
	t := turtle{location: location{0, 0}}
	t.up()
	t.right()
	t.up()
	t.right()
	t.left()
	fmt.Println(t)
}
