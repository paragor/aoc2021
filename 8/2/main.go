package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
)
import "strings"

var a string = `
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
`
var alpha = map[int]string{
	1: "cf",      //2 + len
	7: "acf",     //3 + len
	4: "bcdf",    //4 + len
	5: "abdfg",   //5 + R7
	2: "acdeg",   //5
	3: "acdfg",   //5 + R9
	0: "abcefg",  //6
	6: "abdefg",  //6
	9: "abcdfg",  //6
	8: "abcdefg", //7 + len
}

var beta = map[string]int{}

func init() {
	for k, v := range alpha {
		alpha[k] = SortString(v)
		beta[SortString(v)] = k
	}
}

// a+ b+ c+ d+ e f+ g+

// R1=разница(1, 7) = a
// R2=общее(len(5)) = a,d,g
// R3=общее(R2,4) = d
// R4=разница(1,4) = b,d
// R5=разница(R4,R3) = b
// R6=разница(R1,R2,R3) = g
// R7= len(5) && string.contains(R5) => 5
// R8=разница(R1,R3,R6, R7) = f
// R9 = len(5) && string.contains(R8) && notContains(b) => 3
// R10 = разница(R1,R3,R8,R6, R9) => c
// исключением - e

// далее имея маппинг найти буквы

func Wire(strs []string) map[string]string {
	n1 := findOneByLength(strs, 2)
	n7 := findOneByLength(strs, 3)
	n4 := findOneByLength(strs, 4)
	n8 := findOneByLength(strs, 7)

	len5 := findManyByLength(strs, 5)
	if len(len5) != 3 {
		panic("wtf len5")
	}

	R1 := diff(n1, n7) // a
	assertLen1(R1)
	R2 := union(len5...)
	R3 := union(n4, R2) // d
	assertLen1(R3)
	R4 := diff(n1, n4)
	R5 := diff(R4, R3) // b
	assertLen1(R5)
	R6 := diff(R1, R2, R3) // g
	assertLen1(R6)
	R7 := containsOne(len5, R5)
	R8 := diff(R1+R3+R5+R6, R7) // f
	assertLen1(R8)
	R9 := containsOne(notContains(len5, R5), R8)
	R10 := diff(R1, R3, R8, R6, R9) //c
	assertLen1(R10)
	R11 := diff(n8, R1, R3, R5, R6, R8, R10) //e
	assertLen1(R11)

	res := map[string]string{
		"a": R1, //d
		"b": R5, //e
		"c": R10, //a
		"d": R3, //f
		"e": R11, //g
		"f": R8, //b
		"g": R6, //c
	}
	return res
}

func reverseMap(or map[string]string) map[string]string {
	res := map[string]string{}
	for k, v := range or {
		res[v] = k
	}

	return res
}

func assertLen1(s string) {
	if len(s) != 1 {
		panic("not 1")
	}
}
func notContains(all []string, symb string) []string {
	if len(symb) != 1 {
		panic("wtf symb")
	}
	res := []string{}
	for _, s := range all {
		if !strings.Contains(s, symb) {
			res = append(res, s)
		}
	}

	return res

}
func containsOne(all []string, symb string) string {
	if len(symb) != 1 {
		panic("wtf symb")
	}
	res := ""
	for _, s := range all {
		if strings.Contains(s, symb) {
			if res != "" {
				panic("wtf str symb")
			}
			res = s
		}
	}

	return res

}

func union(all ...string) string {
	res := all[0]
	for _, b := range all[1:] {
		res = union1(res, b)
	}

	return res
}
func union1(a string, b string) string {
	runes := map[rune]int{}
	for _, x := range a {
		if runes[x] == 1 {
			panic("wtf runes u 1")
		}
		runes[x]++
	}
	for _, x := range b {
		if runes[x] == 2 {
			panic("wtf runes u 2")
		}
		runes[x]++
	}
	d := ""
	for r, c := range runes {
		if c == 2 {
			d += string(r)
		}
	}
	return d
}

func diff(all ...string) string {
	res := all[0]
	for _, b := range all[1:] {
		res = diff1(res, b)
	}

	return res
}

func diff1(a, b string) string {
	runes := map[rune]int{}
	for _, x := range a {
		if runes[x] == 1 {
			panic("wtf runes 1")
		}
		runes[x]++
	}
	for _, x := range b {
		if runes[x] == 2 {
			panic("wtf runes 2")
		}
		runes[x]++
	}
	d := ""
	for r, c := range runes {
		if c == 1 {
			d += string(r)
		}
	}
	return d
}
func findManyByLength(strs []string, l int) []string {
	res := []string{}
	for _, str := range strs {
		if len(str) == l {
			res = append(res, str)
		}
	}

	return res
}

func countOne(strs []string, l int) int {

	c := 0
	for _, str := range strs {
		if len(str) == l {
			c++
		}
	}
	return c

}
func findOneByLength(strs []string, l int) string {
	res := ""
	for _, str := range strs {
		if len(str) == l {
			if res != "" {
				panic("wtf")
			}
			res = str
		}
	}
	if res == "" {
		panic("wtf empty str")
	}

	return res
}

//go:embed src.txt
var src string

func main() {
	src = strings.TrimSpace(src)
	strs := strings.Split(src, "\n")
	s := 0
	for _, str := range strs {
		row := strings.Split(str, "|")
		if len(row) != 2 {
			panic("wtf row")
		}
		wired := Wire(strings.Split(strings.TrimSpace(row[0]), " "))
		dnumbers := strings.Split(strings.TrimSpace(row[1]), " ")
		numbers := []int{}
		for _, dnumber := range dnumbers {
			numbers = append(numbers, dnumberToNumber(dnumber, wired))
		}
		x := joinInts(numbers)
		fmt.Println(x)
		s += x
	}
	fmt.Println(s)

	return
}

func joinInts( a []int) int {
	strs := []string{}
	for _,x := range a {
		strs = append(strs, strconv.Itoa(x))
	}
	x, err := strconv.Atoi(strings.Join(strs, ""))
	if err != nil {
		panic("joinints")
	}
	return x
}

func dnumberToNumber(dnumber string, wired map[string]string) int {
	rwired := reverseMap(wired)
	res := ""
	for _, x := range dnumber {
		res += rwired[string(x)]
	}

	x, ok := beta[SortString(res)]
	if !ok {
		fmt.Println(wired, rwired)
		fmt.Println(alpha, beta)
		fmt.Println(SortString(res))
		panic("wtf dnumber")
	}

	return x
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
