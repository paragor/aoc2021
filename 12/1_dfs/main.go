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
	m.FindPaths(START, END, []string{})
	fmt.Println(strings.Join(UniqStrings(m.paths), "\n"))
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
	m.visited[a] = 0
	m.visited[b] = 0
	if b == START || a == END {
		a, b = b, a
	}

	if a == START {
		if b == END {
			m.graph[a] = append(m.graph[a], b)
			return
		}
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

func (m *Graph) GetVariants(from string) []string {
	if m.graph[from] == nil {
		return []string{}
	}
	return m.graph[from]
}

func (m *Graph) FindPaths(from string, end string, curPath []string) {
	if m.visited[from] == 1 && unicode.IsLower(rune(from[0])) {
		return
	}
	variants := m.GetVariants(from)
	newPath := make([]string, len(curPath))
	copy(newPath, curPath)
	newPath = append(newPath, from)
	m.visited[from]++
	for _, to := range variants {
		if to == end {
			m.paths = append(m.paths, strings.Join(newPath, ",")+","+end)
			continue
		}

		m.FindPaths(to, end, newPath)
	}
	m.visited[from]--
}

func UniqStrings(r []string) []string {
	m := map[string]bool{}
	for _, x := range r {
		m[x] = true
	}
	r = []string{}
	for x := range m {
		r = append(r, x)
	}
	return r
}

func (m *Graph) print() {
	for k, v := range m.graph {
		fmt.Printf("%s => %v\n", k, v)
	}

}
