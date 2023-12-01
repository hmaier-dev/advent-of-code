package main

// PuzzleSource: https://adventofcode.com/2022/day/6

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	file, _ := os.Open("./input.txt")

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		puzzleInput := sc.Text()
		findStartMarker(puzzleInput)
		// fmt.Printf("%v \n", sc.Text())
	}
}

func findStartMarker(puzzleInput string){
	for i := 0; i < len(puzzleInput)-13; i++{
		marker := puzzleInput[i:i+14]
		if startOfMessage(marker){
			fmt.Printf("Start: %v \n", i+14)
			break
		}
	}
}

func startOfPacket(seq string) bool{
	fmt.Printf("Seq: %#v \n", seq)
	for i := 0; i < 4; i++{
		for y := 0; y < 4; y++ {
			fmt.Printf("Comparing %s with %s \n", string(seq[i]), string(seq[y]))
			if seq[i] == seq[y] && y != i{
				return false
			}
		}
	}
	return true
}

func startOfMessage(seq string) bool{
	fmt.Printf("Seq: %#v \n", seq)
	for i := 0; i < 14; i++{
		for y := 0; y < 14; y++ {
			fmt.Printf("Comparing %s with %s \n", string(seq[i]), string(seq[y]))
			if seq[i] == seq[y] && y != i{
				return false
			}
		}
	}
	return true
}

