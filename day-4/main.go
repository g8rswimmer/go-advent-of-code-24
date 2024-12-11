package main

import (
	"bufio"
	"fmt"
	"os"
)

type matcher func(r int, c int, grid [][]byte, word string) bool

var testGrid = [][]byte{
	[]byte("MMMSXXMASM"),
	[]byte("MSAMXMSMSA"),
	[]byte("AMXSXMAAMM"),
	[]byte("MSAMASMSMX"),
	[]byte("XMASAMXAMM"),
	[]byte("XXAMMXXAMA"),
	[]byte("SMSMSASXSS"),
	[]byte("SAXAMASAAA"),
	[]byte("MAMMMXMMMM"),
	[]byte("MXMXAXMASX"),
}

func main() {

	matchers := []matcher{
		matchUp,
		matchDown,
		matchLeft,
		matchRight,
		matchUpLeft,
		matchUpRight,
		matchDownLeft,
		matchDownRight,
	}

	grid := makeGrid()
	num := 0
	for r, row := range grid {
		for c, b := range row {
			if b != 'X' {
				continue
			}
			for _, m := range matchers {
				if m(r, c, grid, "XMAS") {
					num++
				}
			}
		}
	}
	fmt.Printf("Number of XMAS: %d\n", num)

	num = 0
	for r, row := range grid {
		for c, b := range row {
			if b != 'A' {
				continue
			}
			if matchX(r, c, grid) {
				num++
			}
		}
	}
	fmt.Printf("Number of X-MAS: %d\n", num)

}

func makeGrid() [][]byte {
	f, err := os.Open("search.txt")
	handleError(err)
	defer f.Close()

	grid := [][]byte{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	return grid
}
func handleError(err error) {
	if err == nil {
		return
	}
	panic(err.Error())
}
