package main

import "fmt"

type Race struct {
	time     int
	distance int
}

var races []Race
var testRace []Race

// Part 1
// func calculateDifferentWays(r Race) []int{
// 	var waysToWin []int
// 	for speed := 0; speed <= r.time; speed++ {
// 		timeToMove := r.time - speed
// 		distance := timeToMove * speed
// 		// fmt.Printf("Holding down %d seconds, you will travel with a speed of %d for %d miliseconds, to distance of %d. ", speed, speed, timeToMove, distance)
// 		if(distance > r.distance){
// 			// fmt.Printf("This way you would win! \n")
// 			waysToWin = append(waysToWin, speed)
// 		}else{
// 			fmt.Printf("\n")
// 		}
// 	}
// 	// fmt.Printf("All ways to win: %#v \n", waysToWin)
// 	return waysToWin

// }

// Part 2

// Help from: https://www.youtube.com/watch?v=JkfOakgO9YI
func calculateDifferentWays(r Race) int {

	n := 0
	// var holds []int

	for hold := 0; hold < r.time; hold++ {
		timeToMove := r.time - hold
		distance := timeToMove * hold
		if distance > r.distance {
			// holds = append(holds, hold)
			n += 1
		}else{

		}
	}
	// fmt.Printf("Winning hold tims: %#v \n",holds)
	return n
}

func main() {

	longRace := Race{
		time:     47986698,
		distance: 400121310111540,
	}
	// longRace := Race{
	// 	time: 71530,
	// 	distance: 940200,
	// }

	// Part 2
	fmt.Printf("Part2: %d \n", calculateDifferentWays(longRace))

	// Tests
	// testRace = append(testRace, Race{
	// 	time:     7,
	// 	distance: 9,
	// })

	// testRace = append(testRace, Race{
	// 	time:     15,
	// 	distance: 40,
	// })
	// testRace = append(testRace, Race{
	// 	time:     30,
	// 	distance: 200,
	// })
	// // Puzzle
	// races = append(races, Race{
	// 	time:     47,
	// 	distance: 400,
	// })
	// races = append(races, Race{
	// 	time:     98,
	// 	distance: 1213,
	// })
	// races = append(races, Race{
	// 	time:     66,
	// 	distance: 1011,
	// })
	// races = append(races, Race{
	// 	time:     98,
	// 	distance: 1540,
	// })

	// calculateDifferentWays(races[0])
	// calculateDifferentWays(testRace[0])
	// calculateDifferentWays(testRace[1])
	// calculateDifferentWays(testRace[2])

	// Part 1
	// var everyLen []int
	// for _, elem := range races{
	// 	everyLen = append(everyLen, len(calculateDifferentWays(elem)))

	// }
	// fmt.Printf("everyLen: %v \n", everyLen)

	// multiply := 1
	// for _, elem := range everyLen{
	// 	multiply *= elem
	// }
	// fmt.Printf("Multiply: %v \n", multiply)

}
