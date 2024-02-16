package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"text/scanner"
)

func main() {
	file, _ := os.Open("day6/input.txt")
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)
	s.Mode = scanner.ScanInts
	s.Whitespace ^= 1 << '\n'

	var time []int
	var distance []int

	resultPartOne, resultPartTwo := 1, 0

	isTime := true

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case scanner.Int:
			if num, _ := strconv.Atoi(s.TokenText()); isTime {
				time = append(time, num)
			} else {
				distance = append(distance, num)
			}
		case '\n':
			isTime = false
		}
	}

	// math package sucks

	for i := 0; i < len(time); i++ {
		t := float64(time[i])
		d := float64(distance[i])

		deltaRoot := math.Sqrt(t*t - 4*d)

		x1 := (-t - deltaRoot) / (-2)
		x2 := (-t + deltaRoot) / (-2)

		if x1 == math.Floor(x1) {
			x1--
		}

		if x2 == math.Ceil(x2) {
			x2++
		}

		resultPartOne *= int(math.Floor(x1)) - int(math.Ceil(x2)) + 1
	}

	var tC string
	for _, t := range time {
		tC += strconv.Itoa(t)
	}

	var dC string
	for _, d := range distance {
		dC += strconv.Itoa(d)
	}

	tCF, _ := strconv.ParseFloat(tC, 64)
	dCF, _ := strconv.ParseFloat(dC, 64)

	t := tCF
	d := dCF

	deltaRoot := math.Sqrt(t*t - 4*d)

	x1 := (-t - deltaRoot) / (-2)
	x2 := (-t + deltaRoot) / (-2)

	if x1 == math.Floor(x1) {
		x1--
	}

	if x2 == math.Ceil(x2) {
		x2++
	}

	resultPartTwo = int(math.Floor(x1)) - int(math.Ceil(x2)) + 1

	fmt.Printf("Part one: %d\n", resultPartOne)
	fmt.Printf("Part two: %d\n", resultPartTwo)
}
