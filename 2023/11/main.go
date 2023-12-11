package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var universe [][]string
var expandedUniverse [][]string

func main() {
	// file, err := os.Open("./input.txt")
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatalf("reading the file failed...")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%#v \n", line)
		buildTheUniverse(line)
	}
	expandTheUniverse()
}

func buildTheUniverse(line string) {
	var row []string
	for _, s := range line {
		row = append(row, string(s))
	}
	universe = append(universe, row)
}

func expandTheUniverse() {
	// expand horizontally
	for x, row := range universe {
		for i, c := range row {
			if c == "#" {
				expandedUniverse = append(expandedUniverse, row)
				continue // don't duplicate
			}
			if checkVertical(x) == true{
				//
			}
			if i == len(row)-1 {
				expandedUniverse = append(expandedUniverse, row)
				expandedUniverse = append(expandedUniverse, row)
				//duplicate

			}
		}
	}

}

func checkVertical(x int)bool{
	// expand vertically
	h := len(universe)
	for x := 0; x < h; x++ {
		for y := 0; y < h; y++ {
			if universe[y][x] == "#" {
				continue // don't duplicate
			}
			if y == h-1 {
				//duplicate

			}

		}

	}
	return false
}
