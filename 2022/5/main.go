package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Order struct {
	quantity int
	from     int
	to       int
}


func main() {
	var cargoBay [][]string
	cargoBay = parseCargoInput()
	// fmt.Printf("%#v \n", cargoBay)
	fmt.Println("Starting position!")
	fmt.Printf("1: %#v \n", cargoBay[0])
	fmt.Printf("2: %#v \n", cargoBay[1])
	fmt.Printf("3: %#v \n", cargoBay[2])
	fmt.Printf("4: %#v \n", cargoBay[3])
	fmt.Printf("5: %#v \n", cargoBay[4])
	fmt.Printf("6: %#v \n", cargoBay[5])
	fmt.Printf("7: %#v \n", cargoBay[6])
	fmt.Printf("8: %#v \n", cargoBay[7])
	fmt.Printf("9: %#v \n", cargoBay[8])

	file := readInput("./movements.txt")
	sc := bufio.NewScanner(file)
	count := 0
	for sc.Scan() {
		line := sc.Text()
		currOrder := parseMovements(line)
		executeMovements(&cargoBay,currOrder)
		fmt.Println("---------------------------------------------")
		fmt.Printf("Order: %#v \n",currOrder)
		fmt.Printf("1: %#v \n", cargoBay[0])
		fmt.Printf("2: %#v \n", cargoBay[1])
		fmt.Printf("3: %#v \n", cargoBay[2])
		fmt.Printf("4: %#v \n", cargoBay[3])
		fmt.Printf("5: %#v \n", cargoBay[4])
		fmt.Printf("6: %#v \n", cargoBay[5])
		fmt.Printf("7: %#v \n", cargoBay[6])
		fmt.Printf("8: %#v \n", cargoBay[7])
		fmt.Printf("9: %#v \n", cargoBay[8])
		count++
		// if count > 2 {
		// 	break
		// }
	}
	fmt.Println("---------------------------------------------")
	fmt.Printf("%#v \n", cargoBay)
	for _, elem := range cargoBay {
		fmt.Printf("%v", elem[0])
	}
	fmt.Printf("\n")
}

func executeMovements(cargoBay *[][]string , orders Order) {
	q := orders.quantity 
	f := orders.from - 1
	t := orders.to - 1
	bay := *cargoBay

	// var fromStack []string
	// var toStack []string
	var fromStack = make([]string, len(bay[f]))
	var toStack = make([]string, len(bay[t]))
	copy(fromStack, bay[f])
	copy(toStack, bay[t])
	var portable []string = fromStack[:q]
	var oldStack []string = fromStack[q:]
	var newStack []string = append(portable, toStack...)
	bay[f] = oldStack
	bay[t] = newStack
}

func parseCargoInput() [][]string {
	var cargoBay = make([][]string, 9)
	file := readInput("./stack.txt")
	s := bufio.NewScanner(file)
	// Parsing
	for s.Scan() {
		rawLine := s.Bytes()
		// fmt.Printf("%#v \n", line)
		line := make([]string, 0)
		// every 4th index is a character (count from 1)
		for i := 1; i < 34; i = i + 4 {
			line = append(line, string(rawLine[i]))
		}
		for x, elem := range line {
			cargoBay[x] = append(cargoBay[x], elem)
		}

	}
	// cleaning the whitespaces
	for x, stack := range cargoBay {
		for y, elem := range stack {
			if elem != " " {
				stack = stack[y:]
				cargoBay[x] = stack
				break
			}
		}
	}
	return cargoBay
}

func parseMovements(rawLine string) Order {
	reg := regexp.MustCompile(`\d+`)
	match := reg.FindAllString(rawLine, -1)
	q, err := strconv.Atoi(match[0])
	f, err := strconv.Atoi(match[1])
	t, err := strconv.Atoi(match[2])
	if err != nil {
		fmt.Println("Error while trying to get movements by regex.")
	}
	return Order{
		quantity: q,
		from:     f,
		to:       t,
	}

}

func readInput(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("%s \n", "Something went wrong reading the file...")
	}
	return file
}
