package main

import (
	"fmt"
	"github.com/ssimunic/gosensors"
)

func main() {
	sensors, err := gosensors.NewFromSystem()

	if err != nil {
		panic(err)
	}

	fmt.Println(sensors)

	for chip := range sensors.Chips {
		// Iterate over entries
		for key, value := range sensors.Chips[chip] {
			// If CPU or GPU, print out
			if key == "CPU" || key == "GPU" {
				fmt.Println(key, value)
			}
		}
	}
}
