package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	sensor := fakeSensor
	fmt.Printf("sensor is of type %T\n", sensor)
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
