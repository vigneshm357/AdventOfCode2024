package day1

import (
	"aoc/helpers"
	"fmt"
	"math"
	"sort"
)

type Day1 struct {
	listA []int
	listB []int
}

func NewDay1Handler() *Day1 {
	lines := helpers.ReadFromFile("day1/input.txt")
	listA, listB := helpers.GetVerticalLists(lines)
	return &Day1{listA: listA, listB: listB}
}

func (d *Day1) day1Prob1() {
	listA, listB := d.listA, d.listB
	sort.Ints(listA[:])
	sort.Ints(listB[:])
	sum := 0
	for i, _ := range listA {
		sum += int(math.Abs(float64(listA[i] - listB[i])))
	}
	fmt.Println(sum)
}

func (d *Day1) day1Prob2() {
	listA, listB := d.listA, d.listB
	myMap := make(map[int]int, len(listB))
	for i, _ := range listB {
		val, _ := myMap[listB[i]]
		myMap[listB[i]] = val + 1
	}
	sum := 0
	for i, _ := range listA {
		sum += listA[i] * (myMap[listA[i]])
	}
	fmt.Println(sum)
}

func (d *Day1) Solve() {
	d.day1Prob1()
	d.day1Prob2()
}
