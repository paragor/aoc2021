package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)
import "strings"

//go:embed src.txt
var src string

func main() {
	src = strings.TrimSpace(src)
	strs := strings.Split(src, "\n")

	numbers := strings.Split(strs[0], ",")

	fields := []*Filed{}
	boardIsWin := map[*Filed]bool{}
	for _, z := range strings.Split(strings.Join(strs[2:], "\n"), "\n\n") {
		filed := NewFiled(z)
		fields = append(fields, filed)
		boardIsWin[filed] = false
		filed.Print()
		fmt.Println()
	}

	for _, number := range numbers {
		for _, f := range fields {
			f.Mark(number)
			x, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			var winner *Filed = nil
			if countNotWinners(boardIsWin) == 1 {
				winner = GetWinner(boardIsWin)
			}

			if f.GetWinner() {
				boardIsWin[f] = true

				//fmt.Println(f.GetUmarkedSum(), x, f.GetUmarkedSum()*x, " ok")
				//return
			}
			if countNotWinners(boardIsWin) == 0 {
				fmt.Println(winner.GetUmarkedSum()*x, x, winner.GetUmarkedSum())
				return
			}
		}
	}

}
func GetWinner(bb map[*Filed]bool) (first *Filed) {

	for f, b := range bb {
		if !b {
			if first == nil {
				first = f
				return first
			}
		}
	}
	panic("wtf")
	return
}
func countNotWinners(bb map[*Filed]bool) int {
	r := 0
	for _, b := range bb {
		if !b {
			r++
		}
	}
	return r
}

type Filed struct {
	sqrs [][]*cell
}
type cell struct {
	ch   string
	mark bool
}

func NewFiled(f string) *Filed {
	strs := strings.Split(f, "\n")
	sqrs := [][]*cell{}
	for _, s := range strs {
		row := []*cell{}
		for _, z := range strings.Split(s, " ") {
			if !regexp.MustCompile("[0-9]+").MatchString(z) {
				continue
			}
			row = append(row, &cell{ch: z})
		}
		sqrs = append(sqrs, row)
	}

	return &Filed{sqrs: sqrs}
}

func (f *Filed) Mark(v string) {
	for _, row := range f.sqrs {
		for _, e := range row {
			if e.ch == v {
				e.mark = true
			}
		}
	}

}
func (f *Filed) Print() {
	for _, row := range f.sqrs {
		str := ""
		for _, e := range row {
			str += e.ch
			z := "-"
			if e.mark {
				z = "+"
			}
			str += z + "_"
		}
		fmt.Println(str)
	}
	return
}
func (f *Filed) GetUmarkedSum() int {
	sum := 0
	for _, row := range f.sqrs {
		for _, e := range row {
			if !e.mark {
				x, err := strconv.Atoi(e.ch)
				if err != nil {
					panic(err)
				}
				sum += x
			}
		}
	}

	return sum
}

func (f *Filed) GetWinner() bool {
	for _, row := range f.sqrs {
		ok := true
		str := []string{}
		for _, e := range row {
			str = append(str, e.ch)
			if !e.mark {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	for i := 0; i < len(f.sqrs[0]); i++ {
		ok := true
		str := []string{}
		for _, e := range f.sqrs {
			str = append(str, e[i].ch)
			if !e[i].mark {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false

}
