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
	numbers := strings.Split(src, ",")

	sum := 0
	days := 256
	for _, n := range numbers {
		x, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		sum += count(days - x - 1)
	}

	fmt.Println(sum)
	//fmt.Println( daysToCount)
}

var daysToCount = map[int]int{}

func count(days int) int {
	if days < 0 {
		return 1
	}
	if res, isExists := daysToCount[days]; isExists {
		return res
	}

	c := count(days-7) + count(days-9)
	daysToCount[days] = c

	return c
}
