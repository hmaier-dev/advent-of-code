package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var numberCollection []int
var sum int = 0

func main(){
	// file, err := os.Open("./example.txt")
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./example2.txt")
	if err != nil{
		log.Fatalf("reading the file failed...")
	}
	scanner :=  bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Printf("%v ", line)
		numbers :=  getNumbers(line)	
		numberCollection = append(numberCollection, addToCollection(numbers))
	}
	for i, elem := range numberCollection{
		sum += elem
		fmt.Printf("Sum: %v, %#v %d\n",sum, numberCollection[i],i)
	}

	fmt.Printf("%v \n", sum)
}

func getNumbers(str string)[]string{
	var asNumbers []string
	// regInt := regexp.MustCompile(`\d`)
	regAll := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)|(\d){1}`)
	matches := regAll.FindAllString(str, -1)
	fmt.Printf("%#v ", matches)
	for _, elem := range matches{
		asNumbers = append(asNumbers, literalToNumber(elem))
	}
	fmt.Printf("%#v ", asNumbers)
	return asNumbers
}

func addToCollection(arr []string)int{
	first := arr[0]
	second := arr[len(arr)-1]
	both := first+second
	integer, _ := strconv.Atoi(both)
	fmt.Printf("%#v \n", integer)
	return integer
}

func literalToNumber(str string)string{
	asNum := str
	switch str{
	case "one":
		asNum = "1"
	case "two":
		asNum = "2"
	case "three":
		asNum = "3"
	case "four":
		asNum = "4"
	case "five":
		asNum = "5"
	case "six":
		asNum = "6"
	case "seven":
		asNum = "7"
	case "eight":
		asNum = "8"
	case "nine":
		asNum = "9"
	}
	return asNum

}

