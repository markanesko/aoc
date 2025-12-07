package main

import (
	"bufio"
	"fmt"
	"os"
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

	lines := [][]byte{}

	for fileScanner.Scan() {
		b := fileScanner.Bytes()
		line := make([]byte, len(b))
		copy(line, b)
		lines = append(lines, line)
	}

	type tuple struct {
		index     int
		operation string
	}

	sum := 0
	lastIndex := 0
	for i := 1; i < len(lines[4]); i++ {
		operation := lines[4][i]

		if operation == 0x20 {
			continue
		}

		lastOperation := lines[4][lastIndex]
		a, _ := strconv.Atoi(strings.ReplaceAll(string(lines[0][lastIndex:i]), " ", ""))
		b, _ := strconv.Atoi(strings.ReplaceAll(string(lines[1][lastIndex:i]), " ", ""))
		c, _ := strconv.Atoi(strings.ReplaceAll(string(lines[2][lastIndex:i]), " ", ""))
		d, _ := strconv.Atoi(strings.ReplaceAll(string(lines[3][lastIndex:i]), " ", ""))

		if lastOperation == '*' {
			sum += a * b * c * d
		} else {
			sum += a + b + c + d
		}
		lastIndex = i
	}

	lineLen := len(lines[4])
	lastOperation := lines[4][lastIndex]
	a, _ := strconv.Atoi(strings.ReplaceAll(string(lines[0][lastIndex:lineLen]), " ", ""))
	b, _ := strconv.Atoi(strings.ReplaceAll(string(lines[1][lastIndex:lineLen]), " ", ""))
	c, _ := strconv.Atoi(strings.ReplaceAll(string(lines[2][lastIndex:lineLen]), " ", ""))
	d, _ := strconv.Atoi(strings.ReplaceAll(string(lines[3][lastIndex:lineLen]), " ", ""))

	if lastOperation == '*' {
		sum += a * b * c * d
	} else {
		sum += a + b + c + d
	}

	fmt.Printf("total sum: %d\n", sum)

	sum = 0
	lastIndex = 0
	for i := 1; i < len(lines[4]); i++ {
		operation := lines[4][i]

		if operation == 0x20 {
			continue
		}

		lastOperation := lines[4][lastIndex]
		a := string(lines[0][lastIndex:i])
		b := string(lines[1][lastIndex:i])
		c := string(lines[2][lastIndex:i])
		d := string(lines[3][lastIndex:i])

		fmt.Printf("a: %s, b: %s, c: %s, d: %s\n", a, b, c, d)
		if lastOperation == '*' {
			product := 1
			for j := 0; j < i-lastIndex-1; j++ {
				v, _ := strconv.Atoi(strings.ReplaceAll(fmt.Sprintf("%c%c%c%c", a[j], b[j], c[j], d[j]), " ", ""))
				fmt.Printf("v: %d, ", v)
				product *= v
			}
			fmt.Println()

			sum += product
		} else {
			for j := 0; j < i-lastIndex-1; j++ {
				v, _ := strconv.Atoi(strings.ReplaceAll(fmt.Sprintf("%c%c%c%c", a[j], b[j], c[j], d[j]), " ", ""))
				fmt.Printf("v: %d, ", v)
				sum += v
			}
			fmt.Println()
		}

		lastIndex = i
	}

	lastOperation = lines[4][lastIndex]
	aStr := string(lines[0][lastIndex:lineLen])
	bStr := string(lines[1][lastIndex:lineLen])
	cStr := string(lines[2][lastIndex:lineLen])
	dStr := string(lines[3][lastIndex:lineLen])

	fmt.Printf("a: %s, b: %s, c: %s, d: %s\n", aStr, bStr, cStr, dStr)

	if lastOperation == '*' {
		product := 1
		for j := 0; j < len(aStr); j++ {
			v, _ := strconv.Atoi(strings.ReplaceAll(fmt.Sprintf("%c%c%c%c", aStr[j], bStr[j], cStr[j], dStr[j]), " ", ""))
			fmt.Printf("v: %d, ", v)
			product *= v
		}
		fmt.Println()

		sum += product
	} else {
		for j := 0; j < len(aStr); j++ {
			v, _ := strconv.Atoi(strings.ReplaceAll(fmt.Sprintf("%c%c%c%c", aStr[j], bStr[j], cStr[j], dStr[j]), " ", ""))
			fmt.Printf("v: %d, ", v)
			sum += v
		}
		fmt.Println()
	}

	fmt.Printf("new total sum: %d\n", sum)

}
