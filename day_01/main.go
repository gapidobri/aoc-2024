package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"regexp"
	"slices"
)

func main() {
	part1()
	part2()
}

func parseLists() ([]int, []int) {
	data := utils.ReadFile("./day_01/input.txt")

	regex := regexp.MustCompile("[0-9]+")
	ids := regex.FindAllString(data, -1)

	var left, right []int
	for i, id := range ids {
		if i%2 == 0 {
			left = append(left, utils.ParseInt(id))
		} else {
			right = append(right, utils.ParseInt(id))
		}
	}

	return left, right
}

func part1() {
	left, right := parseLists()

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i, id := range left {
		sum += utils.IntAbs(id - right[i])
	}

	fmt.Println(sum)
}

func part2() {
	left, right := parseLists()

	sum := 0
	for _, lid := range left {
		count := 0
		for _, rid := range right {
			if lid == rid {
				count++
			}
		}
		sum += lid * count
	}

	fmt.Println(sum)
}
