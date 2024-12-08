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

func (p pos) Diff(other pos) pos {
	return pos{other.x - p.x, other.y - p.y}
}

func (p pos) Add(other pos) pos {
	return pos{p.x + other.x, p.y + other.y}
}

func main() {
	part1()
	part2()
}

func getMap() (map[rune]mapset.Set[pos], int, int) {
	data := utils.ReadFile("day_08/input.txt")
	m := map[rune]mapset.Set[pos]{}
	lines := strings.Split(data, "\n")

	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			if m[char] == nil {
				m[char] = mapset.NewSet[pos]()
			}
			m[char].Add(pos{x, y})
		}
	}
	return m, len(lines[0]), len(lines)
}

func part1() {
	m, width, height := getMap()
	anodes := mapset.NewSet[pos]()

	for _, positions := range m {
		for pos1 := range positions.Iter() {
			for pos2 := range positions.Iter() {
				if pos1 == pos2 {
					continue
				}
				anode1 := pos2.Add(pos1.Diff(pos2))
				if !(anode1.x < 0 || anode1.x >= width || anode1.y < 0 || anode1.y >= height) {
					anodes.Add(anode1)
				}
				anode2 := pos1.Add(pos2.Diff(pos1))
				if !(anode2.x < 0 || anode2.x >= width || anode2.y < 0 || anode2.y >= height) {
					anodes.Add(anode2)
				}
			}
		}
	}

	fmt.Println(anodes.Cardinality())
}

func part2() {
	m, width, height := getMap()
	anodes := mapset.NewSet[pos]()

	for _, positions := range m {
		for pos1 := range positions.Iter() {
			anodes.Add(pos1)
			for pos2 := range positions.Iter() {
				if pos1 == pos2 {
					continue
				}
				diff := pos1.Diff(pos2)
				anode := pos2
				for {
					anode = anode.Add(diff)
					if anode.x < 0 || anode.x >= width || anode.y < 0 || anode.y >= height {
						break
					}
					anodes.Add(anode)
				}
				diff = pos2.Diff(pos1)
				anode = pos1
				for {
					anode = anode.Add(diff)
					if anode.x < 0 || anode.x >= width || anode.y < 0 || anode.y >= height {
						break
					}
					anodes.Add(anode)
				}
			}
		}
	}

	fmt.Println(anodes.Cardinality())
}
