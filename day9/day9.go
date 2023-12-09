package main

import (
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

func main() {
	file, _ := os.Open("day9/input.txt")
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)
	s.Mode = scanner.ScanInts
	s.Whitespace ^= 1 << '\n'

	var resultPartOne, resultPartTwo int
	history := make([][]int, 1)

	isMinus := false
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case scanner.Int:
			num, _ := strconv.Atoi(s.TokenText())

			// I hate this but text/scanner does not add the minus even if the token's name is "Int"
			if isMinus {
				num *= -1
				isMinus = false
			}

			history[0] = append(history[0], num)
		case '-':
			isMinus = true
		case '\n':
			allZeros := false
			for i := 0; !allZeros; i++ {
				var difference []int
				allZeros = true

				for j := 0; j < len(history[i])-1; j++ {
					newValue := history[i][j+1] - history[i][j]
					difference = append(difference, newValue)

					if newValue != 0 {
						allZeros = false
					}
				}

				history = append(history, difference)
			}

			for i := len(history) - 2; i >= 0; i-- {
				// Add the calculated value at the end of the slice
				history[i] = append(history[i], history[i][len(history[i])-1]+history[i+1][len(history[i+1])-1])

				// Add the calculated value at the beginning of the slice
				history[i] = append([]int{history[i][0] - history[i+1][0]}, history[i]...)

				// Remove the last line from the "stack"
				history = history[:i+1]
			}

			resultPartOne += history[0][len(history[0])-1]
			resultPartTwo += history[0][0]

			history = make([][]int, 1)
		}
	}

	fmt.Printf("Part one: %d\nPart two: %d\n", resultPartOne, resultPartTwo)
}
