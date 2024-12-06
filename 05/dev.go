package m05

import (
	"fmt"
	"regexp"

	"github.com/Dozacinc/advent2024/util"
)

func Run() {
	content, err := util.GetFileContentLines("./data/data.txt")
	if err != nil {
		fmt.Println(err)
	}
	ruleSet := getRuleSet(content)
	reportSet := getReportSet(content)
	validReports := make([]ReportSet, 0)
	editedReports := make([]ReportSet, 0)
	for i := 0; i < len(reportSet); i++ {
		if _, valid := checkValidReport(&ruleSet, &reportSet[i]); valid {
			validReports = append(validReports, reportSet[i])
		} else {
			editedReport := getFixedReport(&ruleSet, &reportSet[i])
			editedReports = append(editedReports, editedReport)
		}
	}

	totalValidMidValue := 0
	fmt.Printf("valid reports: %d\n", len(validReports))
	for _, r := range validReports {
		if checkOdd(len(r)) {
			totalValidMidValue += r[len(r)/2]
		} else {
			panic("ASDASD not odd")
		}
	}
	fmt.Printf("total valid mid-value: %d\n", totalValidMidValue)

	totalEditedMidValue := 0
	fmt.Printf("edited reports: %d\n", len(editedReports))
	for _, r := range editedReports {
		if checkOdd(len(r)) {
			totalEditedMidValue += r[len(r)/2]
		} else {
			panic("ASDASD not odd")
		}
	}
	fmt.Printf("total edited mid-value: %d\n", totalEditedMidValue)

}

type RuleSet map[int]map[int]interface{}
type ReportSet []int

func (r ReportSet) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func getFixedReport(ruleSet *RuleSet, r *ReportSet) ReportSet {
	tmp := make(ReportSet, 0)
	tmp = append(tmp, *r...)
	for {
		failedPage, valid := checkValidReport(ruleSet, &tmp)
		if valid {
			break
		}
		for i := 0; i < len(tmp); i++ {
			if tmp[i] == failedPage && i < len(tmp)-1 {
				tmp.Swap(i, i+1)
			}
		}
	}
	return tmp
}

func checkValidReport(ruleSet *RuleSet, r *ReportSet) (int, bool) {
	checkedPages := make(map[int]interface{})
	for _, c := range *r {
		valRule, hasRule := (*ruleSet)[c]
		if !hasRule {
			continue
		}
		for k := range valRule {
			if _, hasKey := checkedPages[k]; hasKey {
				return k, false
			}
		}
		checkedPages[c] = nil
	}
	return -1, true
}

func checkOdd(i int) bool {
	return i%2 != 0
}

func getRuleSet(data []string) RuleSet {
	re := regexp.MustCompile(`\d+\|\d+`)
	mp := make(map[int]map[int]interface{})
	for _, c := range data {
		if !re.MatchString(c) {
			break
		}
		vals := util.SplitStringToInts(c, "|")
		_, hasMap := mp[vals[0]]
		if !hasMap {
			mp[vals[0]] = make(map[int]interface{})
		}
		_, hasPrec := mp[vals[0]][vals[1]]
		if !hasPrec {
			mp[vals[0]][vals[1]] = nil
		}
	}
	return mp
}

func getReportSet(data []string) []ReportSet {
	re := regexp.MustCompile(`\d+\|\d+`)
	res := make([]ReportSet, 0)
	for _, c := range data {
		if re.MatchString(c) || c == "" {
			continue
		}
		vals := util.SplitStringToInts(c, ",")
		res = append(res, vals)
	}
	return res
}
