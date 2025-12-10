package main

import (
	"2k25/common"
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

	coords := []coord{}

	xCoords := map[int][]int{}
	yCoords := map[int][]int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])

		coords = append(coords, coord{x: x, y: y})
		xCoords[x] = append(xCoords[x], y)
		yCoords[y] = append(yCoords[y], x)
	}

	for k := range xCoords {
		sort.Ints(xCoords[k])
	}
	for k := range yCoords {
		sort.Ints(yCoords[k])
	}

	largest := -1

	for i := 0; i < len(coords)-1; i++ {
		a := coords[i]
		for j := i + 1; j < len(coords); j++ {
			b := coords[j]
			p := area(a, b)
			if p > largest {
				largest = p
			}
		}
	}
	fmt.Printf("largest area is: %d\n", largest)

	for x, yArr := range xCoords {
		filledArray := []int{}
		for i := yArr[0]; i <= yArr[len(yArr)-1]; i++ {
			filledArray = append(filledArray, i)
		}
		xCoords[x] = filledArray
	}
	for y, xArr := range yCoords {
		filledArray := []int{}
		for i := xArr[0]; i <= xArr[len(xArr)-1]; i++ {
			filledArray = append(filledArray, i)
		}
		yCoords[y] = filledArray
	}

	largest = -1

	for i := 0; i < len(coords)-1; i++ {
		a := coords[i]
		for j := i + 1; j < len(coords); j++ {
			b := coords[j]

			if boxHasEdgeInside(a, b, xCoords, yCoords) {
				continue
			}

			p := area(a, b)
			if p > largest {
				largest = p
			}
		}
	}
	fmt.Printf("largest area is: %d\n", largest)
}

type coord struct {
	x, y int
}

func area(a, b coord) int {
	return (common.Abs(a.x-b.x) + 1) * (common.Abs(a.y-b.y) + 1)
}

func boxHasEdgeInside(a, b coord, xCoords, yCoords map[int][]int) bool {
	x1, x2 := a.x, b.x
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	y1, y2 := a.y, b.y
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	for x, ys := range xCoords {
		if x <= x1 || x >= x2 {
			continue
		}

		for i := 0; i+1 < len(ys); i++ {
			yLow := ys[i]
			yHigh := ys[i+1]

			if yHigh > y1 && yLow < y2 {
				return true
			}
		}
	}

	for y, xs := range yCoords {
		if y <= y1 || y >= y2 {
			continue
		}

		for i := 0; i+1 < len(xs); i++ {
			xLow := xs[i]
			xHigh := xs[i+1]

			if xHigh > x1 && xLow < x2 {
				return true
			}
		}
	}

	return false
}
