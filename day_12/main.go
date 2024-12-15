package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

type pos struct {
	x, y int
}

func (p pos) add(other pos) pos {
	return pos{p.x + other.x, p.y + other.y}
}

var checks = []pos{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func main() {
	part1()
	part2()
}

func inBounds(rows []string, p pos) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(rows[0]) && p.y < len(rows)
}

func part1() {
	data := utils.ReadFile("day_12/input.txt")

	rows := strings.Split(data, "\n")

	visited := mapset.NewSet[pos]()
	sum := 0
	for y, row := range rows {
		for x := range row {
			area, fences := find(rows, visited, pos{x, y})
			sum += area * fences
		}
	}

	fmt.Println(sum)
}

func find(rows []string, visited mapset.Set[pos], p pos) (int, int) {
	if visited.Contains(p) {
		return 0, 0
	}
	visited.Add(p)
	curr := rows[p.y][p.x]
	fences := 4
	area := 1

	for _, check := range checks {
		newPos := p.add(check)
		if !inBounds(rows, newPos) {
			continue
		}
		if rows[newPos.y][newPos.x] == curr {
			fences--
			if !visited.Contains(newPos) {
				a, f := find(rows, visited, newPos)
				area += a
				fences += f
			}
		}
	}

	return area, fences
}

func part2() {
	data := utils.ReadFile("day_12/input.txt")

	rows := strings.Split(data, "\n")

	visited := mapset.NewSet[pos]()
	sum := 0
	for y, row := range rows {
		for x := range row {
			area, fences := find2(rows, visited, pos{x, y})
			sum += area * fences
		}
	}

	fmt.Println(sum)
}

func find2(rows []string, visited mapset.Set[pos], p pos) (int, int) {
	if visited.Contains(p) {
		return 0, 0
	}
	visited.Add(p)

	corners := 0
	area := 1

	var edges [4]bool
	for i, check := range checks {
		newPos := p.add(check)
		if !inBounds(rows, newPos) {
			edges[i] = true
			continue
		}
		if rows[newPos.y][newPos.x] != rows[p.y][p.x] {
			edges[i] = true
		} else {
			a, c := find2(rows, visited, newPos)
			area += a
			corners += c
		}
	}

	for i := 0; i < 4; i++ {
		if edges[i] && edges[(i+1)%4] {
			corners++
		} else if !edges[i] && !edges[(i+1)%4] {
			diag := p.add(checks[i]).add(checks[(i+1)%4])
			if !inBounds(rows, diag) {
				continue
			}
			if rows[diag.y][diag.x] != rows[p.y][p.x] {
				corners++
			}
		}
	}

	return area, corners
}
