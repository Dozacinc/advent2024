package m04

import (
	"fmt"

	"github.com/Dozacinc/advent2024/util"
)

func Run() {
	content, err := util.GetFileContentLines("./data/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	letterGrid := getLetterGrid(content)
	forward := getForward(letterGrid)
	backward := getBackward(letterGrid)
	downward := getDownward(letterGrid)
	upward := getUpward(letterGrid)
	diagonalDown := getDiagonalDown(letterGrid)
	diagonalUp := getDiagonalUp(letterGrid)
	total := forward + backward + downward + upward + diagonalDown + diagonalUp
	fmt.Printf("total value: %d\n", total)

	totalXMas := getXMas(letterGrid)
	fmt.Printf("total x-mas: %d\n", totalXMas)
}

func getForward(g [][]rune) int {
	totalXmas := 0
	for _, row := range g {
		for i := 0; i < len(row); i++ {
			if string(row[i]) == "X" && i+3 < len(row) {
				if string(row[i+1]) == "M" &&
					string(row[i+2]) == "A" &&
					string(row[i+3]) == "S" {
					totalXmas++
				}
			}
		}
	}
	return totalXmas
}

func getBackward(g [][]rune) int {
	totalXmas := 0
	for _, row := range g {
		for i := len(row) - 1; i >= 0; i-- {
			if string(row[i]) == "X" && i-3 >= 0 {
				if string(row[i-1]) == "M" &&
					string(row[i-2]) == "A" &&
					string(row[i-3]) == "S" {
					totalXmas++
				}
			}
		}
	}
	return totalXmas
}

func getDownward(g [][]rune) int {
	totalXmas := 0
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if string(g[y][x]) == "X" && y+3 < len(g) {
				if string(g[y+1][x]) == "M" &&
					string(g[y+2][x]) == "A" &&
					string(g[y+3][x]) == "S" {
					totalXmas++
				}
			}
		}
	}
	return totalXmas
}

func getUpward(g [][]rune) int {
	totalXmas := 0
	for y := len(g) - 1; y >= 0; y-- {
		for x := 0; x < len(g[y]); x++ {
			if string(g[y][x]) == "X" && y-3 >= 0 {
				if string(g[y-1][x]) == "M" &&
					string(g[y-2][x]) == "A" &&
					string(g[y-3][x]) == "S" {
					totalXmas++
				}
			}
		}
	}
	return totalXmas
}

func getDiagonalDown(g [][]rune) int {
	totalXmas := 0
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			// right
			if string(g[y][x]) == "X" && y+3 < len(g) && x+3 < len(g[y]) {
				if string(g[y+1][x+1]) == "M" &&
					string(g[y+2][x+2]) == "A" &&
					string(g[y+3][x+3]) == "S" {
					totalXmas++
				}
			}
			// left
			if string(g[y][x]) == "X" && y+3 < len(g) && x-3 >= 0 {
				if string(g[y+1][x-1]) == "M" &&
					string(g[y+2][x-2]) == "A" &&
					string(g[y+3][x-3]) == "S" {
					totalXmas++
				}
			}
		}
	}
	return totalXmas
}

func getDiagonalUp(g [][]rune) int {
	totalXmas := 0
	for y := len(g) - 1; y >= 0; y-- {
		for x := 0; x < len(g[y]); x++ {
			// right
			if string(g[y][x]) == "X" && y-3 >= 0 && x+3 < len(g[y]) {
				if string(g[y-1][x+1]) == "M" &&
					string(g[y-2][x+2]) == "A" &&
					string(g[y-3][x+3]) == "S" {
					totalXmas++
				}
			}
			// left
			if string(g[y][x]) == "X" && y-3 >= 0 && x-3 >= 0 {
				if string(g[y-1][x-1]) == "M" &&
					string(g[y-2][x-2]) == "A" &&
					string(g[y-3][x-3]) == "S" {
					totalXmas++
				}
			}
		}
	}
	return totalXmas
}

func getXMas(g [][]rune) int {
	totalXmas := 0
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if string(g[y][x]) != "A" {
				continue
			}
			if x == 0 || x == len(g[y])-1 {
				continue
			}
			if y == 0 || y == len(g)-1 {
				continue
			}
			// potential XMAS - A
			valA := string([]rune{
				g[y-1][x-1],
				g[y][x],
				g[y+1][x+1],
			})
			valB := string([]rune{
				g[y+1][x-1],
				g[y][x],
				g[y-1][x+1],
			})

			if (valA == "MAS" || valA == "SAM") &&
				(valB == "MAS" || valB == "SAM") {
				totalXmas++
			}
		}
	}
	return totalXmas
}

func getLetterGrid(data []string) [][]rune {
	res := make([][]rune, 0)
	for _, r := range data {
		chars := []rune(r)
		res = append(res, chars)
	}
	return res
}
