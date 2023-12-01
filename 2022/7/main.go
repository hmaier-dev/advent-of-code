package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	// "regexp"
)

type Directory struct {
	name            string
	parent          *Directory
	children        []*Directory
	containingFiles []File
}

type File struct {
	name   string
	size   int
	parent *Directory
}

var currDir *Directory
var lsing bool
var root *Directory
var totalSum int

var totalDiskSpace int = 70000000
var neededSpace int = 30000000
var usedSpace int
var dirToDelete int = 0

var allTotalDirSizes []int

func main() {

	file, _ := os.Open("./input.txt")
	// file, _ := os.Open("./example.txt")
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		parse(line)
	}

	// fmt.Printf("%#v \n", currDir)

	// printTree(root, "")
	// _ = part1_calc(root)
	// fmt.Println("Part 1")
	// fmt.Printf("Total sum: %#v \n", totalSum)

	_ = part2_calc(root)
	fmt.Println("Part 2")
	fmt.Printf("Used Space: %#v \n", usedSpace)
	freeSpace := totalDiskSpace - usedSpace
	fmt.Printf("Free Space: %#v \n", freeSpace)
	t := usedSpace + freeSpace
	fmt.Printf("Total: %#v \n", t)
	sort.Ints(allTotalDirSizes)
	for _, elem := range allTotalDirSizes {
		if (freeSpace + elem) >= neededSpace {
			fmt.Printf("Smallest possible dir size: %d \n", elem)
			break
		}
	}
	fmt.Printf("All dir Sizes: %#v \n", allTotalDirSizes)
}

func parse(line string) {
	parts := strings.Split(line, " ")

	switch {
	// create root
	case len(parts) >= 3 && parts[2] == "/":
		root = &Directory{
			name:            parts[2],
			parent:          nil,
			children:        nil,
			containingFiles: nil,
		}
		currDir = root
		break
	// DOWN the tree
	case len(parts) >= 3 && parts[2] != "..":
		// adding as child an descending into it
		newDir := &Directory{
			name:     string(parts[2]),
			parent:   currDir,
			children: nil,
		}
		currDir.children = append(currDir.children, newDir)
		currDir = newDir
		break
	// UP the tree
	case len(parts) >= 3 && parts[2] == "..":
		if currDir.parent != nil {
			currDir = currDir.parent
		}
		break
	// just ls
	case parts[1] == "ls":
		break
	// its a file
	case parts[0] != "dir":
		name := parts[1]
		size, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("str to int conversion went wrong.")
			return
		}
		newFile := File{
			name:   name,
			size:   size,
			parent: currDir}
		currDir.containingFiles = append(currDir.containingFiles, newFile)
		break
	// its a dir
	case parts[0] == "dir":
		// add the dir when cding into it
		break
	default:
		panic("Unkown command.")
	}

}

func printTree(node *Directory, indent string) {
	fmt.Println(indent + node.name)
	for _, file := range node.containingFiles {
		fmt.Println(indent+"  ", file.name, file.size)
	}
	for _, child := range node.children {
		printTree(child, indent+"  ")
	}
}

func part1_calc(dir *Directory) int {
	var limit int = 100000
	var dirSum int
	// adding all files
	for _, file := range dir.containingFiles {
		dirSum += file.size
	}
	// traverse into sub directorys
	for _, child := range dir.children {
		dirSum += part1_calc(child)
	}
	// add to total if greater equal than limit
	if dirSum <= limit {
		totalSum += dirSum
	}
	return dirSum
}


func part2_calc(dir *Directory) int {
	var dirSize int
	// var totalDirSize int
	// adding all files
	for _, file := range dir.containingFiles {
		dirSize += file.size
	}
	// only size of current files
	usedSpace += dirSize
	// traverse into sub directorys
	for _, child := range dir.children {
		dirSize += part2_calc(child)
	}
	allTotalDirSizes = append(allTotalDirSizes, dirSize)
	return dirSize
}
