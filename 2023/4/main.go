package main

import (
	"bufio"
	"os"
)

func main() {

	file, _ := os.Open("./input.txt")
	// file, err := os.Open("./exmaple.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

	}
}
