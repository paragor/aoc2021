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

	rows := []*Row{}
	for _, s := range strs {
		r := NewRow(s)
		rows = append(rows, r)
	}
	xMax, yMax := 0, 0
	for _, r := range rows {
		xMax = max(max(r.fX, r.tX), xMax)
		yMax = max(max(r.fY, r.tY), yMax)
	}
	f := NewField(xMax, yMax)
	for _, r := range rows {
		f.Mark(r)
	}

	f.Print()
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

type Row struct {
	fX int
	fY int

	tX int
	tY int
}

func (r *Row) isHor() bool {
	return r.tX != r.fX
}
func (r *Row) isVer() bool {
	return r.tY != r.fY
}

func NewRow(str string) *Row {
	fr := strings.Split(str, "->")
	if len(fr) != 2 {
		panic("count ->")
	}
	r := &Row{}

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

func (f *Field) Mark(row *Row) {
	if row.isHor() != row.isVer() {
		for i := min(row.fX, row.tX); i <= max(row.fX, row.tX); i++ {
			for j := min(row.fY, row.tY); j <= max(row.fY, row.tY); j++ {
				f.f[j][i]++
			}
		}
		return
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
