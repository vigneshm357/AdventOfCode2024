package day2

import (
	"aoc/helpers"
	"fmt"
	"math"
)

type Day2 struct {
	lists [][]int
}

func NewDay2Handler() *Day2 {
	lines := helpers.ReadFromFile("day2/input.txt")
	lists := helpers.GetHorizontalLists(lines)
	return &Day2{lists: lists}
}

func isStrictlyDecreasing(list []int) bool {
	for i := range len(list) - 1 {
		if list[i] <= list[i+1] {
			return false
		}
	}
	return true
}

func isStrictlyIncreasing(list []int) bool {
	for i := range len(list) - 1 {
		if list[i] >= list[i+1] {
			return false
		}
	}
	return true
}

func isMaxDifferenceLessThan3(list []int) bool {
	for i := range len(list) - 1 {
		if int(math.Abs(float64(list[i]-list[i+1]))) > 3 {
			return false
		}
	}
	return true
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func (d *Day2) day2Prob1() {
	lists := d.lists

	ans := 0
	for i, _ := range lists {
		if isStrictlyDecreasing(lists[i]) || isStrictlyIncreasing(lists[i]) {
			if isMaxDifferenceLessThan3(lists[i]) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func (d *Day2) day2Prob2() {
	lists := d.lists

	ans := 0
	for i, _ := range lists {
		for j, _ := range lists[i] {
			cpy := make([]int, len(lists[i]))
			copy(cpy, lists[i])
			targetList := RemoveIndex(cpy, j)
			if isStrictlyDecreasing(targetList) || isStrictlyIncreasing(targetList) {
				if isMaxDifferenceLessThan3(targetList) {
					ans++
					break
				}
			}

		}
	}
	fmt.Println(ans)

}

func (d *Day2) Solve() {
	d.day2Prob1()
	d.day2Prob2()
}
