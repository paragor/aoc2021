package main

import (
	_ "embed"
	"fmt"
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
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
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
	s := 0
	for i, row := range rows {
		fmt.Printf("%1d ",i)
		the := getScore([]rune(row))
		fmt.Printf("\n")
		s += the
	}
	fmt.Println(s)

}

func getScore(str []rune) int {
	s := &Stack{}
	for i, x := range str {
		if _, ok := beta[x]; ok {
			s.Push(beta[x])
			continue
		}

		should := s.Pop()
		got := x
		if got != should {
			fmt.Printf("Expected: %s, but found: %s, score: %d, pos: %d\n", string(should), string(x), score[x], i)
			return score[x]
		}

	}
	//if s.len() != 0 {
	//	panic("wtf")
	//}

	return 0
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
	r := s.s[len(s.s) -1]
	s.s = s.s[:len(s.s) - 1]
	return r
}
