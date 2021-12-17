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

	x := 0
	y := 0

	for _, str := range strs {
		row := strings.Split(str, " ")

		i, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}

		switch row[0] {
		case "forward":
			x += i
		case "up":
			y -= i
		case "down":
			y += i
		default:
			panic("hui")
		}

	}

	fmt.Println(x * y)
}
