package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("files/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)

	var totalMass float64 = 0.0
	// Read through the file, print out the output
	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		totalMass += calculateFuel(mass)
	}
	fmt.Printf("%f", totalMass)
}

/** Calculate the fuel needed for a given mass */
func calculateFuel(mass float64) float64 {
	return math.Round(mass/3.0) - 2.0
}
