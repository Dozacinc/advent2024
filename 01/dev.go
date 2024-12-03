package m01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Dozacinc/advent2024/util"
)

func Run() {
	content, err := util.GetFileContentLines("./data/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	l1, l2, totalDistance := getTotalDistance(content)
	fmt.Printf("total distance: %d\n", totalDistance)
	similarityScore := getSimilarityScore(l1, l2)
	fmt.Printf("score: %d\n", similarityScore)
}

func getLists(data []string) ([]int, []int) {
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, c := range data {
		rowData := strings.Split(c, "   ")
		l1v, _ := strconv.Atoi(rowData[0])
		l2v, _ := strconv.Atoi(rowData[1])
		l1 = append(l1, l1v)
		l2 = append(l2, l2v)
	}
	return l1, l2
}

func getTotalDistance(content []string) ([]int, []int, int) {
	l1, l2 := getLists(content)
	sortData(&l1)
	sortData(&l2)

	totalDistance := 0
	for i := range l1 {
		x := l1[i]
		y := l2[i]
		diff := y - x
		if x > y {
			diff = diff * -1
		}
		totalDistance += diff
	}
	return l1, l2, totalDistance
}

func getSimilarityScore(l1 []int, l2 []int) int {
	freq := getFreq(l2)
	totalScore := 0
	for _, k := range l1 {
		if val, has := freq[k]; !has {
			// not exist, no change
		} else {
			score := val * k
			totalScore += score
		}
	}

	return totalScore
}

func getFreq(l []int) map[int]int {
	mp := make(map[int]int)
	for _, v := range l {
		if val, has := mp[v]; !has {
			mp[v] = 1
		} else {
			mp[v] = val + 1
		}
	}
	return mp
}

func sortData(list *[]int) {
	swapped := false
	n := len(*list)
	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if (*list)[j] > (*list)[j+1] {
				val := (*list)[j]
				(*list)[j] = (*list)[j+1]
				(*list)[j+1] = val
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
