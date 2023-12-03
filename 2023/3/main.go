package main

import(
	"bufio"
	"os"
	"log"
	
)
func main(){
	// file, err := os.Open("./example.txt")
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./example2.txt")
	if err != nil{
		log.Fatalf("reading the file failed...")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()

	}
}

func parseIntoGrid(line string) [][]string{
	return 
}
