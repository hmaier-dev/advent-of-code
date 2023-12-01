package main

// https://adventofcode.com/2022/day/2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("%s", "Something went wrong reading the file...")
	}

	// Part 1
	// --------------------------------------------------------
	// // opponent
	// oRock := "A"
	// oPaper := "B"
	// oScissors := "C"

	// // myself
	// mRock := "X"
	// mPaper := "Y"
	// mScissors := "Z"
	
	// Score consist of shape and outcome
	// 1 for Rock
	// 2 for Paper
	// 3 for Scissors
	// 0 for Losing
	// 3 for Draw
	// 6 for Winning

	// Part 2
	// --------------------------------------------------------
	// X => need to lose
	// Y => need to draw
	// Z => need to win


	sc := bufio.NewScanner(file)
	var myScore int = 0
	for sc.Scan() {
		line := sc.Text()
		round := strings.Split(line, " ")// [opponent, myself]
		myScore += calcScore(round[0], round[1])
	}
	fmt.Printf("Total score is %v \n", myScore)
}

func calcScore(opponent string, outcome string) int {
 	var score int = 0
	switch true {
	// Rock
	case opponent == "A" && outcome == "X": // loose/scissor
		score += 3
	case opponent == "A" && outcome == "Y": // draw/rock
		score += 3 + 1
	case opponent == "A" && outcome == "Z": // win/paper
		score += 2 + 6
	// Paper
	case opponent == "B" && outcome == "X": // loose/rock
		score += 1
	case opponent == "B" && outcome == "Y": // draw/paper
		score += 3 + 2
	case opponent == "B" && outcome == "Z": // win/scissors
		score += 6 + 3
	// Scissors
	case opponent == "C" && outcome == "X": // loose/paper
		score += 2
	case opponent == "C" && outcome == "Y": // draw/scissors
		score += 3 + 3
	case opponent == "C" && outcome == "Z": // win/rock
		score += 6 + 1
	}
	return score

}




// func calcScore(opponent string, myself string) int{
// 	var score int = 0
// 	switch true {
// 		case myself == "A": //rock
// 			score += 1
// 		case myself == "B": //paper
// 			score += 2
// 		case myself == "C": //scissors
// 			score += 3	
// 	}
// 	switch true{
// 		case myself == opponent: //draw
// 			score += 3	
// 		case myself == "A" && opponent == "C": //winning with Rock
// 			score += 6 
// 		case myself == "B" && opponent == "A": //winning with Paper
// 			score += 6 
// 		case myself == "C" && opponent == "B": //winning with Scissors
// 			score += 6 
// 	}
// 	fmt.Printf("%v \n", score)
// 	
// 	return score
// }

