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

	preints := []int{}
	for i := 0; i < len(strs); i++ {
		v, err := strconv.Atoi(strs[i])
		if err != nil {
			panic(err)
		}
		preints = append(preints, v)
	}

	ints := []int{}
	for i := 0; i< len(preints) - 2; i++ {
		ints = append(ints, preints[i] + preints[i+1]+preints[i+2])
	}


	fmt.Println(ints)

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
