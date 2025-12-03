package main

import (
	"2k25/common"
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

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		bankJoltage := 0
		numOfJoltageDigits := 12

		lastLeftIndex := 0

		for ; numOfJoltageDigits > 0; numOfJoltageDigits-- {
			lastBiggest := line[lastLeftIndex]
			for i := lastLeftIndex + 1; i < len(line)-numOfJoltageDigits+1; i++ {
				if lastBiggest < line[i] {
					lastBiggest = line[i]
					lastLeftIndex = i
				}
			}
			lastLeftIndex++

			bankJoltage += int(lastBiggest-'0') * common.Pow(10, numOfJoltageDigits-1)
		}

		sum += bankJoltage

	}

	fmt.Printf("total output joltage: %d\n", sum)
}
