package main

import (
	"2k25/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		return
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	counter := 50
	nullCounter := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			break
		}
		operation := line[0]
		operand := line[1:]

		operandValue, err := strconv.Atoi(operand)

		if err != nil {
			fmt.Printf("error formating operand value %s", operand)
			break
		}
		numberOfCertainClicks := operandValue / 100

		nullCounter += numberOfCertainClicks
		operandValue = common.Mod(operandValue, 100)

		if string(operation) == "L" {
			if operandValue > counter && counter != 0 {
				nullCounter++
			}
			counter -= operandValue
			counter = common.Mod(counter, 100)
		}

		if string(operation) == "R" {
			if counter+operandValue > 100 {
				nullCounter++
			}
			counter += operandValue
			counter = common.Mod(counter, 100)
		}

		if counter == 0 {
			nullCounter++
		}
	}

	fmt.Printf("password is: %d\n", nullCounter)
}
