package main

import (
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

func main() {
	file, _ := os.Open("day2/input.txt")
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)
	s.Whitespace ^= 1 << '\n'

	gameID := 1
	var prev string
	var maxRed, maxGreen, maxBlue int
	var resultPartOne, resultPartTwo int

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch s.TokenText() {
		case "red":
			if num, _ := strconv.Atoi(prev); num > maxRed {
				maxRed = num
			}
		case "green":
			if num, _ := strconv.Atoi(prev); num > maxGreen {
				maxGreen = num
			}
		case "blue":
			if num, _ := strconv.Atoi(prev); num > maxBlue {
				maxBlue = num
			}
		case "\n":
			// Add game ID to the part one result if RGB values satisfy given conditions
			if !(maxRed > 12 || maxGreen > 13 || maxBlue > 14) {
				resultPartOne += gameID
			}

			// Calculate powers of the minimum set of cubes and add it to the part two result
			resultPartTwo += maxRed * maxGreen * maxBlue

			// Zero out RGB variables and increment the game ID
			maxRed, maxGreen, maxBlue = 0, 0, 0
			gameID++
		}

		prev = s.TokenText()
	}

	fmt.Printf("Part one: %d\nPart two: %d\n", resultPartOne, resultPartTwo)
}
