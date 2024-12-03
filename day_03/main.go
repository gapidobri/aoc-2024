package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"regexp"
	"strings"
)

func main() {
	part1()
	part2()
}

func calculateSum(data string) int {
	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := r.FindAllStringSubmatch(data, -1)

	sum := 0
	for _, match := range matches {
		sum += utils.ParseInt(match[1]) * utils.ParseInt(match[2])
	}
	return sum
}

func part1() {
	data := utils.ReadFile("day_03/input.txt")

	sum := calculateSum(data)

	fmt.Println(sum)
}

func part2() {
	data := utils.ReadFile("day_03/input.txt")

	parts := strings.Split(data, "do()")

	sum := 0
	for _, part := range parts {
		valid := strings.Split(part, "don't()")[0]
		sum += calculateSum(valid)
	}

	fmt.Println(sum)
}
