package main

import (
	_ "embed"
	"fmt"
	"sync"
)
import "strings"

//go:embed src.txt
var src string

const CacheLen = 3

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
	for i := 0; i < 40; i++ {
		template = concurrentCalculate(pairs, template, 12)
		//fmt.Println(i+1, template)
		fmt.Println(i + 1)
	}

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

	fmt.Println(min, max, max-min)
}

func generatesCache(pairs map[string]rune, str string, count int) map[string]string {
	alpha := getAlpabet(str)
	input := []string{}
	for _, x := range alpha {
		input = append(input, string(x))
	}
	gen := generate(alpha, count, input)
	result := make(map[string]string, len(gen))
	for _, str := range gen {
		result[str] = calculate(pairs, str)
	}
	return result
}
func generate(alpha string, count int, input []string) []string {

	results := make([]string, 0, len(input)*len(alpha))
	for _, str := range input {
		for _, w := range alpha {
			results = append(results, str+string(w))
		}
	}

	if len(results[0]) == count {
		return results
	}
	return generate(alpha, count, results)
}

func getAlpabet(str string) string {
	for _, pattern := range []string{" ", "\n", "-", ">"} {
		str = strings.ReplaceAll(str, pattern, "")
	}
	return Uniq(str)

}
func Uniq(r string) string {
	m := map[string]bool{}
	for _, x := range r {
		m[string(x)] = true
	}
	r = ""
	for x := range m {
		r += x
	}
	return r
}

func concurrentCalculate(pairs map[string]rune, template string, count int) string {
	chunks := Chunks(template, len(template)/count+1)
	if len(chunks) == 1 {
		return calculate(pairs, template)
	}

	result := make([]string, len(chunks))

	wg := sync.WaitGroup{}
	for i := 0; i < len(chunks); i++ {
		wg.Add(1)
		go func(i int) {
			result[i] = calculate(pairs, chunks[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	template = result[0]
	for i := 1; i < len(result); i++ {
		template = merge(pairs, template, result[i])
	}

	return template

}
func calculateWithCache(cache map[string]string, pairs map[string]rune, template string) string {
	if len(template) < CacheLen*2 {
		return calculate(pairs, template)
	}

	last := calculate(pairs, template[len(template)-CacheLen-1:])
	first := cache[template[:CacheLen]]

	result := first
	for i := CacheLen * 2; i < len(template)-(CacheLen); i += CacheLen {
		result = merge(pairs, result, cache[template[i-CacheLen:i]])
	}

	return merge(pairs, result, last)

}

func merge(pairs map[string]rune, a, b string) string {
	middle := calculate(pairs, string(a[len(a)-1])+string(b[0]))

	return a[:len(a)-1] + middle + b[1:]
}

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
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
