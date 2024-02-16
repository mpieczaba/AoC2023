package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type operation bool

const (
	greaterThan operation = true
	lessThan              = false
)

type rule struct {
	part      rune
	operation operation
	rating    int
	result    string
}

func main() {
	file, _ := os.Open("day19/input.txt")
	defer file.Close()

	workflows := make(map[string][]rule)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(scanner.Bytes()) > 0 {
		var workflowString string
		workflowString = strings.ReplaceAll(scanner.Text(), "{", " ")
		workflowString = strings.ReplaceAll(workflowString, "}", "")

		var name, rules string
		fmt.Sscanf(workflowString, "%s %s", &name, &rules)

		var newRules []rule

		for _, r := range strings.Split(rules, ",") {
			if strings.ContainsAny(r, "<>:") {
				var part rune
				var operation rune
				var rating int
				var result string

				fmt.Sscanf(r, "%c%c%d:%s", &part, &operation, &rating, &result)

				switch operation {
				case '>':
					newRules = append(newRules, rule{part, greaterThan, rating, result})
				case '<':
					newRules = append(newRules, rule{part, lessThan, rating, result})
				}
			} else {
				newRules = append(newRules, rule{result: r})
			}
		}

		workflows[name] = newRules
	}

	var partRatings []map[rune]int

	for scanner.Scan() {
		partRating := make(map[rune]int)

		partRatingString := strings.ReplaceAll(scanner.Text(), "{", "")
		partRatingString = strings.ReplaceAll(partRatingString, "}", "")

		for _, pr := range strings.Split(partRatingString, ",") {
			var part rune
			var rating int

			fmt.Sscanf(pr, "%c=%d", &part, &rating)

			partRating[part] = rating
		}

		partRatings = append(partRatings, partRating)
	}

	var resultPartOne int

	for _, partRating := range partRatings {
		key := "in"

		for key != "A" && key != "R" {
		ruleExit:
			for _, rule := range workflows[key] {
				if rule.part != 0 {
					switch rule.operation {
					case greaterThan:
						if partRating[rule.part] > rule.rating {
							key = rule.result

							break ruleExit
						}
					case lessThan:
						if partRating[rule.part] < rule.rating {
							key = rule.result

							break ruleExit
						}
					}
				} else {
					key = rule.result

					break ruleExit
				}
			}
		}

		if key == "A" {
			for _, ratings := range partRating {
				resultPartOne += ratings
			}
		}
	}

	fmt.Printf("Part one: %d\n", resultPartOne)
}
