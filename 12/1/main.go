package main

import (
	_ "embed"
	"fmt"
	"unicode"
)
import "strings"

//go:embed src.txt
var src string
var debug = true

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
	count, paths := m.FindEnd(START)
	fmt.Println(strings.Join(UniqStrings(paths), "\n"))
	fmt.Println(len(UniqStrings(paths)))
	fmt.Println(count)

}

const (
	START = "start"
	EMPTY = "_"
	END   = "end"
)

type Map struct {
	paths map[string][]string
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
		//for _, bb := range b {
		//	m.paths[a] = append(m.paths[a], string(bb))
		//}
		return
	}
	if b == END {
		//for _, aa := range a {
		//	m.paths[string(aa)] = append(m.paths[string(aa)], b)
		//}
		m.paths[a] = append(m.paths[a], b)
		return
	}
	m.paths[a] = append(m.paths[a], b)
	m.paths[b] = append(m.paths[b], a)
	//for _, aa := range a {
	//	for _, bb := range b {
	//		m.paths[string(aa)] = append(m.paths[string(aa)], string(bb))
	//		m.paths[string(bb)] = append(m.paths[string(bb)], string(aa))
	//	}
	//}
}

func (m *Map) GetVariants(from string) []string {
	if m.paths[from] == nil {
		return []string{}
	}
	return m.paths[from]
}

func (m *Map) FindEnd(from string) (int, []string) {
	count := 0
	paths := []string{}
	variants := m.GetVariants(from)
	for _, to := range variants {
		if to == EMPTY {
			continue
		}
		if to == END {
			count++
			paths = append(paths, END)
			continue
		}

		newM := m.ChooseVariant(from, to)
		newCount, newpaths := newM.FindEnd(to)
		count += newCount

		if debug {
			paths = append(paths, newpaths...)
		}

	}

	if debug {
		for i, path := range paths {
			paths[i] = from + " -> " + path
		}
	}

	return count, paths
}

func (m *Map) ChooseVariant(from string, to string) *Map {
	if to == START || to == END {
		panic("cant choose variant")
	}
	for _, variant := range m.paths[from] {
		if variant == to {
			if unicode.IsUpper(rune(to[0])) {
				return m
			}
			n := m.Copy()

			for _, row := range n.paths {
				for i, vv := range row {
					if vv == to {
						row[i] = EMPTY
					}
				}
			}
			return n
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

	return &Map{paths: paths}
}
