package main

import (
	"bufio"
	"fmt"
	"os"
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

	lines := [][]byte{}

	for fileScanner.Scan() {
		b := fileScanner.Bytes()
		line := make([]byte, len(b))
		copy(line, b)
		lines = append(lines, line)
	}

	lineLen := len(lines[0])
	allLinesLen := len(lines)

	total := 0
	totalWithRemove := 0
	shouldRun := true

	type tuple struct {
		row int
		col int
	}

	removed := []tuple{}

	for shouldRun {
		for j, line := range lines {
			for i, v := range line {
				if v == '@' {
					counter := 0
					if i-1 >= 0 {
						if line[i-1] == byte('@') {
							counter++
						}
						if j-1 >= 0 && lines[j-1][i-1] == byte('@') {
							counter++
						}
						if j+1 < allLinesLen && lines[j+1][i-1] == byte('@') {
							counter++
						}
					}

					if i+1 < lineLen {
						if line[i+1] == byte('@') {
							counter++
						}
						if j-1 >= 0 && lines[j-1][i+1] == byte('@') {
							counter++
						}
						if j+1 < allLinesLen && lines[j+1][i+1] == byte('@') {
							counter++
						}
					}

					if j-1 >= 0 && lines[j-1][i] == byte('@') {
						counter++
					}

					if j+1 < allLinesLen && lines[j+1][i] == byte('@') {
						counter++
					}

					if counter < 4 {
						removed = append(removed, tuple{j, i})
						total++
					}
				}
			}
		}

		for _, r := range removed {
			col := r.col
			row := r.row

			lines[row][col] = byte('.')

		}

		removed = []tuple{}

		if total == 0 {
			shouldRun = false
		}

		totalWithRemove += total
		total = 0

	}

	fmt.Printf("total paper rolls to move: %d\n", totalWithRemove)
}
