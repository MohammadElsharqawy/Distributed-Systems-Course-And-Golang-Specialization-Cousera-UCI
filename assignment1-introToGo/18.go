package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pict := make([][]uint8, dx)

	for i := range pict {
		pict[i] = make([]uint8, dy)

		for j := range pict[i] {
			pict[i][j] = uint8(i ^ j)
		}

	}

	return pict
}

func main() {
	pic.Show(Pic)
}
