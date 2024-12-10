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
	part2()
}

func getMap() [][]int {
	data := utils.ReadFile("day_10/input.txt")
	var m [][]int
	for i, row := range strings.Split(data, "\n") {
		m = append(m, make([]int, len(row)))
		for j, char := range row {
			m[i][j] = int(char - '0')
		}
	}
	return m
}

func find1(m [][]int, reachable mapset.Set[pos], currPos pos) {
	curr := m[currPos.y][currPos.x]
	if curr == 9 {
		reachable.Add(currPos)
		return
	}
	next := curr + 1
	if currPos.y > 0 && m[currPos.y-1][currPos.x] == next {
		find1(m, reachable, pos{currPos.x, currPos.y - 1})
	}
	if currPos.y+1 < len(m) && m[currPos.y+1][currPos.x] == next {
		find1(m, reachable, pos{currPos.x, currPos.y + 1})
	}
	if currPos.x > 0 && m[currPos.y][currPos.x-1] == next {
		find1(m, reachable, pos{currPos.x - 1, currPos.y})
	}
	if currPos.x+1 < len(m[0]) && m[currPos.y][currPos.x+1] == next {
		find1(m, reachable, pos{currPos.x + 1, currPos.y})
	}
}

func part1() {
	m := getMap()

	sum := 0
	for y, row := range m {
		for x, height := range row {
			if height == 0 {
				reachable := mapset.NewSet[pos]()
				find1(m, reachable, pos{x, y})
				sum += reachable.Cardinality()
			}
		}
	}

	fmt.Println(sum)
}

func find2(m [][]int, currPos pos) int {
	curr := m[currPos.y][currPos.x]
	if curr == 9 {
		return 1
	}
	next := curr + 1
	sum := 0
	if currPos.y > 0 && m[currPos.y-1][currPos.x] == next {
		sum += find2(m, pos{currPos.x, currPos.y - 1})
	}
	if currPos.y+1 < len(m) && m[currPos.y+1][currPos.x] == next {
		sum += find2(m, pos{currPos.x, currPos.y + 1})
	}
	if currPos.x > 0 && m[currPos.y][currPos.x-1] == next {
		sum += find2(m, pos{currPos.x - 1, currPos.y})
	}
	if currPos.x+1 < len(m[0]) && m[currPos.y][currPos.x+1] == next {
		sum += find2(m, pos{currPos.x + 1, currPos.y})
	}
	return sum
}

func part2() {
	m := getMap()

	sum := 0
	for y, row := range m {
		for x, height := range row {
			if height == 0 {
				sum += find2(m, pos{x, y})
			}
		}
	}

	fmt.Println(sum)
}
