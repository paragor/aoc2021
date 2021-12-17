package main

import (
	_ "embed"
	"fmt"
	"strconv"
)
import "strings"

//go:embed src.txt
var src string

type Cell struct {
	val     int
	flushed bool
}

func (c *Cell) reset() int {
	if !c.flushed && c.val >= 10 {
		panic("wtf reset")
	}
	if c.flushed {
		c.flushed = false
		c.val = 0
		return 1
	}
	return 0

}

func (c *Cell) inc() {
	if c.flushed {
		return
	}
	c.val++
	if c.val == 10 {
		c.flushed = true
	}
}
func main() {
	src = strings.TrimSpace(src)
	rows := strings.Split(src, "\n")

	field := [][]*Cell{}
	for _, row := range rows {
		r := []*Cell{}

		for _, x := range row {
			xi, err := strconv.Atoi(string(x))
			if err != nil {
				panic("wtf atoi")
			}
			r = append(r, &Cell{
				val:     xi,
				flushed: false,
			})
		}

		field = append(field, r)
	}

	flushed := 0
	for i := 0; i < 200; i++ {
		for y, row := range field {
			for x, cell := range row {
				if cell.flushed {
					continue
				}

				cell.inc()
				if cell.flushed {
					flushNearly(field, y, x)
				}
			}
		}
		for _, row := range field {
			for _, cell := range row {
				flushed += cell.reset()
			}
		}
		print(field)
		fmt.Println()
	}
	fmt.Println(flushed)
}
func print(field [][]*Cell) {
	for _, row := range field {
		s := ""
		for _, cell := range row {
			if cell.flushed {
				s += "_"
				continue
			}
			s += strconv.Itoa(cell.val)
		}
		fmt.Println(s)
	}

}

func flushNearly(field [][]*Cell, y, x int) {
	for _, point := range getPontsFor(field, y, x) {
		cell := field[point[0]][point[1]]
		if cell.flushed {
			continue
		}
		cell.inc()
		if cell.flushed {
			flushNearly(field, point[0], point[1])
		}
	}

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
	r := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return r
}

func getPontsFor(field [][]*Cell, y, x int) [][]int {
	res := [][]int{}
	if y+1 < len(field) {
		res = append(res, []int{y + 1, x})
		if x+1 < len(field[0]) {
			res = append(res, []int{y + 1, x + 1})
		}
		if x > 0 {
			res = append(res, []int{y + 1, x - 1})
		}
	}
	if y > 0 {
		res = append(res, []int{y - 1, x})
		if x+1 < len(field[0]) {
			res = append(res, []int{y - 1, x + 1})
		}
		if x > 0 {
			res = append(res, []int{y - 1, x - 1})
		}
	}
	if x+1 < len(field[0]) {
		res = append(res, []int{y, x + 1})
	}
	if x > 0 {
		res = append(res, []int{y, x - 1})
	}
	return res

}
