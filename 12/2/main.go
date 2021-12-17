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
	m := &Map{}
	for _, row := range rows {
		vars := strings.Split(row, "-")
		if len(vars) != 2 {
			panic("vars")
		}

		m.AddVariant(vars[0], vars[1])
	}

	m.print()
	paths := m.FindPaths(START, END)
	//fmt.Println(strings.Join(UniqStrings(paths), "\n"))
	fmt.Println(len(paths))

}

const (
	START = "start"
	EMPTY = "_"
	END   = "end"
)

type Map struct {
	paths          map[string][]string
	denyDoubleCave bool
}

func (m *Map) AddVariant(a, b string) {
	if m.paths == nil {
		m.paths = map[string][]string{}
	}
	if b == START || a == END {
		a, b = b, a
	}

	if a == START {
		if b == END {
			m.paths[a] = append(m.paths[a], b)
			return
		}
		m.paths[a] = append(m.paths[a], b)
		return
	}
	if b == END {
		m.paths[a] = append(m.paths[a], b)
		return
	}
	m.paths[a] = append(m.paths[a], b)
	m.paths[b] = append(m.paths[b], a)
}

func (m *Map) GetVariants(from string) []string {
	if m.paths[from] == nil {
		return []string{}
	}
	return m.paths[from]
}

func (m *Map) FindPaths(from string, end string) []string {
	count := 0
	paths := []string{}
	variants := m.GetVariants(from)
	for _, to := range variants {
		if to == EMPTY {
			continue
		}
		if to == end {
			count++
			paths = append(paths, end)
			continue
		}

		for _, newM := range m.NewMapsForPair(from, to) {
			newpaths := newM.FindPaths(to, end)

			paths = append(paths, newpaths...)

		}
	}

	for i, path := range paths {
		paths[i] = from + "," + path
	}

	return UniqStrings(paths)
}

func (m *Map) NewMapsForPair(from string, to string) []*Map {
	if to == START || to == END {
		panic("cant choose variant")
	}
	for _, variant := range m.paths[from] {
		if variant == to {
			if unicode.IsUpper(rune(to[0])) {
				return []*Map{m}
			}
			n := m.Copy()
			for _, row := range n.paths {
				for i, vv := range row {
					if vv == to {
						row[i] = EMPTY
						break
					}
				}
			}
			res := []*Map{n}

			if !m.denyDoubleCave {
				nn := m.Copy()
				nn.denyDoubleCave = true
				res = append(res, nn)
			}

			return res
		}
	}
	panic("wtf variant")
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

func (m *Map) print() {
	for k, v := range m.paths {
		fmt.Printf("%s => %v\n", k, v)
	}

}
func (m *Map) Copy() *Map {
	paths := map[string][]string{}
	for k, src := range m.paths {
		dst := make([]string, len(src))
		copy(dst, src)
		paths[k] = dst
	}

	return &Map{paths: paths, denyDoubleCave: m.denyDoubleCave}
}
