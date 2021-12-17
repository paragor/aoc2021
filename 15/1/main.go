package main

import (
	_ "embed"
	"fmt"
	"strconv"
)
import "strings"

//go:embed src.txt
var src string

const CacheLen = 3

func main() {
	src = strings.TrimSpace(src)
	strs := strings.Split(src, "\n")
	f := NewField(strs)
	s := &Stack{}
	s.Push(Element{})

	maxY := len(f.f) - 1
	maxX := len(f.f[maxY]) - 1
	for s.Len() > 0 {
		prevCoard := s.Pop()
		f.f[prevCoard.y][prevCoard.x].was = true
		for _, nextCoard := range getPontsFor(f.f, prevCoard.y, prevCoard.x) {

			next := f.f[nextCoard.y][nextCoard.x]
			prev := f.f[prevCoard.y][prevCoard.x]

			if !next.was {
				s.Push(Element{x: nextCoard.x, y: nextCoard.y})
			}

			w := next.w + prev.wFromStart

			if w < next.wFromStart || !next.was {
				next.wFromStart = w
			}
			next.was = true

		}
		//f.print()
		//fmt.Println()

		if s.Len() > 2000 {
			panic("wtf len")
		}
	}
	last := f.f[maxY][maxX]
	fmt.Println(last.wFromStart - last.w)

}

type Point struct {
	w          int
	wFromStart int
	was        bool
}

type Field struct {
	f [][]*Point
}

//func (f*Field) NewField() *Field {
func (f *Field) print() {
	for _, row := range f.f {
		s := ""
		for _, cell := range row {
			s += fmt.Sprintf("%d[%d] ", cell.w, cell.wFromStart)
		}
		fmt.Println(s)
	}

}
func NewField(strs []string) *Field {
	f := [][]*Point{}
	for _, rstr := range strs {
		row := []*Point{}
		for _, c := range rstr {
			x, err := strconv.Atoi(string(c))
			assert(err == nil, "atoi")
			row = append(row, &Point{
				w:          x,
				wFromStart: 0,
			})

		}

		f = append(f, row)
	}

	x := &Field{f: f}
	x.f[0][0].wFromStart = x.f[0][0].w
	return x
}

func assert(u bool, msg string) {
	if !u {
		panic(msg)
	}

}

type Element struct {
	x, y int
}

type Stack struct {
	s []Element
}

func (s *Stack) Len() int {
	return len(s.s)

}
func (s *Stack) Push(res Element) {
	s.s = append(s.s, res)
}
func (s *Stack) Pop() Element {

	if len(s.s) == 0 {
		panic("wtf")
	}
	r := s.s[0]
	s.s = s.s[1:]
	//r := s.s[len(s.s)-1]
	//s.s = s.s[:len(s.s)-1]
	return r
}
func getPontsFor(field [][]*Point, y, x int) []Element {
	res := []Element{}

	for _, point := range [][]int{
		/*     */ {-1, 0},
		{0, -1} /*     */, {0, 1},
		/*     */ {1, 0},
	} {
		yy, xx := point[0], point[1]
		if y+yy >= len(field) || y+yy < 0 || x+xx >= len(field[y]) || x+xx < 0 {
			continue
		}
		res = append(res, Element{y: yy + y, x: xx + x})
	}

	return res
}
