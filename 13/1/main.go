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
	all := strings.Split(src, "\n\n")
	assert(len(all) == 2, "all")

	field := NewFiled(strings.Split(all[0], "\n"))
	for _, foldFull := range strings.Split(all[1], "\n") {
		fold := strings.Split(strings.ReplaceAll(foldFull, "fold along ", ""), "=")
		val, err := strconv.Atoi(fold[1])
		assert(err == nil, "fold val")
		if fold[0] == "y" {
			field.SplitY(val)
		}
		if fold[0] == "x" {
			field.SplitX(val)
		}
		break;
		//panic("wtf is going on")
	}
	fmt.Println(field.print())
	//folds := strings.Split(all[1], "\n")

}

func assert(u bool, msg string) {
	if !u {
		panic(msg)
	}

}

type Filed struct {
	sqrs [][]*cell
}

func (f *Filed) SplitY(y int) *Filed {

	up := f.sqrs[:y]
	down := f.sqrs[y+1:]
	for i, uRow := range up {
		for j, uCell := range uRow {
			if down[len(down)-i-1][j].mark {
				uCell.mark = true
			}
		}

	}

	f.sqrs = up
	return f

}
func (f *Filed) SplitX(x int) *Filed {

	right := make([][]*cell, 0)
	left := make([][]*cell, 0)
	for _, row := range f.sqrs {
		right = append(right, row[:x])
		left = append(left, row[x+1:])
	}

	for i, rRrow := range right {
		for j, rCell := range rRrow {
			if left[i][len(left[i])-1-j].mark {
				rCell.mark = true
			}
		}

	}

	f.sqrs = right
	return f

}

func (f *Filed) print() int {
	c := 0
	for y := 0; y < len(f.sqrs); y++ {
		s := ""
		for x := 0; x < len(f.sqrs[y]); x++ {
			if f.sqrs[y][x].mark {
				s += "# "
				c++
			} else {
				s += ". "

			}
		}
		fmt.Println(s)
	}
	return c
}

type cell struct {
	mark bool
}

func NewFiled(rawPoints []string) *Filed {
	points := [][]int{}
	maxY := 0
	maxX := 0
	for _, s := range rawPoints {
		split := strings.Split(s, ",")
		assert(len(split) == 2, "split raw points")
		x, err := strconv.Atoi(split[0])
		assert(err == nil, "x "+split[0])
		y, err := strconv.Atoi(split[1])
		assert(err == nil, "y "+split[1])

		points = append(points, []int{y, x})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	sqrs := [][]*cell{}
	for y := 0; y < maxY+1; y++ {
		row := []*cell{}
		for x := 0; x < maxX+1; x++ {
			row = append(row, &cell{mark: false})
		}
		sqrs = append(sqrs, row)
	}
	for _, point := range points {
		sqrs[point[0]][point[1]].mark = true
	}

	return &Filed{sqrs: sqrs}
}
