package main

import (
	"2k25/common"
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

	counter := 0

	fileScanner.Scan()
	line := fileScanner.Text()

	ranges := strings.Split(line, ",")

	for _, r := range ranges {
		idRange := strings.Split(r, "-")
		lower, _ := strconv.Atoi(idRange[0])
		upper, _ := strconv.Atoi(idRange[1])

		for i := lower; i <= upper; i++ {
			s := strconv.Itoa(i)
			mid := len(s) / 2
			if len(s)%2 != 0 {
				continue
			}

			left := s[:mid]
			right := s[mid:]

			if left == right {
				sillyPattern := fmt.Sprintf("%s%s", left, right)
				sillyNumber, _ := strconv.Atoi(sillyPattern)

				counter += sillyNumber
			}
		}
	}

	fmt.Printf("sillyIds: %d\n", counter)

	counter = 0
	set := make(map[int]struct{})

	for _, r := range ranges {
		idRange := strings.Split(r, "-")
		lower, _ := strconv.Atoi(idRange[0])
		upper, _ := strconv.Atoi(idRange[1])

		for i := lower; i <= upper; i++ {
			s := strconv.Itoa(i)
			sLen := len(s)
			divs := common.LowerHalfDivisors(sLen)

			sillyNumber := common.SameDigits(i)
			if _, ok := set[sillyNumber]; !ok {
				counter += sillyNumber
				set[sillyNumber] = struct{}{}
			}

			for _, d := range divs {
				sillyNumber := checkByNumDigits(i, d)
				if _, ok := set[sillyNumber]; !ok {
					counter += sillyNumber
					set[sillyNumber] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("even sillier ids: %d\n", counter)

}

func checkByNumDigits(n, m int) int {
	originalN := n

	increment := common.Pow(10, m)
	lastM := n % increment
	n = n / increment

	for n > 0 {
		if lastM != n%increment {
			return 0
		}
		lastM = n % increment
		n = n / increment
	}

	return originalN
}
