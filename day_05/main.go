package main

import (
	"fmt"
	"github.com/deckarep/golang-set/v2"
	"github.com/gapidobri/aoc-2024/utils"
	"slices"
	"strings"
)

type rulesMap map[string]mapset.Set[string]

func main() {
	part1()
	part2()
}

func getData() ([]string, rulesMap, rulesMap) {
	data := utils.ReadFile("day_05/input.txt")

	parts := strings.Split(data, "\n\n")

	fwdRules := rulesMap{}
	bkwRules := rulesMap{}

	rules := strings.Split(parts[0], "\n")
	for _, rule := range rules {
		p := strings.Split(rule, "|")
		if fwdRules[p[0]] == nil {
			fwdRules[p[0]] = mapset.NewSet[string](p[1])
		} else {
			fwdRules[p[0]].Add(p[1])
		}
		if bkwRules[p[1]] == nil {
			bkwRules[p[1]] = mapset.NewSet[string](p[0])
		} else {
			bkwRules[p[1]].Add(p[0])
		}
	}

	return strings.Split(parts[1], "\n"), fwdRules, bkwRules
}

func sortPages(fwdRules rulesMap, bkwRules rulesMap, pages []string) bool {
	valid := true
	slices.SortFunc(pages, func(a string, b string) int {
		rule, ok := fwdRules[b]
		if ok && rule.Contains(a) {
			return 1
		}
		rule, ok = bkwRules[a]
		if ok && rule.Contains(b) {
			return 1
		}
		valid = false
		return -1
	})
	return valid
}

func part1() {
	updates, fwdRules, bkwRules := getData()

	sum := 0
	for _, update := range updates {
		pages := strings.Split(update, ",")

		if sortPages(fwdRules, bkwRules, pages) {
			sum += utils.ParseInt(pages[len(pages)/2])
		}
	}

	fmt.Println(sum)
}

func part2() {
	updates, fwdRules, bkwRules := getData()

	sum := 0
	for _, update := range updates {
		pages := strings.Split(update, ",")

		if !sortPages(fwdRules, bkwRules, pages) {
			sum += utils.ParseInt(pages[len(pages)/2])
		}
	}

	fmt.Println(sum)
}
