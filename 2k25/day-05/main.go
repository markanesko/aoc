package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	type tuple struct {
		low  int
		high int
	}

	freshIds := []tuple{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			break
		}

		low, _ := strconv.Atoi(strings.Split(line, "-")[0])
		high, _ := strconv.Atoi(strings.Split(line, "-")[1])

		freshIds = append(freshIds, tuple{
			low:  low,
			high: high,
		})

		sort.Slice(freshIds, func(i, j int) bool {
			if freshIds[i].low == freshIds[j].low {
				return freshIds[i].high < freshIds[j].high
			}
			return freshIds[i].low < freshIds[j].low
		})
	}

	freshIngredientsCount := freshIds[0].high - freshIds[0].low + 1

	currentHigh := freshIds[0].high

	for i := 1; i < len(freshIds); i++ {

		low := freshIds[i].low
		high := freshIds[i].high

		if high <= currentHigh {
			continue
		}

		if low > currentHigh {
			freshIngredientsCount += high - low + 1
			currentHigh = high
			continue
		}

		freshIngredientsCount += high - currentHigh
		currentHigh = high

	}

	fmt.Printf("fresh ingredients count: %d\n", freshIngredientsCount)

	counter := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		id, _ := strconv.Atoi(line)

		for _, r := range freshIds {
			if id >= r.low && id <= r.high {
				counter++
				break
			}
		}
	}

	fmt.Printf("number of fresh ingredients is: %d\n", counter)
}

func checkInBetween(a, b, c int) bool {
	if a >= b && a <= c {
		return true
	}

	return false
}
