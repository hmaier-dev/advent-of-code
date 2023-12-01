package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("%s", "Something went wrong reading the file...")
	}

	sc := bufio.NewScanner(file)
	current := 0
	calories := 0
	arr := []int{}
	for sc.Scan() {
		b := sc.Bytes()
		fmt.Printf("%v \n", b)
		current += 1
		if len(b) == 0 {
			arr = append(arr, calories)
			fmt.Printf("%v \n", calories)
			calories = 0
		} else {
			num, _ := strconv.Atoi(string(b))
			calories += num
		}
	}
	sort.Ints(arr)
	total := arr[len(arr)-1] + arr[len(arr)-2] + arr[len(arr)-3]
	fmt.Printf("Total: %v \n", total)

}
