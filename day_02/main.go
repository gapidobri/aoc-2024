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

func getReports() []string {
	data := utils.ReadFile("day_02/input.txt")
	return strings.Split(data, "\n")
}

func checkReport(levels []string) bool {
	safe := true
	last := 0
	decreasing := false
	for i, strLevel := range levels {
		level := utils.ParseInt(strLevel)
		if i == 0 {
			last = level
			continue
		}
		diff := utils.IntAbs(level - last)
		if diff == 0 || diff > 3 {
			safe = false
			break
		}
		if i == 1 {
			decreasing = last > level
			last = level
			continue
		}
		if decreasing && last < level || !decreasing && last > level {
			safe = false
			break
		}
		last = level
	}
	return safe
}

func part1() {
	reports := getReports()

	safeCount := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")
		if checkReport(levels) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func part2() {
	reports := getReports()

	safeCount := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")

		if checkReport(levels) {
			safeCount++
			continue
		}

		for i := range levels {
			n := make([]string, len(levels))
			copy(n, levels)
			if checkReport(append(n[:i], n[i+1:]...)) {
				safeCount++
				break
			}
		}
	}

	fmt.Println(safeCount)
}
