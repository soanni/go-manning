package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

func (l location) String() string {
	// fmt.Println("Name:", l.name, "Lat:", l.lat, "Long:", l.long)
	return fmt.Sprintf("%v (%.1f°, %.1f°)", l.name, l.lat, l.long)
}

type world struct {
	radius float64
}

// distance calculation using the Spherical Law of Cosines
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type gps struct {
	current, destination location
	world
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	// fmt.Println("Distance:", g.distance())
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.destination)
}

type rover struct {
	gps
}

func main() {
	mars := gps{
		current:     location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		destination: location{name: "Elysium Planitia", lat: 4.5, long: 135.9},
		world:       world{radius: 3389.5},
	}
	curiosity := rover{gps: mars}
	// curiosity.message()
	fmt.Println(curiosity.message())
}
