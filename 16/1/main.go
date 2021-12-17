package main

import (
	_ "embed"
	"fmt"
	"strconv"
)
import "strings"

//go:embed src.txt
var src string

const CacheLen = 3

func main() {
	src = strings.TrimSpace(src)
	strs := strings.Split(src, "\n")



}
