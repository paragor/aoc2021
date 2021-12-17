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
	ints := []int{}

	for i := 0; i < len(strs); i++ {
		v, err := strconv.Atoi(strs[i])
		if err != nil {
			panic(err)
		}
		ints = append(ints, v)
	}
	ups := 0
	for i := 1; i < len(ints); i++ {
		prev := ints[i-1]
		cur := ints[i]
		if prev < cur {
			ups++
		}
	}
	fmt.Println(ups)
}
