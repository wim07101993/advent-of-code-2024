package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	listA := make([]int, 0)
	listB := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		a, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		listA = append(listA, a)
		listB = append(listB, b)
	}

	sort.Ints(listA)
	sort.Ints(listB)

	var ret int
	for i := 0; i < len(listA); i++ {
		diff := listA[i] - listB[i]
		if diff > 0 {
			ret += diff
		} else {
			ret -= diff
		}
	}
	return ret
}
