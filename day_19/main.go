package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	data := utils.ReadFile("day_19/input.txt")

	parts := strings.Split(data, "\n\n")
	towels := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")

	count := 0
	for _, design := range designs {
		if dfs1(towels, design) {
			count++
		}
	}

	fmt.Println(count)
}

var cache = map[string]bool{}

func dfs1(towels []string, design string) bool {
	if val, ok := cache[design]; ok {
		return val
	}
	if len(design) == 0 {
		cache[design] = true
		return true
	}
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			if dfs1(towels, design[len(towel):]) {
				cache[design] = true
				return true
			}
		}
	}
	cache[design] = false
	return false
}

func part2() {
	data := utils.ReadFile("day_19/input.txt")

	parts := strings.Split(data, "\n\n")
	towels := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")

	count := 0
	for _, design := range designs {
		count += dfs2(towels, design)
	}

	fmt.Println(count)
}

var cache2 = map[string]int{}

func dfs2(towels []string, design string) int {
	if val, ok := cache2[design]; ok {
		return val
	}
	if len(design) == 0 {
		cache2[design] = 1
		return 1
	}
	count := 0
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			count += dfs2(towels, design[len(towel):])
		}
	}
	cache2[design] = count
	return count
}
