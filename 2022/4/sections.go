package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("%s \n", "Something went wrong reading the file...")
	}

	overlapSum := 0
	notOverlapSum := 0
	all := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		re := regexp.MustCompile(`\d+`)
		match := re.FindAllString(line, -1)
		fmt.Printf("%v \n", match)

		firstStart, err := strconv.Atoi(match[0])
		firstEnd, err := strconv.Atoi(match[1])
		secStart, err := strconv.Atoi(match[2])
		secEnd, err := strconv.Atoi(match[3])
		if err != nil {
			fmt.Println("str to int conversion failed...")
		}
		// switch true {
		// // Anfang
		// case firstStart <= secStart && firstEnd <= secEnd && firstEnd >= secStart:
		// 	overlapSum += 1
		// // Ende
		// case firstStart >= secStart && firstEnd >= secEnd && firstStart <= secEnd:
		// 	overlapSum += 1
		// // Anfang
		// case secStart <= firstStart && secEnd <= firstEnd && secEnd >= firstStart:
		// 	overlapSum += 1
		// // Ende
		// case secStart >= firstStart && secEnd >= firstEnd && secStart <= firstEnd:
		// 	overlapSum += 1
		// default:
		// 	fmt.Println("Does not overlap.")
		// 	notOverlapSum += 1
		// }

		if (firstEnd >= secStart && firstEnd <= secEnd) || (secEnd >= firstStart && secEnd <= firstEnd) {
			overlapSum++
		} else {
			fmt.Println("Does not overlap.")
			notOverlapSum += 1

		}
		all += 1

		fmt.Println("---------------------------------------------")
		// // seconds contains first
		// if firstStart >= secStart && firstEnd <= secEnd {
		// 	overlapSum += 1
		// }
		// // first contains second
		// if firstStart <= secStart && firstEnd >= secEnd {
		// 	overlapSum += 1
		// }
	}
	fmt.Printf("Overlap: %v \n", overlapSum)
	fmt.Printf("Not: %v \n", notOverlapSum)
	fmt.Printf("All: %v \n", all)

}
