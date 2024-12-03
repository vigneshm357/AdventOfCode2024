package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"fmt"
)

func main() {
	day1.NewDay1Handler().Solve()
	fmt.Println("---------")
	day2.NewDay2Handler().Solve()
	fmt.Println("---------")
	day3.Solve()
	fmt.Println("---------")
}
