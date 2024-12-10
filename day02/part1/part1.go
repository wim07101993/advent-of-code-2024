package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		var previousDiff int
		isFirst := true
		isSafe := true

		previous, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		for i := 1; i < len(split); i++ {
			v, err := strconv.Atoi(split[i])
			if err != nil {
				panic(err)
			}

			if v == previous {
				// diff has to be at least 1
				isSafe = false
				break
			}
			newDiff := v - previous
			if newDiff > 3 || newDiff < -3 {
				// all diffs must be less than 3
				isSafe = false
				break
			}

			if isFirst {
				isFirst = false
				previousDiff = v - previous
			} else {
				if newDiff < 0 && previousDiff > 0 || newDiff > 0 && previousDiff < 0 {
					// all diffs must either increase or decrease, no mixing
					isSafe = false
					break
				}
			}

			previous = v
		}

		if isSafe {
			safeCount += 1
		}

		isSafe = true
	}

	return safeCount
}
