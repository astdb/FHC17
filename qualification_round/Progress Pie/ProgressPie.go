
// facebook hacker cup qualification round 2017 - progress pie
// ------------------------------------------------------------
// given a a percentage P and a point on a 100x100 plane (in the middle of which there's a progress pie whose P%*360 
// sector from the vertical is darkened to indicate a P% progress) determine if the given point is shaded black 
// or white. all points within 10^-6 units of a given (X,Y) is considered to be of the color of that point.

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	// ensure input filename entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Enter input filename on command line as an argument to ProgressPie")
		return
	}

	// capture input file name
	inFileName := os.Args[1]
	file, err := os.Open(inFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}

	defer file.Close()

	// setup scanner to read commands input file
	scanner := bufio.NewScanner(file)
	var numPoints int

	// iterate through the commands
	i := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		 // fmt.Println(line)

		 if i == 0 {
			 // read number of points
			 numPoints, err = strconv.Atoi(line)

			 if err != nil || numPoints < 1 || numPoints > 1000 {
				fmt.Fprintf(os.Stderr, "Invalid number of points: %d\n", numPoints)
				return
			 }

			 fmt.Printf("Line #%d: %d\n", i+1, numPoints)

			 // go on to reading tescases from next line
			 i++
			 continue
		 }

		 // only required to read the given number of test cases
		 if i > numPoints {
			 break
		 }

		 // read testcase: percentage, X and Y
		 line_split := strings.Split(line, " ")
		 if len(line_split) < 3 {
			 fmt.Fprintf(os.Stderr, "Invalid test case: %s\n", line)
			 // i++
			 // continue
			 return
		 }

		 percentage, err1 := strconv.ParseFloat(strings.TrimSpace(line_split[0]), 32)
		 x, err2 := strconv.ParseFloat(strings.TrimSpace(line_split[1]), 32)
		 y, err3 := strconv.ParseFloat(strings.TrimSpace(line_split[2]), 32)

		 if err1 != nil || err2 != nil || err3 != nil {
			fmt.Fprintf(os.Stderr, "Invalid component(s) in test case: %s\n", line)
			// i++
			 // continue
			 return
		 }

		 fmt.Printf("Line #%d: %.2f %.2f %.2f\n", i+1, percentage, x, y)

		 i++
	}
}