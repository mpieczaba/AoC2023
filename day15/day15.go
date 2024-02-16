package main

import (
	"fmt"
	"os"
	"text/scanner"
)

func main() {
	file, _ := os.Open("day15/input.txt")
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)
	s.Mode = scanner.ScanChars
	s.Whitespace ^= 1 << '\n'

	var resultPartOne int
	var hash int

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case ',', '\n':
			resultPartOne += hash
			hash = 0
		default:
			hash += int(tok)
			hash *= 17
			hash %= 256
		}
	}

	fmt.Printf("Part one: %d\n", resultPartOne)
}
