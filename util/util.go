package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFileContentLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content, nil
}

func KeepCharInString(input string, keepChar string) string {
	resultRunes := make([]rune, 0)
	inputRunes := []rune(input)
	keepRunes := []rune(keepChar)
	for i := 0; i < len(inputRunes); i++ {
		for k := 0; k < len(keepRunes); k++ {
			if inputRunes[i] == keepRunes[k] {
				resultRunes = append(resultRunes, inputRunes[i])
				break
			}
		}
	}
	resString := string(resultRunes)
	return resString
}

func ConvertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

func ConvertStringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ReplaceOnceInString(input string, old string, new string) string {
	return strings.Replace(input, old, new, 1)
}

func SplitStringToInts(s string, d string) []int {
	res := make([]int, 0)
	slice := strings.Split(s, d)
	for _, v := range slice {
		if len(v) == 0 {
			continue
		}
		vi, err := strconv.Atoi(v)
		if err != nil {
			panic("int conv failed")
		}
		res = append(res, vi)
	}
	return res
}
