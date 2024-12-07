package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

type (
	pos struct {
		x, y int
	}
	posdir struct {
		x, y, dir int
	}
)

const (
	up = iota
	right
	down
	left
)

func main() {
	//overengineeredPart1()
	//part1()
	part2()
}

func getRows() [][]rune {
	data := utils.ReadFile("day_06/input.txt")

	var rows [][]rune
	for _, row := range strings.Split(data, "\n") {
		rows = append(rows, []rune(row))
	}

	return rows
}

func getGuard(rows [][]rune) posdir {
	for y, row := range rows {
		for x, char := range row {
			dir := strings.Index("^>v<", string(char))
			if dir != -1 {
				return posdir{x, y, dir}
			}
		}
	}
	return posdir{}
}

func part1() {
	rows := getRows()
	guard := getGuard(rows)

	visited := mapset.NewSet[pos](pos{guard.x, guard.y})

outer:
	for {
		switch guard.dir {
		case up:
			if guard.y == 0 {
				break outer
			}
			if rows[guard.y-1][guard.x] == '#' {
				guard.dir = right
				continue
			}
			guard.y--
		case right:
			if guard.x == len(rows[0])-1 {
				break outer
			}
			if rows[guard.y][guard.x+1] == '#' {
				guard.dir = down
				continue
			}
			guard.x++
		case down:
			if guard.y == len(rows)-1 {
				break outer
			}
			if rows[guard.y+1][guard.x] == '#' {
				guard.dir = left
				continue
			}
			guard.y++
		case left:
			if guard.x == 0 {
				break outer
			}
			if rows[guard.y][guard.x-1] == '#' {
				guard.dir = up
				continue
			}
			guard.x--
		}
		visited.Add(pos{guard.x, guard.y})
	}

	fmt.Println(visited.Cardinality())
}

func simulate(rows [][]rune, guard posdir) bool {
	visited := mapset.NewSet[posdir](guard)

outer:
	for {
		switch guard.dir {
		case up:
			if guard.y == 0 {
				break outer
			}
			if rows[guard.y-1][guard.x] == '#' {
				guard.dir = right
				continue
			}
			guard.y--
		case right:
			if guard.x == len(rows[0])-1 {
				break outer
			}
			if rows[guard.y][guard.x+1] == '#' {
				guard.dir = down
				continue
			}
			guard.x++
		case down:
			if guard.y == len(rows)-1 {
				break outer
			}
			if rows[guard.y+1][guard.x] == '#' {
				guard.dir = left
				continue
			}
			guard.y++
		case left:
			if guard.x == 0 {
				break outer
			}
			if rows[guard.y][guard.x-1] == '#' {
				guard.dir = up
				continue
			}
			guard.x--
		}
		if visited.Contains(guard) {
			return true
		}
		visited.Add(guard)
	}

	return false
}

func part2() {
	data := utils.ReadFile("day_06/input.txt")

	var rows [][]rune
	for _, row := range strings.Split(data, "\n") {
		rows = append(rows, []rune(row))
	}

	guard := getGuard(rows)
	origGuard := guard

	okObst := mapset.NewSet[pos]()

	sum := 0
outer:
	for {
		var rows2 [][]rune
		for i, row := range rows {
			rows2 = append(rows2, []rune{})
			for _, char := range row {
				rows2[i] = append(rows2[i], char)
			}
		}

		skip := false
		var obstacle pos
		switch guard.dir {
		case up:
			if guard.y == 0 {
				break outer
			}
			if rows[guard.y-1][guard.x] == '#' {
				skip = true
				break
			}
			obstacle = pos{guard.x, guard.y - 1}
			rows2[guard.y-1][guard.x] = '#'
		case right:
			if guard.x == len(rows[0])-1 {
				break outer
			}
			if rows[guard.y][guard.x+1] == '#' {
				skip = true
				break
			}
			obstacle = pos{guard.x + 1, guard.y}
			rows2[guard.y][guard.x+1] = '#'
		case down:
			if guard.y == len(rows)-1 {
				break outer
			}
			if rows[guard.y+1][guard.x] == '#' {
				skip = true
				break
			}
			obstacle = pos{guard.x, guard.y + 1}
			rows2[guard.y+1][guard.x] = '#'
		case left:
			if guard.x == 0 {
				break outer
			}
			if rows[guard.y][guard.x-1] == '#' {
				skip = true
				break
			}
			obstacle = pos{guard.x - 1, guard.y}
			rows2[guard.y][guard.x-1] = '#'
		}

		if !skip && simulate(rows2, origGuard) {
			okObst.Add(obstacle)
			sum++
		}

		switch guard.dir {
		case up:
			if guard.y == 0 {
				break outer
			}
			if rows[guard.y-1][guard.x] == '#' {
				guard.dir = right
				continue
			}
			guard.y--
		case right:
			if guard.x == len(rows[0])-1 {
				break outer
			}
			if rows[guard.y][guard.x+1] == '#' {
				guard.dir = down
				continue
			}
			guard.x++
		case down:
			if guard.y == len(rows)-1 {
				break outer
			}
			if rows[guard.y+1][guard.x] == '#' {
				guard.dir = left
				continue
			}
			guard.y++
		case left:
			if guard.x == 0 {
				break outer
			}
			if rows[guard.y][guard.x-1] == '#' {
				guard.dir = up
				continue
			}
			guard.x--
		}
	}
	okObst.Remove(pos{origGuard.x, origGuard.y})

	fmt.Println(okObst.Cardinality())
}
