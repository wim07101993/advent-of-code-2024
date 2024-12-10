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

	listA := make([]int, 0)
	listB := make(map[int]int)

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
		listB[b] += b
	}

	var ret int
	for i := 0; i < len(listA); i++ {
		ret += listB[listA[i]]
	}
	return ret
}
