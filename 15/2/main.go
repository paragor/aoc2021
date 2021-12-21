package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strconv"
)
import "strings"

//go:embed src.txt
var src string

const CacheLen = 3

func main() {
	fmt.Println(0%10, 9%10, 10%10)
	src = strings.TrimSpace(src)
	strs := strings.Split(src, "\n")
	f := NewField(strs)
	f = f.superCopy()
	first := f.f[0][0]
	first.wFromStart = first.w
	maxY := len(f.f) - 1
	maxX := len(f.f[maxY]) - 1
	f.search()
	last := f.f[maxY][maxX]
	fmt.Println(last.wFromStart - first.w)
	//f.print()
}

func (f *Field) search() {

	s := &Stack{}
	heap.Init(s)
	heap.Push(s, WeightPoint{})
	for s.Len() > 0 {
		prevCoard := heap.Pop(s).(WeightPoint)
		if prevCoard.x == len(f.f[0]) && prevCoard.y == len(f.f) {
			return
		}
		f.f[prevCoard.y][prevCoard.x].was = true
		for _, nextCoard := range getPontsFor(f.f, prevCoard.y, prevCoard.x) {

			next := f.f[nextCoard.y][nextCoard.x]
			prev := f.f[prevCoard.y][prevCoard.x]

			w := next.w + prev.wFromStart
			if !next.was {
				heap.Push(s, WeightPoint{x: nextCoard.x, y: nextCoard.y, W: w})
			}

			if w < next.wFromStart || !next.was {
				next.wFromStart = w
			}
			next.was = true

		}

	}
}

type Point struct {
	w          int
	wFromStart int
	was        bool
	BEST       bool
}

type Field struct {
	f [][]*Point
}

func (f *Field) markbest(startP WeightPoint) {
	start := f.f[startP.y][startP.x]
	start.BEST = true
	if startP.y == 0 && startP.x == 0 {
		return
	}
	for _, nextP := range getPontsFor(f.f, startP.y, startP.x) {
		next := f.f[nextP.y][nextP.x]
		//if nextP.y == len(f.f)-2 && nextP.x == len(f.f[0])-2 {
		//	continue
		//}
		if next.BEST {
			continue
		}
		if next.wFromStart == start.wFromStart-start.w {
			f.markbest(nextP)
		}
	}
	//panic("wtf")
}

//func (f*Field) NewField() *Field {
func (f *Field) print() {
	f.markbest(WeightPoint{
		x: len(f.f[0]) - 1,
		y: len(f.f) - 1,
	})
	for i, row := range f.f {
		s := ""
		for j, cell := range row {
			//s += fmt.Sprintf("%d[%d] ", cell.w, cell.wFromStart)
			if j%(len(row)/5) == 0 {
				//s += fmt.Sprintf(" ")
			}
			if cell.BEST {

				s += fmt.Sprintf("@")

			} else {
				s += fmt.Sprintf("%d", cell.w)
			}
		}
		if i%(len(row)/5) == 0 {
			//fmt.Println()
		}
		fmt.Println(s)
	}

}
func (f *Field) superCopy() *Field {
	nf := [][]*Point{}
	for i := 0; i < 5*len(f.f); i++ {
		nf = append(nf, make([]*Point, 5*len(f.f[0])))
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			startI := i * len(f.f)
			startJ := j * len(f.f[0])

			for ii, row := range f.f {
				for jj, cell := range row {
					if i == 0 && j == 0 {
						nf[startI+ii][startJ+jj] = &Point{
							w:          cell.w,
							wFromStart: cell.wFromStart,
							was:        cell.was,
						}
					} else if j != 0 {
						prev := nf[startI+ii][startJ+jj-len(f.f[0])]
						r := (prev.w + 1) % 10
						if r == 0 {
							r = 1
						}
						nf[startI+ii][startJ+jj] = &Point{
							w:          r,
							wFromStart: prev.wFromStart,
							was:        prev.was,
						}
					} else {
						prev := nf[startI+ii-len(f.f)][startJ+jj]
						r := (prev.w + 1) % 10
						if r == 0 {
							r = 1
						}
						nf[startI+ii][startJ+jj] = &Point{
							w:          r,
							wFromStart: prev.wFromStart,
							was:        prev.was,
						}
					}
				}
			}
		}
	}

	return &Field{f: nf}
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
	return x
}

func assert(u bool, msg string) {
	if !u {
		panic(msg)
	}

}

type WeightPoint struct {
	x, y int
	W    int
}

type Stack struct {
	s []WeightPoint
}

func (s *Stack) Len() int {
	return len(s.s)

}
func (s *Stack) Less(i, j int) bool {
	return s.s[i].W < s.s[j].W
}
func (s *Stack) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s *Stack) Push(res interface{}) {
	s.s = append(s.s, res.(WeightPoint))
}
func (s *Stack) Pop() interface{} {
	if len(s.s) == 0 {
		panic("wtf")
	}
	r := s.s[len(s.s) - 1]
	s.s = s.s[:len(s.s) - 1]

	return r
}

func getPontsFor(field [][]*Point, y, x int) []WeightPoint {
	res := []WeightPoint{}

	for _, point := range [][]int{
		/*     */ {-1, 0},
		{0, -1} /*     */, {0, 1},
		/*     */ {1, 0},
	} {
		yy, xx := point[0], point[1]
		if y+yy >= len(field) || y+yy < 0 || x+xx >= len(field[y]) || x+xx < 0 {
			continue
		}
		res = append(res, WeightPoint{y: yy + y, x: xx + x})
	}

	return res
}
