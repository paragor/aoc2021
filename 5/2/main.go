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
	strs := strings.Split(src, "\n")

	rows := []*MarkerPlan{}
	xMax, yMax := 0, 0
	for _, s := range strs {
		r := NewRow(s)
		rows = append(rows, r)
		xMax = max(max(r.fX, r.tX), xMax)
		yMax = max(max(r.fY, r.tY), yMax)
	}
	f := NewField(xMax, yMax)
	for _, r := range rows {
		f.Mark(r)
	}

	//f.Print()
	fmt.Println(f.Count())

}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type MarkerPlan struct {
	fX int
	fY int

	tX int
	tY int
}

func (r *MarkerPlan) getDiagonalPoints() []Point {
	xDelta := 1
	if r.fX > r.tX {
		xDelta = -1
	}

	yDelta := 1
	if r.fY > r.tY {
		yDelta = -1
	}

	res := []Point{}
	x, y := r.fX, r.fY
	for {
		if x*xDelta > r.tX*xDelta || y*yDelta > r.tY*yDelta {
			break
		}
		res = append(res, Point{x, y})
		x += xDelta
		y += yDelta
	}

	return res
}

type Point struct {
	x, y int
}

func (r *MarkerPlan) isHor() bool {
	return r.tX != r.fX
}
func (r *MarkerPlan) isVer() bool {
	return r.tY != r.fY
}
func (r *MarkerPlan) isDiag() bool {
	return r.isVer() == r.isHor()
}

func NewRow(str string) *MarkerPlan {
	fr := strings.Split(str, "->")
	if len(fr) != 2 {
		panic("count ->")
	}
	r := &MarkerPlan{}

	r.fX, r.fY = findNumbers(fr[0])
	r.tX, r.tY = findNumbers(fr[1])

	return r
}

func findNumbers(str string) (int, int) {
	ss := strings.Split(str, ",")
	if len(ss) != 2 {
		panic("count ->")
	}
	f, err := strconv.Atoi(strings.TrimSpace(ss[0]))
	if err != nil {
		panic(err)
	}
	s, err := strconv.Atoi(strings.TrimSpace(ss[1]))
	if err != nil {
		panic(err)
	}
	return f, s
}

type Field struct {
	f [][]int
}

func (f *Field) Print() {
	for _, r := range f.f {
		fmt.Println(r)
	}

}

func (f *Field) Mark(row *MarkerPlan) {
	if row.isDiag() {
		for _, p := range row.getDiagonalPoints() {
			f.f[p.y][p.x]++
		}
		return
	}

	for i := min(row.fX, row.tX); i <= max(row.fX, row.tX); i++ {
		for j := min(row.fY, row.tY); j <= max(row.fY, row.tY); j++ {
			f.f[j][i]++
		}
	}
}
func (f *Field) Count() int {
	count := 0
	for _, r := range f.f {
		for _, cell := range r {
			if cell >= 2 {
				count++
			}
		}

	}
	return count
}

func NewField(xmax, ymax int) *Field {
	f := [][]int{}
	for i := 0; i < ymax+1; i++ {
		f = append(f, make([]int, xmax+1))
	}
	return &Field{f: f}
}
