package main

import "fmt"

func main() {
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const distance = 24000000000000000000
	// km := distance
	days := distance / lightSpeed / secondsPerDay
	// fmt.Println("Andromeda Galaxy is", distance, "km away.")
	fmt.Println("Andromeda Galaxy is", days, "light days away.")
	fmt.Printf("That is %v years of travel at light speed.\n", days/365)
}
