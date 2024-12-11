package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"strconv"
	"strings"
)

func main() {
	data := utils.ReadFile("day_11/input.txt")

	stones := strings.Split(data, " ")

	count := 0
	for _, stone := range stones {
		count += transform(75, stone)
	}

	fmt.Println(count)
}

type p struct {
	depth int
	stone string
}

var cache = map[p]int{}

func transform(depth int, stone string) int {
	params := p{depth, stone}
	if depth == 0 {
		return 1
	}
	if res, ok := cache[params]; ok {
		return res
	}
	if stone == "0" {
		res := transform(depth-1, "1")
		cache[params] = res
		return res
	}
	if len(stone)%2 == 0 {
		newLen := len(stone) / 2
		right := strings.TrimLeft(stone[newLen:], "0")
		if right == "" {
			right = "0"
		}
		res := transform(depth-1, stone[:newLen]) + transform(depth-1, right)
		cache[params] = res
		return res
	}
	res := transform(depth-1, strconv.Itoa(utils.ParseInt(stone)*2024))
	cache[params] = res
	return res
}
