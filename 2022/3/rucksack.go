package main

// Puzzle Source: https://adventofcode.com/2022/day/3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
var inBoth string = ""
func main(){

	// A-Z => 65-90 
	// a-z => 97-122

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("%s \n", "Something went wrong reading the file...")
	}

	sc := bufio.NewScanner(file)
	totalPriority := 0
	for sc.Scan() {
		line := sc.Text()
		firstComp := line[:len(line)/2]
		secComp := line[len(line)/2:]
		fmt.Printf("First Compartment: %v \n", firstComp)
		fmt.Printf("Second Compartment: %v \n", secComp)

		for i := 0; i < len(firstComp); i++{
			for j := 0; j < len(secComp); j++{
				if firstComp[i] == secComp[j]{
					totalPriority += checkPriority(secComp[j])
					fmt.Printf("Matching Item is: %v \n", string(secComp[j]))
				}
			}
		}
	}
	fmt.Printf("Total Priority: %d \n", totalPriority)

}

// need the %d of a byte to calc the priority
func checkPriority(char byte) int{
	resp := 0
	charInt, err := strconv.Atoi(fmt.Sprintf("%d",char))
	if err != nil{
		log.Fatal(err)
	}
	// A-Z => 65-90 
	if charInt >= 65 && charInt <= 90{ 
		resp = charInt - 64 + 26
		fmt.Printf("A-Z => 65-90 \n")
	}
	// a-z => 97-122
	if charInt >= 97 && charInt <= 122{
		resp = charInt - 96
		fmt.Printf("a-z => 97-122 \n")
		fmt.Printf("Priority added: %v \n", resp)
	}
	fmt.Printf("%d \n", resp)
	return resp
}


