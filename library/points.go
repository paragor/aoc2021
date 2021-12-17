package library

func getPontsFor(field [][]interface{}, y, x int) [][]int {
	res := [][]int{}

	for _, point := range [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*     */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	} {
		yy, xx := point[0], point[1]
		if y+yy >= len(field) || y+yy < 0 || x+xx >= len(field[y]) || x+xx < 0 {
			continue
		}
		res = append(res, []int{yy + y, xx + x})
	}

	return res
}
