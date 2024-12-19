package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

type pos struct {
	x, y int
	cost int
}

func (p pos) noCost() pos {
	return pos{p.x, p.y, 0}
}

type move struct {
	pos pos
	f   int
}

func main() {
	part1()
	part2()
}

func part1() {
	data := utils.ReadFile("day_18/input.txt")

	bytes := strings.Split(data, "\n")

	var memory [71][71]rune

	for y, row := range memory {
		for x := range row {
			memory[y][x] = '.'
		}
	}

	for _, b := range bytes[:1024] {
		cord := strings.Split(b, ",")
		x := utils.ParseInt(cord[0])
		y := utils.ParseInt(cord[1])

		memory[y][x] = '#'
	}

	heuristic := func(p pos) int {
		return len(memory[0]) - 1 - p.x + len(memory) - 1 - p.y
	}

	visited := mapset.NewSet[pos]()

	generateMoves := func(p pos) (moves []move) {
		if p.x+1 < len(memory[0]) && memory[p.y][p.x+1] != '#' {
			newPos := pos{p.x + 1, p.y, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}
		if p.x-1 >= 0 && memory[p.y][p.x-1] != '#' {
			newPos := pos{p.x - 1, p.y, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}
		if p.y+1 < len(memory) && memory[p.y+1][p.x] != '#' {
			newPos := pos{p.x, p.y + 1, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}
		if p.y-1 >= 0 && memory[p.y-1][p.x] != '#' {
			newPos := pos{p.x, p.y - 1, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}

		return
	}

	goalCheck := func(p pos) bool {
		return p.x == len(memory[0])-1 && p.y == len(memory)-1
	}

	var moves []move
	p := pos{0, 0, 0}

	for {
		if goalCheck(p) {
			fmt.Println(p.cost)
			break
		}
		gen := generateMoves(p)
		for i1, m1 := range moves {
			for i2, m2 := range gen {
				if m1.pos.x != m2.pos.x || m1.f != m2.f {
					continue
				}
				if m1.f > m2.f {
					moves = append(moves[:i1], moves[i1+1:]...)
				} else {
					gen = append(gen[:i2], gen[i2+1:]...)
				}
			}
		}

		moves = append(moves, gen...)

		p = moves[0].pos
		moves = moves[1:]
		visited.Add(p.noCost())
	}
}

func part2() {
	data := utils.ReadFile("day_18/input.txt")

	bytes := strings.Split(data, "\n")

	var memory [71][71]rune

	for y, row := range memory {
		for x := range row {
			memory[y][x] = '.'
		}
	}

	for _, b := range bytes[:1024] {
		cord := strings.Split(b, ",")
		x := utils.ParseInt(cord[0])
		y := utils.ParseInt(cord[1])

		memory[y][x] = '#'
	}

	heuristic := func(p pos) int {
		return len(memory[0]) - 1 - p.x + len(memory) - 1 - p.y
	}

	visited := mapset.NewSet[pos]()

	generateMoves := func(p pos) (moves []move) {
		if p.x+1 < len(memory[0]) && memory[p.y][p.x+1] != '#' {
			newPos := pos{p.x + 1, p.y, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}
		if p.x-1 >= 0 && memory[p.y][p.x-1] != '#' {
			newPos := pos{p.x - 1, p.y, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}
		if p.y+1 < len(memory) && memory[p.y+1][p.x] != '#' {
			newPos := pos{p.x, p.y + 1, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}
		if p.y-1 >= 0 && memory[p.y-1][p.x] != '#' {
			newPos := pos{p.x, p.y - 1, p.cost + 1}
			if !visited.Contains(newPos.noCost()) {
				moves = append(moves, move{newPos, heuristic(newPos)})
			}
		}

		return
	}

	goalCheck := func(p pos) bool {
		return p.x == len(memory[0])-1 && p.y == len(memory)-1
	}

	hasPath := func() bool {
		visited = mapset.NewSet[pos]()
		var moves []move
		p := pos{0, 0, 0}

		for {
			if goalCheck(p) {
				return true
			}
			gen := generateMoves(p)
			for i1, m1 := range moves {
				for i2, m2 := range gen {
					if m1.pos.x != m2.pos.x || m1.f != m2.f {
						continue
					}
					if m1.f > m2.f {
						moves = append(moves[:i1], moves[i1+1:]...)
					} else {
						gen = append(gen[:i2], gen[i2+1:]...)
					}
				}
			}

			moves = append(moves, gen...)

			if len(moves) == 0 {
				return false
			}

			p = moves[0].pos
			moves = moves[1:]
			visited.Add(p.noCost())
		}
	}

	for _, b := range bytes[1024:] {
		cord := strings.Split(b, ",")
		x := utils.ParseInt(cord[0])
		y := utils.ParseInt(cord[1])

		memory[y][x] = '#'

		if !hasPath() {
			fmt.Println(b)
			break
		}
	}
}
