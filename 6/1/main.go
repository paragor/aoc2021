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
	numbers := strings.Split(src, ",")

	initial := []int{}
	for _, n := range numbers {
		x, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		initial = append(initial, x)
	}

	for i := 0; i < 80; i++ {
		for i, x := range initial {
			if x == 0 {
				initial = append(initial, 8)
				initial[i] = 6
				continue
			}

			initial[i]--
		}
	}


	fmt.Println(len(initial))

}
