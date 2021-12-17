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
	rows := strings.Split(src, "\n")

	mins := []int{}
	for y, row := range rows {
		for x := range row {
			v := getValue(rows, y, x)
			min := true
			for _, coards := range getPontsForV(rows, y, x) {
				if v >= getValue(rows, coards[0], coards[1]) {
					min = false
				}
			}
			if min {
				mins = append(mins, v)
			}
		}

	}
	fmt.Println(mins)
	s := 0
	for _, x := range mins {
		s += x + 1
	}
	fmt.Println(s)
}
func getPontsForV(rows []string, y, x int) [][]int {
	res := [][]int{}
	if y+1 < len(rows) {
		res = append(res, []int{y + 1, x})
	}
	if y > 0 {
		res = append(res, []int{y - 1, x})
	}
	if x+1 < len(rows[0]) {
		res = append(res, []int{y, x + 1})
	}
	if x > 0 {
		res = append(res, []int{y, x - 1})
	}
	return res

}

func getValue(rows []string, y, x int) int {
	v, err := strconv.Atoi(string(rows[y][x]))
	if err != nil {
		panic("atoi")
	}

	return v
}
