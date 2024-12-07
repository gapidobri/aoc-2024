package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"slices"
	"strconv"
	"strings"
)

const (
	add = iota
	mul
	cat
)

func main() {
	part1()
	part2()
}

func part1() {
	run(false)
}

func part2() {
	run(true)
}

func catInt(a, b int) int {
	return utils.ParseInt(strconv.Itoa(a) + strconv.Itoa(b))
}

func calculate(nums *utils.Stack[int], op int, target int, part2 bool) bool {
	if nums.Len() == 0 {
		return false
	}
	if nums.Len() == 1 {
		return nums.Pop() == target
	}
	if nums.Peek() > target {
		return false
	}
	switch op {
	case add:
		nums.Push(nums.Pop() + nums.Pop())
	case mul:
		nums.Push(nums.Pop() * nums.Pop())
	case cat:
		nums.Push(catInt(nums.Pop(), nums.Pop()))
	}
	return calculate(nums.Clone(), add, target, part2) || calculate(nums.Clone(), mul, target, part2) || (part2 && calculate(nums.Clone(), cat, target, part2))
}

func run(part2 bool) {
	data := utils.ReadFile("day_07/input.txt")

	sum := 0
	for _, line := range strings.Split(data, "\n") {
		p := strings.Split(line, ": ")
		target := utils.ParseInt(p[0])
		var nums []int
		for _, n := range strings.Split(p[1], " ") {
			nums = append(nums, utils.ParseInt(n))
		}

		slices.Reverse(nums)
		stack := utils.NewStack(nums)

		if calculate(stack.Clone(), add, target, part2) || calculate(stack.Clone(), mul, target, part2) || (part2 && calculate(stack.Clone(), cat, target, part2)) {
			sum += target
		}
	}

	fmt.Println(sum)
}
