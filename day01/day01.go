package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var digits = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	file, _ := os.Open("day1/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var resultPartOne, resultPartTwo int = 0, 0

	// Loop through each line of the file
	for scanner.Scan() {
		// Get the first digit literal from the scanned line
		firstIndex := strings.IndexAny(scanner.Text(), "0123456789")
		first := int(scanner.Text()[firstIndex] - '0')

		// Get the last digit literal from the scanned line
		lastIndex := strings.LastIndexAny(scanner.Text(), "0123456789")
		last := int(scanner.Text()[lastIndex] - '0')

		// Calculate the part one result
		resultPartOne += 10*first + last

		// Iterate through the digits array
		for i, name := range digits {
			// Change the first digit if the name is before the first digit literal or any other name from the digits array
			if index := strings.Index(scanner.Text(), name); index < firstIndex && index >= 0 {
				firstIndex = index
				first = i + 1
			}

			// Change the last digit if the name is after the last digit literal or any other name from the digits array
			if index := strings.LastIndex(scanner.Text(), name); index > lastIndex && index >= 0 {
				lastIndex = index
				last = i + 1
			}
		}

		// Calculate the part two result
		resultPartTwo += 10*first + last
	}

	fmt.Printf("Part one: %d\n", resultPartOne)
	fmt.Printf("Part two: %d\n", resultPartTwo)
}
