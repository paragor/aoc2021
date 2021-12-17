package main

import (
	_ "embed"
	"fmt"
)
import "strings"

//go:embed src.txt
var src string

func main() {
	src = strings.TrimSpace(src)
	strs := strings.Split(src, "\n")

	maxCommon := findMaxCommon(strs, false)
	leastCommon := findMaxCommon(strs, true)
	fmt.Println(maxCommon, bitToNumber(maxCommon), leastCommon, bitToNumber(leastCommon), bitToNumber(maxCommon)*bitToNumber(leastCommon))
}

func findMaxCommon(strs []string, revert bool) string {
	var mostCommon rune
	for i := 0; i < len(strs[0]); i++ {
		s0, s1 := findS0AndS1(strs, i)

		max := s0
		if s0 > s1 {
			max = s0
			mostCommon = '0'
		} else {
			max = s1
			mostCommon = '1'
		}

		if revert {
			if mostCommon == '1' {
				mostCommon = '0'
			} else {
				mostCommon = '1'
			}
		}

		strs = keepNumber(strs, mostCommon, i, max)
		fmt.Println(strs, "", "")
		if len(strs) == 1 {
			break
		}
	}
	return strs[0]
}

func keepNumber(strs []string, bit rune, pos int, count int) []string {
	newstrs := []string{}

	for _, str := range strs {
		if ([]rune(str))[pos] == bit {
			newstrs = append(newstrs, str)
		}
	}
	if len(newstrs) > count {
		return newstrs[:count]
	}

	return newstrs
}

func findS0AndS1(strs []string, i int) (s0, s1 int) {
	for _, str := range strs {
		if str[i] == '0' {
			s0++
		} else {
			s1++
		}
	}

	return s0, s1
}

func bitToNumber(str string) int {
	result := 0
	for i, ch := range str {
		if ch == '1' {
			result += 1 << (len(str) - i - 1)
		}
	}

	return result

}
