package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error opening file: %s", err.Error())
		return
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	tachyons := []int{}

	fileScanner.Scan()
	line := fileScanner.Text()

	startIndex := strings.Index(line, "S")
	fmt.Printf("start index: %d\n", startIndex)

	tachyons = append(tachyons, startIndex)
	counter := 0

	lines := []string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
		mutated := []int{}

		for _, tId := range tachyons {
			if line[tId] == '.' && !slices.Contains(mutated, tId) {
				mutated = append(mutated, tId)
				continue
			}

			if line[tId] == '^' {
				counter++
				left := tId - 1
				right := tId + 1

				if left >= 0 && !slices.Contains(mutated, left) {
					mutated = append(mutated, left)
				}

				if right < len(line) && !slices.Contains(mutated, right) {
					mutated = append(mutated, right)
				}

			}
		}

		tachyons = mutated
	}

	fmt.Printf("number of splits: %d\n", counter)

	calculated := map[tuple]int{}

	timeSplited := calculate(1, startIndex, lines, calculated)
	fmt.Printf("timeSplited: %d\n", timeSplited)

}

type tuple struct {
	x int
	y int
}

func calculate(level, index int, tachyons []string, calculated map[tuple]int) int {
	result := 0

	v := tachyons[level][index]

	if level == len(tachyons)-1 {
		return 1
	}

	if v == '.' {
		result += calculate(level+1, index, tachyons, calculated)
	}

	if _, ok := calculated[tuple{level, index}]; ok {
		return calculated[tuple{level, index}]
	}

	if v == '^' {
		left := index - 1
		right := index + 1

		if left < 0 {
			return 1
		} else {
			result += calculate(level+1, left, tachyons, calculated)
		}

		if right >= len(tachyons[0]) {
			return 1
		} else {
			result += calculate(level+1, right, tachyons, calculated)
		}

	}
	calculated[tuple{level, index}] = result

	return result
}
