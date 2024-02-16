package main

import (
	"bufio"
	"fmt"
	"os"
)

type digPlanEntry struct {
	direction rune
	depth     int
	color     int
}

type position struct {
	x, y int
}

func main() {
	file, _ := os.Open("day18/input.txt")
	defer file.Close()

	var digPlan []position
	digPlan = append(digPlan, position{0, 0})

	var digPlan2 []position
	digPlan2 = append(digPlan2, position{0, 0})

	circumference := 0
	circumference2 := 0

	s := bufio.NewScanner(file)
	for s.Scan() {
		var entry digPlanEntry
		previousEntry := digPlan[len(digPlan)-1]
		previousEntry2 := digPlan2[len(digPlan2)-1]

		fmt.Sscanf(s.Text(), "%c %d (#%x)", &entry.direction, &entry.depth, &entry.color)

		switch entry.direction {
		case 'U':
			digPlan = append(digPlan, position{previousEntry.x, previousEntry.y - entry.depth})
		case 'R':
			digPlan = append(digPlan, position{previousEntry.x + entry.depth, previousEntry.y})
		case 'D':
			digPlan = append(digPlan, position{previousEntry.x, previousEntry.y + entry.depth})
		case 'L':
			digPlan = append(digPlan, position{previousEntry.x - entry.depth, previousEntry.y})
		}

		depth := entry.color >> 4
		direction := entry.color - (depth << 4)

		switch direction {
		case 3:
			digPlan2 = append(digPlan2, position{previousEntry2.x, previousEntry2.y - depth})
		case 0:
			digPlan2 = append(digPlan2, position{previousEntry2.x + depth, previousEntry2.y})
		case 1:
			digPlan2 = append(digPlan2, position{previousEntry2.x, previousEntry2.y + depth})
		case 2:
			digPlan2 = append(digPlan2, position{previousEntry2.x - depth, previousEntry2.y})
		}

		circumference += entry.depth
		circumference2 += depth
	}

	// Calculate the polygon area with Shoelace formula
	area := 0
	j := len(digPlan) - 1
	for i := 0; i < len(digPlan); i++ {
		area += (digPlan[j].x + digPlan[i].x) * (digPlan[j].y - digPlan[i].y)
		j = i
	}

	resultPartOne := area/2 - circumference/2 - 1
	resultPartOne *= -1

	area2 := 0
	j = len(digPlan2) - 1
	for i := 0; i < len(digPlan2); i++ {
		area2 += (digPlan2[j].x + digPlan2[i].x) * (digPlan2[j].y - digPlan2[i].y)
		j = i
	}

	resultPartTwo := area2/2 - circumference2/2 - 1
	resultPartTwo *= -1

	fmt.Printf("Part one: %d\n", resultPartOne)
	fmt.Printf("Part two: %d\n", resultPartTwo)
}
