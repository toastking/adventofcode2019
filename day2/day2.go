// Advent of Code Day 2
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitCommas)

	var intCode program = readFileToArray(scanner)

	//Setup the program
	intCode[1] = 12
	intCode[2] = 2

	runProgram(intCode[0:])

	fmt.Printf("Result: %d\n", intCode[0])
}

type program []int64

// Reads in a file line by line and turns it into the array of integers
// for the "Intcode" program
func readFileToArray(scanner *bufio.Scanner) program {
	var intCode program
	for scanner.Scan() {
		var num string = scanner.Text()
		converted, err := strconv.ParseInt(num, 10, 64)

		if err != nil {
			panic(err)
		}

		intCode = append(intCode, converted)
	}
	return intCode
}

// SplitFunc to split a file based on commas
func splitCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchString := []byte(",")

	// If we are at the end of the file and there's no more data then we're done
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// If we are at the end of the file and there IS more data return it
	if atEOF {
		return len(data), data, nil
	}

	//If we reach a comma then advance to the next token
	if i := bytes.Index(data, searchString); i >= 0 {
		return i + len(searchString), data[0:i], nil
	}

	return
}

// Run the computation for each line of intcode
func runProgram(fullProgram program) {
	// Iterate over each line of the program
	for i, continueRunning := 0, true; continueRunning; i += 4 {
		opcode, firstOperand, secondOperand, resultLocation := getValuesForCalculation(fullProgram[i:i+4], fullProgram)
		switch opcode {
		case 1:
			fullProgram[resultLocation] = firstOperand + secondOperand
			continueRunning = true
		case 2:
			fullProgram[resultLocation] = firstOperand * secondOperand
			continueRunning = true
		case 99:
			// 99 means end the program
			continueRunning = false
		default:
			continueRunning = false
		}
	}
}

// Helper function to get the values needed for doing math
func getValuesForCalculation(lineOfCode []int64, fullProgram program) (int64, int64, int64, int64) {
	var (
		opcode         int64
		firstOperand   int64
		secondOperand  int64
		resultLocation int64
	)
	// Use the locations of the operands
	opcode = lineOfCode[0]
	if opcode != 99 {
		firstOperand = fullProgram[lineOfCode[1]]
		secondOperand = fullProgram[lineOfCode[2]]
		resultLocation = lineOfCode[3]
	}
	return opcode, firstOperand, secondOperand, resultLocation
}
