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
			divs := lowerHalfDivisors(sLen)

			sillyNumber := sameDigits(i)
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

// 1 * 36 = 36
// 2 * 18 = 36
// 3 * 12 = 36
// 4 * 9  = 36
// 6 * 6  = 36
func divisors(n int) []int {
	small := []int{}
	large := []int{}

	// first part till sqrt(n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			small = append(small, i)

			// make sure that there are no duplicates
			if i != n/i {
				large = append(large, n/i)
			}
		}
	}

	for i := len(large) - 1; i >= 0; i-- {
		small = append(small, large[i])
	}

	return small
}

func lowerHalfDivisors(n int) []int {
	small := []int{}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			small = append(small, i)
		}
	}

	return small
}

func sameDigits(n int) int {
	if n < 10 {
		return 0
	}
	lastDigit := n % 10
	withoutLast := n / 10

	for withoutLast > 0 {
		if lastDigit != withoutLast%10 {
			return 0
		}

		lastDigit = withoutLast % 10
		withoutLast = withoutLast / 10
	}

	return n
}

func checkByNumDigits(n, m int) int {
	originalN := n

	increment := pow(10, m)
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

func pow(a, b int) int {
	res := 1

	for b > 0 {
		res *= a
		b--
	}

	return res
}
