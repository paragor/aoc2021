package main

import (
	_ "embed"
	"fmt"
	"sort"
)
import "strings"

//go:embed src.txt
var src string

var alpha = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
	'>': '<',
}
var score = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
var beta = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

func main() {
	src = strings.TrimSpace(src)
	rows := strings.Split(src, "\n")
	s := []int{}
	for _, row := range rows {
		the := getAutoComplite([]rune(row))
		r := getScore(the)
		fmt.Println(string(the), r)
		if len(the) == 0 {
			continue
		}
		s = append(s, r)
	}
	sort.Ints(s)
	fmt.Println(s)

	fmt.Println(s[len(s)/2])

}

func getScore(the []rune) int {
	sum := 0
	for _, x := range the {
		sum = sum*5 + score[x]
	}
	return sum

}

func getAutoComplite(str []rune) []rune {
	s := &Stack{}
	for _, x := range str {
		if _, ok := beta[x]; ok {
			s.Push(beta[x])
			continue
		}

		should := s.Pop()
		got := x
		if got != should {
			return nil
		}

	}

	res := []rune{}
	for s.len() > 0 {
		x := s.Pop()
		res = append(res, x)
	}
	return res

}

type Stack struct {
	s []rune
}

func (s *Stack) len() int {
	return len(s.s)

}
func (s *Stack) Push(res rune) {
	s.s = append(s.s, res)
}
func (s *Stack) Pop() rune {

	if len(s.s) == 0 {
		panic("wtf")
	}
	r := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return r
}
