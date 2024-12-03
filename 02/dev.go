package m02

import (
	"fmt"

	"github.com/Dozacinc/advent2024/util"
)

func Run() {
	content, err := util.GetFileContentLines("./data/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	totalSafe := 0
	totalMarginSafe := 0
	reports := getReports(content)
	fmt.Printf("total reports: %d\n", len(reports))
	for _, r := range reports {
		if checkSafeReport(r) {
			totalSafe++
			totalMarginSafe++
		} else if checkMarginSafeReport(r) {
			totalMarginSafe++
		}
	}
	fmt.Printf("total safe: %d\n", totalSafe)
	fmt.Printf("total margin safe: %d\n", totalMarginSafe)
}

func checkMarginSafeReport(r []int) bool {

	marginSafe := false
	for i := 0; i < len(r); i++ {
		arr := getArrayWithoutElement(r, i)
		if checkSafeReport(arr) {
			marginSafe = true
			break
		}
	}
	return marginSafe

}

func getArrayWithoutElement(r []int, x int) []int {
	res := make([]int, 0)
	for i, v := range r {
		if i == x {
			continue
		}
		res = append(res, v)
	}
	return res
}

func checkSafeReport(r []int) bool {

	pairs := len(r) - 1
	lastDir := 0
	for i := 0; i < pairs; i++ {
		x := r[i]
		y := r[i+1]

		dir := getDir(x, y)
		absDiff := (x - y) * dir

		if !(absDiff >= 1 && absDiff <= 3) {
			return false
		}

		if lastDir != 0 && dir != lastDir {
			return false
		}
		lastDir = dir
	}
	return true
}

func getDir(x, y int) int {
	if x == y {
		return 0
	}
	if x > y {
		return 1
	}
	return -1
}

func getReports(data []string) [][]int {
	res := make([][]int, 0)
	for _, v := range data {
		res = append(res, util.SplitStringToInts(v, " "))
	}
	return res
}
