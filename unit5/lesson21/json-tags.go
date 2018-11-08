package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	type location struct {
		Name string `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	//location1 := location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417}
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
	}
	bytes, err := json.MarshalIndent(locations,"","")
	exitOnError(err)
	fmt.Println(string(bytes))
	// fmt.Printf("%+v\n", locations)
}
