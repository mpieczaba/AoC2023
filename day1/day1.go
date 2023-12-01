package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var digits = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func main() {
	file, _ := os.Open("day1/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var resultPartOne, resultPartTwo int = 0, 0

	// Loop through each line of the file
	for scanner.Scan() {
		// Get the first digit literal from the scanned line
		firstIndex := strings.IndexAny(scanner.Text(), "0123456789")
		first := scanner.Text()[firstIndex]

		// Get the last digit literal from the scanned line
		lastIndex := strings.LastIndexAny(scanner.Text(), "0123456789")
		last := scanner.Text()[lastIndex]

		// Concatenate digits and add the parsed number to the part one result
		parsedNumber, _ := strconv.ParseInt(fmt.Sprintf("%c%c", first, last), 10, 32)
		resultPartOne += int(parsedNumber)

		// Iterate through the digits map
		for name, digit := range digits {
			// Change the first digit if the name is before the first digit literal or any other name from the map
			if index := strings.Index(scanner.Text(), name); index < firstIndex && index >= 0 {
				firstIndex = index
				first = digit
			}

			// Change the last digit if the name is after the last digit literal or any other name from the map
			if index := strings.LastIndex(scanner.Text(), name); index > lastIndex && index >= 0 {
				lastIndex = index
				last = digit
			}
		}

		// Concatenate digits and add the parsed number to the part two result
		parsedNumber, _ = strconv.ParseInt(fmt.Sprintf("%c%c", first, last), 10, 32)
		resultPartTwo += int(parsedNumber)
	}

	fmt.Printf("Part one: %d\nPart two: %d\n", resultPartOne, resultPartTwo)
}
