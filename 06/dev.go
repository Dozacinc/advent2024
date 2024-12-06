package m06

import (
	"fmt"
	"strings"

	"github.com/Dozacinc/advent2024/util"
)

func Run() {
	content, err := util.GetFileContentLines("./data/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	movementMap, startY, startX := getMovementMap(content)
	fmt.Printf("Start: %d, %d\n", startY, startX)
	totalTiles := resolveMovement(movementMap, startY, startX)
	fmt.Printf("total tiles: %d\n", totalTiles)

	totalHits := 0
	for _, r := range movementMap {
		for _, c := range r {
			if getMapPres(c) == "*" {
				totalHits++
			}
		}
	}
	fmt.Printf("Total stars: %d\n", totalHits)

	printMap(movementMap)
}

func printMap(g Map) {
	fmt.Println("MAP: ")
	for ri, r := range g {
		fmt.Printf("%d\t", ri)
		for _, c := range r {
			fmt.Printf("%s", getMapPres(c))

		}
		fmt.Println("")
	}
}

func getMapPres(i int) string {
	switch i {
	case -1:
		return "#"
	case 0:
		return "."
	}
	return "*"
}

type Map [][]int

func resolveMovement(g Map, origY int, origX int) int {
	direction := 0 // N, E, S, W
	endOfMap := false
	totalTiles := 0
	y, x := origY, origX
	for {
		switch direction {
		case 0:
			tiles, hitObstacle := g.MoveUp(y, x)
			if hitObstacle {
				direction = 1
				y -= tiles
			} else {
				endOfMap = true
			}
			totalTiles += tiles
		case 1:
			tiles, hitObstacle := g.MoveRight(y, x)
			if hitObstacle {
				direction = 2
				x += tiles
			} else {
				endOfMap = true
			}
			totalTiles += tiles
		case 2:
			tiles, hitObstacle := g.MoveDown(y, x)
			if hitObstacle {
				direction = 3
				y += tiles
			} else {
				endOfMap = true
			}
			totalTiles += tiles
		case 3:
			tiles, hitObstacle := g.MoveLeft(y, x)
			if hitObstacle {
				direction = 0
				x -= tiles
			} else {
				endOfMap = true
			}
			totalTiles += tiles
		}
		if endOfMap {
			break
		}
	}
	return totalTiles
}

func (g Map) MoveUp(startY, startX int) (int, bool) {
	movementId := g[startY][startX] + 1
	totalMoves := 0
	for y := startY; y > 0; y-- {
		if g[y-1][startX] >= 0 {
			g[y-1][startX] = movementId
			totalMoves++
		} else {
			return totalMoves, true
		}
	}
	return totalMoves, false
}

func (g Map) MoveDown(startY, startX int) (int, bool) {
	movementId := g[startY][startX] + 1
	totalMoves := 0
	for y := startY; y < len(g)-1; y++ {
		if g[y+1][startX] >= 0 {
			g[y+1][startX] = movementId
			totalMoves++
		} else {
			return totalMoves, true
		}
	}
	return totalMoves, false
}

func (g Map) MoveLeft(startY, startX int) (int, bool) {
	movementId := g[startY][startX] + 1
	totalMoves := 0
	for x := startX; x > 0; x-- {
		if g[startY][x-1] >= 0 {
			g[startY][x-1] = movementId
			totalMoves++
		} else {
			return totalMoves, true
		}
	}
	return totalMoves, false
}

func (g Map) MoveRight(startY, startX int) (int, bool) {
	movementId := g[startY][startX] + 1
	totalMoves := 0
	for x := startX; x < len(g[startY])-1; x++ {
		if g[startY][x+1] >= 0 {
			g[startY][x+1] = movementId
			totalMoves++
		} else {
			return totalMoves, true
		}
	}
	return totalMoves, false
}

func getMovementMap(data []string) (Map, int, int) {
	res := make(Map, 0)
	x, y := 0, 0
	for rI, r := range data {
		chars := strings.Split(r, "")
		arr := make([]int, 0)
		for i, c := range chars {
			switch c {
			case ".":
				arr = append(arr, 0)
			case "#":
				arr = append(arr, -1)
			case "^":
				arr = append(arr, 1)
				x = i
				y = rI
			}
		}
		res = append(res, arr)
	}
	return res, y, x
}
