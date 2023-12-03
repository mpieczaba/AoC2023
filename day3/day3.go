package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"text/scanner"
)

func main() {
	file, _ := os.Open("day3/input.txt")
	defer file.Close()

	// TODO: Make it use bool mask instead
	var mask [140][140]*int

	var s scanner.Scanner
	s.Init(file)
	s.Mode = scanner.ScanInts

	// Fill the mask
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		num, err := strconv.Atoi(s.TokenText())
		if err != nil {
			continue
		}

		x := s.Position.Line - 1
		y := s.Position.Column - 1

		for digitPosition := range s.TokenText() {
			mask[x][y+digitPosition] = &num
		}
	}

	// Seek the file and re-init the scanner
	file.Seek(0, io.SeekStart)
	s.Init(file)
	s.Mode = scanner.ScanChars

	var resultPartOne, resultPartTwo int

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		// Continue  if tok is not a symbol
		if tok >= '0' && tok <= '9' || tok == '.' {
			continue
		}

		x := s.Position.Line - 1
		y := s.Position.Column - 1

		var parts []int

		// Iterate through tok's neighbors
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if mask[x+i][y+j] != nil && *mask[x+i][y+j] != 0 {
					resultPartOne += *mask[x+i][y+j]

					parts = append(parts, *mask[x+i][y+j])

					*mask[x+i][y+j] = 0
				}
			}
		}

		// Add the gear ratio to the part two result
		if tok == '*' && len(parts) == 2 {
			resultPartTwo += parts[0] * parts[1]
		}
	}

	fmt.Printf("Part one: %d\nPart two: %d\n", resultPartOne, resultPartTwo)
}
