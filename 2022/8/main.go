package main

// Puzzle Source: https://adventofcode.com/2022/day/8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 99,99
// var grid = make([][]int, 99)
var height int = 5
var width int = 5
// var grid [99][99]int
// var visible [99][99]int
var grid [5][5]int
var visible [5][5]bool


func main() {
	// file, _ := os.Open("./input.txt")
	file, _ := os.Open("./example.txt")
	s := bufio.NewScanner(file)
	var y int = 0
	for s.Scan() {
		line := s.Text()
		for x, tree := range line {
			integer, err := strconv.Atoi(string(tree))
			if err != nil {
				log.Fatal("Conversion failed.")
			}
			grid[y][x] = integer
		}
		y++
	}
	fmt.Printf("%#v \n", grid)

	// for v := 0; v < 98; v++{
	// 	for h := 0; h < 98; h++{
	for h := 0; h < height - 1; h++{
		for v := 0; v < width - 1; v++{
			isVisible(h, v)
		}
	}

	fmt.Printf("%#v \n", visible)


}


func isVisible(y int, x int) {
	positive := true
	switch {
	case fromTop(x,y):
		visible[y][x] = positive
		fmt.Printf("%d at %d,%d is visible from Top \n",grid[y][x], y, x)
		break
	case fromRight(x,y):
		visible[y][x] = positive
		fmt.Printf("%d at %d,%d is visible from Right \n",grid[y][x], y, x)
		break
	case fromBottom(x,y):
		visible[y][x] = positive
		fmt.Printf("%d at %d,%d is visible from Bottom \n",grid[y][x], y, x)
		break
	case fromLeft(x,y):
		visible[y][x] = positive
		fmt.Printf("%d at %d,%d is visible from Left \n",grid[y][x], y, x)
		break
	default:
		visible[y][x] = false
		fmt.Printf("%d at %d,%d is not visible \n",grid[y][x], y, x)
	}

}
func fromTop(x int, y int) bool{
	for i := y; i > 0; i--{
		fmt.Printf("Is %d greater than %d ?\n", grid[i][x], grid[y][x])
		if grid[i][x] > grid[y][x]{
			return false
		}
	}
	return true
}
func fromRight(x int, y int) bool{
	for i := x; i < width; i++{
		if grid[y][i] > grid[y][x]{
			return false
		}
	}
	return true
}

func fromBottom(x int, y int) bool{
	for i := y; i < height; i++{
		if grid[i][x] > grid[y][x]{
			return false
		}
	}
	return true
}

func fromLeft(x int, y int)bool{
	for i := x; i > 0; i--{
		if grid[y][i] > grid[y][x]{
			return false
		}
	}
	return true
}
