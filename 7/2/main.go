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

	minX := 99999
	maxX := -1
	for _, s := range strings.Split(src, ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, x)
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
	}

	minTarget := 9999999999
	val := 9999999999999
	for target := minX; target <= maxX; target++ {
		l := 0
		for _, x := range numbers {
			l += sum(max(x, target) - min(x, target))
		}
		if l < val {
			minTarget = target
			val = l
		}
	}

	fmt.Println(val)
	fmt.Println(minTarget)

}

func sum(x int) int {
	return x * (x + 1) / 2

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
