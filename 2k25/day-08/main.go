package main

import (
	"2k25/common"
	"bufio"
	"fmt"
	"math"
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

	for fileScanner.Scan() {
		line := fileScanner.Text()
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		z, _ := strconv.Atoi(strings.Split(line, ",")[2])

		coords = append(coords, coord{
			x: x,
			y: y,
			z: z,
		})
	}

	distances := []distancePair{}
	for i := 0; i < len(coords)-1; i++ {
		a := coords[i]
		for j := i + 1; j < len(coords); j++ {
			b := coords[j]
			d := calculateDistance(a, b)
			distances = append(distances, distancePair{
				a: a,
				b: b,
				d: d,
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].d < distances[j].d
	})

	pairs := map[pair]struct{}{}
	circuits := []map[coord]struct{}{}
	num := 0

	for _, d := range distances {
		if inIncludedPairs(pair{a: d.a, b: d.b}, pairs) {
			continue
		}

		num++
		circuits = addPair(pair{a: d.a, b: d.b}, circuits)
		pairs[pair{a: d.a, b: d.b}] = struct{}{}

		if num == 1000 {
			break
		}
	}

	sizes := []int{}
	for _, c := range circuits {
		sizes = append(sizes, len(c))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	fmt.Printf("product is: %d\n", sizes[0]*sizes[1]*sizes[2])

	pairs = map[pair]struct{}{}
	circuits = []map[coord]struct{}{}

	for _, d := range distances {
		if inIncludedPairs(pair{a: d.a, b: d.b}, pairs) {
			continue
		}

		num++
		circuits = addPair(pair{a: d.a, b: d.b}, circuits)
		pairs[pair{a: d.a, b: d.b}] = struct{}{}

		if len(circuits[0]) == len(coords) {
			fmt.Printf("a.x * b.x = %d\n", d.a.x*d.b.x)
			break
		}
	}

}

type coord struct {
	x, y, z int
}

type pair struct {
	a, b coord
}

type distancePair struct {
	a, b coord
	d    float64
}

func inIncludedPairs(p pair, pairs map[pair]struct{}) bool {
	p1 := p
	p2 := pair{
		a: p.b,
		b: p.a,
	}

	if _, ok := pairs[p1]; ok {
		return true
	}

	if _, ok := pairs[p2]; ok {
		return true
	}

	return false
}

func addPair(p pair, circuits []map[coord]struct{}) []map[coord]struct{} {

	iA := -1
	iB := -1

	for i := 0; i < len(circuits); i++ {
		if _, aOk := circuits[i][p.a]; aOk {
			iA = i
		}
		if _, bOk := circuits[i][p.b]; bOk {
			iB = i
		}
	}

	if iA == -1 && iB == -1 {
		c := map[coord]struct{}{}
		c[p.a] = struct{}{}
		c[p.b] = struct{}{}
		circuits = append(circuits, c)
		return circuits
	}

	if iA != -1 && iB == -1 {
		circuits[iA][p.b] = struct{}{}
		return circuits
	}

	if iB != -1 && iA == -1 {
		circuits[iB][p.a] = struct{}{}
		return circuits
	}

	if iA == iB {
		return circuits
	}

	if iB < iA {
		iA, iB = iB, iA
	}
	for k := range circuits[iB] {
		circuits[iA][k] = struct{}{}
	}

	circuits = append(circuits[:iB], circuits[iB+1:]...)

	return circuits
}

func calculateDistance(a, b coord) float64 {
	return math.Sqrt(
		float64(
			common.IntSum(
				common.Pow(a.x-b.x, 2),
				common.Pow(a.y-b.y, 2),
				common.Pow(a.z-b.z, 2),
			),
		),
	)
}
