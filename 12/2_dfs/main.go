package main

import (
	_ "embed"
	"fmt"
	"unicode"
)
import "strings"

//go:embed src.txt
var src string

func main() {
	src = strings.TrimSpace(src)
	rows := strings.Split(src, "\n")
	m := &Graph{}
	for _, row := range rows {
		vars := strings.Split(row, "-")
		if len(vars) != 2 {
			panic("vars")
		}

		m.AddVariant(vars[0], vars[1])
	}

	m.print()
	m.FindPaths(START, END, []string{}, true)
	fmt.Println(len(m.paths))

}

const (
	START = "start"
	END   = "end"
)

type Graph struct {
	graph          map[string][]string
	paths          []string
	visited        map[string]int
	denyDoubleCave bool
}

func (m *Graph) AddVariant(a, b string) {
	if m.graph == nil {
		m.graph = map[string][]string{}
		m.visited = map[string]int{}
	}
	if b == START || a == END {
		a, b = b, a
	}

	if a == START {
		m.graph[a] = append(m.graph[a], b)
		return
	}
	if b == END {
		m.graph[a] = append(m.graph[a], b)
		return
	}
	m.graph[a] = append(m.graph[a], b)
	m.graph[b] = append(m.graph[b], a)
}

func (m *Graph) FindPaths(from string, end string, path []string, allowTwice bool) {
	if unicode.IsLower(rune(from[0])) {
		if m.visited[from] == 1 {
			if !allowTwice {
				return
			} else {
				allowTwice = false
			}

		}
		if m.visited[from] == 2 {
			return
		}
	}

	path = append(path, from)
	m.visited[from]++
	for _, to := range m.graph[from] {
		if to == end {
			m.paths = append(m.paths, strings.Join(path, ",")+","+end)
			continue
		}

		m.FindPaths(to, end, path, allowTwice)
	}
	m.visited[from]--
}

func (m *Graph) print() {
	for k, v := range m.graph {
		fmt.Printf("%s => %v\n", k, v)
	}

}
