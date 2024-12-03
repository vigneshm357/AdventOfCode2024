package helpers

import (
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func ReadFromFile(file string) (lines []string) {
	data, err := os.ReadFile(file)
	checkErr(err)
	lines = strings.Split(string(data), "\n")
	return
}

func GetVerticalLists(lines []string) (listA []int, listB []int) {
	listA = make([]int, len(lines))
	listB = make([]int, len(lines))
	var err error
	for i, val := range lines {
		if len(val) == 0 {
			continue
		}
		pairs := strings.Split(lines[i], "   ")
		listA[i], err = strconv.Atoi(pairs[0])
		checkErr(err)
		listB[i], err = strconv.Atoi(pairs[1])
		checkErr(err)
	}
	return
}

func GetHorizontalLists(lines []string) (lists [][]int) {
	lists = make([][]int, len(lines))
	for i, _ := range lists {
		strList := strings.Split(lines[i], " ")
		lists[i] = make([]int, len(strList))
	}
	var err error
	for i, _ := range lines {
		strList := strings.Split(lines[i], " ")
		for j, _ := range strList {
			lists[i][j], err = strconv.Atoi(strList[j])
			checkErr(err)
		}
	}
	return
}
