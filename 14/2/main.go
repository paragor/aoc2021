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
	all := strings.Split(src, "\n\n")
	assert(len(all) == 2, "all")

	pairs := map[string]rune{}
	for _, pairS := range strings.Split(all[1], "\n") {
		pairA := strings.Split(strings.TrimSpace(pairS), "->")
		assert(len(pairA) == 2, "pairA")

		pairs[strings.TrimSpace(pairA[0])] = rune(strings.TrimSpace(pairA[1])[0])
	}

	//fmt.Println(cache)
	template := strings.TrimSpace(all[0])
	last := 0
	for i := 0; i < 20; i++ {
		fmt.Println(countSymInTemplate(template))
		template = calculate(pairs, template)
		//fmt.Println(i+1, template)
		minMaxDiffInTemplate := countMinMaxDiffInTemplate(template)
		fmt.Println(i+1, " minmax:", minMaxDiffInTemplate, "diff:", minMaxDiffInTemplate-last, float64(last)/float64(minMaxDiffInTemplate))
		last = minMaxDiffInTemplate
	}

}

func countMinMaxDiffInTemplate(template string) int {
	count := map[string]int{}
	for _, x := range template {
		count[string(x)]++
	}
	min := 999999999
	max := -1
	for _, v := range count {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min

}
func countSymInPairs(pairs map[string]rune) (map[string]int, map[string]int) {
	before := map[string]int{}
	after := map[string]int{}

}

func countSymInTemplate(template string) map[string]int {
	res := map[string]int{}
	for i, x := range template {
		res[string(x)]++
		if i != 0 && i != len(template)-1 {
			res[string(x)]++
		}
	}

	return res
}

func calculate(pairs map[string]rune, template string) string {
	builder := strings.Builder{}
	builder.Reset()
	builder.Grow(len(template))
	builder.WriteRune(rune(template[0]))
	for j := 1; j < len(template); j++ {
		ab := string(template[j-1]) + string(template[j])

		if res, ok := pairs[ab]; ok {
			builder.WriteRune(res)
		}
		builder.WriteRune(rune(template[j]))
	}

	return builder.String()
}

func assert(u bool, msg string) {
	if !u {
		panic(msg)
	}

}
