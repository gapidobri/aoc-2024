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

func main() {
	part1()
}

func find(rows []string, visited mapset.Set[pos], p pos) (int, int) {
	if visited.Contains(p) {
		return 0, 0
	}
	visited.Add(p)
	curr := rows[p.y][p.x]
	fences := 4
	area := 1
	if p.y > 0 && rows[p.y-1][p.x] == curr {
		fences--
		newPos := pos{p.x, p.y - 1}
		if !visited.Contains(newPos) {
			a, f := find(rows, visited, newPos)
			area += a
			fences += f
		}
	}
	if p.y+1 < len(rows) && rows[p.y+1][p.x] == curr {
		fences--
		newPos := pos{p.x, p.y + 1}
		if !visited.Contains(newPos) {
			a, f := find(rows, visited, newPos)
			area += a
			fences += f
		}
	}
	if p.x > 0 && rows[p.y][p.x-1] == curr {
		fences--
		newPos := pos{p.x - 1, p.y}
		if !visited.Contains(newPos) {
			a, f := find(rows, visited, newPos)
			area += a
			fences += f
		}
	}
	if p.x+1 < len(rows) && rows[p.y][p.x+1] == curr {
		fences--
		newPos := pos{p.x + 1, p.y}
		if !visited.Contains(newPos) {
			a, f := find(rows, visited, newPos)
			area += a
			fences += f
		}
	}
	return area, fences
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
