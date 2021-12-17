package main

import (
	_ "embed"
	"fmt"
	"strconv"
)
import "strings"

//go:embed src.txt
var src string

func main() {
	src = strings.TrimSpace(src)
	numbers := []int{}
	for _, s := range strings.Split(src, ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, x)
	}

	lengths := []int{}
	minIdx := 0
	for i, target := range numbers {
		l := 0
		for _, x := range numbers {
			l += max(x, target) - min(x, target)
		}
		lengths = append(lengths, l)
		if l < lengths[minIdx] {
			minIdx = i
		}
	}

	fmt.Println(lengths)
	fmt.Println(minIdx, lengths[minIdx])

}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
