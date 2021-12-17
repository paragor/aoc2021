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
	strs := strings.Split(src, "\n\n")

	fmt.Println(strs[0], extractTile(strs[0]))
	tiles := []*Tile{}
	for _, str := range strs {
		tiles = append(tiles, extractTile(str))
	}

	eqMap := map[string][]*Tile{}

	for _, tile := range tiles {
		for _, row := range tile.lines {
			eqMap[row] = append(eqMap[row], tile)
		}
	}

	tilesToNonBorderds := map[*Tile]int{}
	for _, tile := range tiles {
		for _, row := range tile.lines {
			if len(eqMap[row]) == 1 {
				tilesToNonBorderds[tile]++
			}
		}
	}

	counter := 1
	for _, tile := range tiles {
		if tilesToNonBorderds[tile] == 2 {
			counter *= tile.name
		}
	}
	fmt.Println(counter)
}

func extractTile(source string) *Tile {
	source = strings.TrimSpace(source)

	rows := strings.Split(source, "\n")

	name := strings.ReplaceAll(rows[0], ":", "")
	name = strings.ReplaceAll(name, "Tile ", "")
	nameI, err := strconv.Atoi(name)
	if err != nil {
		panic(err)
	}

	tile := Tile{
		lines: []string{prepare(rows[1]), prepare(rows[len(rows)-1])},
		name:  nameI,
	}
	rows = rows[1:]

	left, right := "", ""
	for _, row := range rows {
		left = left + string(row[0])
		right = right + string(row[len(row)-1])
	}
	tile.lines = append(tile.lines, prepare(left), prepare(right))

	return &tile
}

func prepare(row string) string {
	return min(row, Reverse(row))
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type Tile struct {
	lines []string
	name  int
}
