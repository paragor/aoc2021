package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
)
import "strings"

//go:embed src.txt
var src string

func main() {
	src = strings.TrimSpace(src)
	rows := strings.Split(src, "\n")

	fields := [][]int{}
	for y, row := range rows {
		f := []int{}
		for x := range row {
			v := getValue(rows, y, x)
			res := 0
			if v == 9 {
				res = 1
			}
			f = append(f, res)
		}
		fields = append(fields, f)

	}
	res := []int{}
	for y, row := range fields {
		for x, cell := range row {
			if cell == 0 {
				s := &Stack{}
				s.Push([]int{y, x})
				res = append(res, run(fields, s))
			}
		}
	}
	sort.Ints(res)
	fmt.Println(res[len(res)-1] * res[len(res)-2] * res[len(res)-3])
}

func run(fields [][]int, s *Stack) int {
	res := 0
	for s.len() > 0 {
		coards := s.Pop()
		y, x := coards[0], coards[1]
		if fields[y][x] == 0 {
			fields[y][x] = 1
			for _, c := range getPonts(fields, y, x) {
				s.Push(c)
			}
			res += 1
		}

	}
	return res

}

type Stack struct {
	s [][]int
}

func (s *Stack) len() int {
	return len(s.s)

}
func (s *Stack) Push(res []int) {
	s.s = append(s.s, res)
}
func (s *Stack) Pop() []int {

	if len(s.s) == 0 {
		panic("wtf")
	}
	r := s.s[0]
	s.s = s.s[1:]
	return r
}

func getPonts(rows [][]int, y, x int) [][]int {
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
