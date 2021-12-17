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
	for i := 0; i < 20000; i++ {
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
		ok := true
		for _, row := range field {
			for _, cell := range row {
				ok = ok && cell.flushed
				flushed += cell.reset()
			}
		}
		print(field)
		fmt.Println(flushed)
		fmt.Println()
		if ok {
			panic(i + 1)
		}
	}
}
func print(field [][]*Cell) {
	for _, row := range field {
		s := ""
		for _, cell := range row {
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

func getPontsFor(field [][]*Cell, y, x int) [][]int {
	res := [][]int{}

	for _, point := range [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*     */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	} {
		yy, xx := point[0], point[1]
		if y+yy >= len(field) || y+yy < 0 || x+xx >= len(field[y]) || x+xx < 0 {
			continue
		}
		res = append(res, []int{yy + y, xx + x})
	}

	return res
}
