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

	gamma := ""
	epsilon := ""

	for i := 0; i < len(strs[0]); i ++ {
		s0 := 0
		s1 := 0
		for _, str := range strs {
			if str[i] == '0' {
				s0 ++
			} else {
				s1++
			}
		}
		if s0 > s1 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	fmt.Println(gamma, epsilon, bitToNumber(gamma), bitToNumber(epsilon), bitToNumber(gamma) * bitToNumber(epsilon))
}

func bitToNumber(str string) int  {
	result := 0
	for i, ch := range str {
		if ch == '1' {
			result += 1 << (len(str) - i - 1)
		}
	}

	return result

}
