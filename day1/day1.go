package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)

	var totalMass float64 = 0.0
	// read through the file, print out the output
	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		totalMass += calculateFuel(mass)
	}

	fmt.Printf("%d\n", int(totalMass))
}

/** Calculate the fuel needed for a given mass */
func calculateFuel(mass float64) float64 {
	fuel := math.Floor(mass/3.0) - 2.0
	// Base case, negative mass
	if fuel <= 0 {
		return 0
	}

	// Calculate the mass needed for the fuel's fuel as well
	return fuel + calculateFuel(fuel)
}
