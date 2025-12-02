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
}
