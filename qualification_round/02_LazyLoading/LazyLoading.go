// facebook hacker cup qualification round 2017 - lazy loading
// -----------------------------------------------------------
// given a set of weights to be divided into sets of atleast 50lb, find the maximum
// number of sets they can be divided into. the total weight of a set is determined
// by assuming all items of that set each weigh at least as much as the first item.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ensure input filename entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Usage: > go run LazyLoading.go <input.file>")
		return
	}

	// capture input file name
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}
	defer file.Close()

	// setup scanner to read input file
	scanner := bufio.NewScanner(file)

	i := 0
	days := 0
	day := 0

	readingWeights := false
	weightTotal := 0
	weight := 0

	for scanner.Scan() {
		// line := strings.TrimSpace(scanner.Text())
		// num, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		// fmt.Printf("%d\n", days)

		if i == 0 {
			// read in number of days to follow
			days, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Printf("%d DAYS\n------\n", days)
		} else {
			if !readingWeights {
				// reading in a weight total
				day++
				weightTotal, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))
				readingWeights = true
				fmt.Printf("Total Weights for day #%d: %d\n", day, weightTotal)
			} else {
				// reading in a weight
				weight, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))
				fmt.Printf("\t%d\n", weight)
				weightTotal--
				if weightTotal <= 0 {
					readingWeights = false
				}
			}
		}

		i++
	}

	return

	// days := 0
	// if scanner.Scan() {
	// 	days, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	// 	fmt.Printf("%d\n----\n", days)
	// }

	// for i := 0; i < days; i++ {
	// 	numWeights := 0
	// 	if scanner.Scan() {
	// 		numWeights, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	// 		fmt.Printf("%d\n", numWeights)
	// 	}

	// 	weight := 0
	// 	for j := 0; j < numWeights; j++ {
	// 		if scanner.Scan() {
	// 			weight, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	// 			fmt.Printf("\t%d\n", weight)
	// 		}
	// 	}
	// }

	// // iterate through the commands
	// i := 0
	// for scanner.Scan() {
	// 	line := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// 	if i == 0 {
	// 		// number of days

	// 	}

	// 	i++
	// }
}
