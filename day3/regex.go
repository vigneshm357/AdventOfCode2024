package day3

import (
	"aoc/helpers"
	"fmt"
	"regexp"
)

func day3Prob2() {
	lines := helpers.ReadFromFile("day3/input.txt")
	r := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don\'t\(\)`)
	sum := 0
	run := true
	for i, _ := range lines {
		matches := r.FindAllString(lines[i], -1)
		for _, val := range matches {
			if val == "do()" {
				run = true
				continue
			}
			if val == "don't()" {
				run = false
				continue
			}
			if !run {
				continue
			}
			var a, b int
			_, err := fmt.Sscanf(val, "mul(%d,%d)", &a, &b)
			if err != nil {
				panic(err)
			}
			sum += (a * b)
		}
	}
	fmt.Println(sum)
}

func day3Prob1() {
	lines := helpers.ReadFromFile("day3/input.txt")
	r := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	sum := 0
	for i, _ := range lines {
		matches := r.FindAllString(lines[i], -1)
		for _, val := range matches {
			var a, b int
			_, err := fmt.Sscanf(val, "mul(%d,%d)", &a, &b)
			if err != nil {
				panic(err)
			}
			sum += (a * b)
		}
	}
	fmt.Println(sum)
}

func Solve() {
	day3Prob1()
	day3Prob2()
}
