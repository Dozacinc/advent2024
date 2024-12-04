package m03

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/Dozacinc/advent2024/util"
)

func Run() {
	content, err := util.GetFileContentLines("./data/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	dataScans := getAllMultis(content)

	totalMultiplications := getMultiplications(dataScans)
	fmt.Printf("total mults value: %d\n", totalMultiplications)
	totalDoDontMultis := getDoDontMultis(dataScans)
	fmt.Printf("total do dont mults value: %d\n", totalDoDontMultis)

}

func getMultiplications(dataScans []DataScan) int {
	totalMultiplications := 0
	for _, ds := range dataScans {
		for _, mult := range ds.Multis {
			v1, v2 := getNumbers(mult)
			totalMultiplications += v1 * v2
		}
	}
	return totalMultiplications
}

func getDoDontMultis(dataScans []DataScan) int {
	initDo := true
	totalMulti := 0
	for _, ds := range dataScans {
		val, outDo := getDoDontMultiByScan(ds, initDo)
		initDo = outDo
		totalMulti += val
	}
	return totalMulti
}

func getDoDontMultiByScan(ds DataScan, initDo bool) (int, bool) {
	multiMp := make(map[int]string)
	for id, index := range ds.MultisIds {
		multiMp[index] = ds.Multis[id]
	}
	mp := make(map[int]int)
	for _, doId := range ds.DoIds {
		mp[doId] = 1
	}
	for _, dontId := range ds.DontIds {
		mp[dontId] = -1
	}
	for _, multiId := range ds.MultisIds {
		mp[multiId] = 0
	}
	keys := make([]int, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	countMulti := initDo
	multiList := make([]string, 0)
	for _, k := range keys {
		switch mp[k] {
		case -1:
			countMulti = false
		case 1:
			countMulti = true
		case 0:
			if countMulti {
				multiList = append(multiList, multiMp[k])
			}
		}
	}
	// now multiList have all indexes of multis that are legal
	totalMulti := 0
	for _, val := range multiList {
		v1, v2 := getNumbers(val)
		totalMulti += v1 * v2
	}
	return totalMulti, countMulti
}

func getNumbers(s string) (int, int) {
	re := regexp.MustCompile(`\d+`)
	nums := re.FindAllString(s, 2)
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return n1, n2
}

type DataScan struct {
	Multis    []string
	MultisIds []int
	DoIds     []int
	DontIds   []int
}

func getAllMultis(data []string) []DataScan {
	res := make([]DataScan, 0)
	reMul := regexp.MustCompile(`mul\(\d+,\d+\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don\'t\(\)`)

	for _, r := range data {
		ds := DataScan{}
		// multis
		ds.Multis = reMul.FindAllString(r, -1)
		multiIds := make([]int, 0)
		for _, e := range reMul.FindAllStringIndex(r, -1) {
			multiIds = append(multiIds, e[0])
		}
		ds.MultisIds = multiIds
		// dos
		doIds := make([]int, 0)
		for _, e := range reDo.FindAllStringIndex(r, -1) {
			doIds = append(doIds, e[0])
		}
		ds.DoIds = doIds
		// dont
		dontIds := make([]int, 0)
		for _, e := range reDont.FindAllStringIndex(r, -1) {
			dontIds = append(dontIds, e[0])
		}
		ds.DontIds = dontIds

		res = append(res, ds)
	}
	return res
}
