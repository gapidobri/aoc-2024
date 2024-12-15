package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

const (
	up    = '^'
	right = '>'
	down  = 'v'
	left  = '<'
)

type pos struct {
	x, y int
}

func (p1 pos) add(p2 pos) pos {
	return pos{p1.x + p2.x, p1.y + p2.y}
}

var nextMove = map[rune]pos{
	up:    {0, -1},
	right: {1, 0},
	down:  {0, 1},
	left:  {-1, 0},
}

func main() {
	part1()
	part2()
}

func part1() {
	data := utils.ReadFile("day_15/input.txt")

	parts := strings.Split(data, "\n\n")

	var robot pos
	var m [][]rune
	for y, row := range strings.Split(parts[0], "\n") {
		m = append(m, []rune{})
		for x, char := range row {
			m[y] = append(m[y], char)
			if char == '@' {
				robot = pos{x, y}
			}
		}
	}

	instr := strings.Replace(parts[1], "\n", "", -1)
	for _, dir := range instr {
		if move1(m, robot, dir) {
			next := nextMove[dir]
			robot.x += next.x
			robot.y += next.y
		}
	}

	sum := 0
	for y, row := range m {
		for x, char := range row {
			if char != 'O' {
				continue
			}
			sum += y*100 + x
		}
	}

	fmt.Println(sum)
}

func move1(m [][]rune, p pos, dir rune) bool {
	next := p.add(nextMove[dir])
	switch m[next.y][next.x] {
	case '#':
		return false
	case 'O':
		if !move1(m, next, dir) {
			return false
		}
	}
	m[next.y][next.x] = m[p.y][p.x]
	m[p.y][p.x] = '.'
	return true
}

var repl = map[rune][]rune{
	'#': {'#', '#'},
	'O': {'[', ']'},
	'.': {'.', '.'},
	'@': {'@', '.'},
}

func part2() {
	data := utils.ReadFile("day_15/input.txt")

	parts := strings.Split(data, "\n\n")

	var robot pos
	var m [][]rune
	for y, row := range strings.Split(parts[0], "\n") {
		m = append(m, []rune{})
		for x, char := range row {
			m[y] = append(m[y], repl[char]...)
			if char == '@' {
				robot = pos{x * 2, y}
			}
		}
	}

	instr := strings.Replace(parts[1], "\n", "", -1)
	for _, dir := range instr {
		if canMove(m, robot, dir, true) {
			move2(m, robot, dir, true)
			next := nextMove[dir]
			robot.x += next.x
			robot.y += next.y
		}
	}

	sum := 0
	for y, row := range m {
		for x, char := range row {
			if char != '[' {
				continue
			}
			sum += y*100 + x
		}
	}

	fmt.Println(sum)
}

func canMove(m [][]rune, p pos, dir rune, check bool) bool {
	nextPos := p.add(nextMove[dir])
	curr := m[p.y][p.x]
	next := m[nextPos.y][nextPos.x]

	if next == '#' {
		return false
	}

	// check if the second part of the box can move
	if check && (dir == up || dir == down) &&
		(curr == '[' && !canMove(m, p.add(pos{1, 0}), dir, false) ||
			curr == ']' && !canMove(m, p.add(pos{-1, 0}), dir, false)) {
		return false
	}

	if (next == '[' || next == ']') && !canMove(m, nextPos, dir, true) {
		return false
	}

	return true
}

func move2(m [][]rune, p pos, dir rune, move bool) {
	nextPos := p.add(nextMove[dir])
	curr := m[p.y][p.x]
	next := m[nextPos.y][nextPos.x]

	if next == '[' || next == ']' {
		move2(m, nextPos, dir, true)
	}

	// move second part of the box
	if move && (dir == up || dir == down) {
		if curr == '[' {
			move2(m, p.add(pos{1, 0}), dir, false)
		} else if curr == ']' {
			move2(m, p.add(pos{-1, 0}), dir, false)
		}
	}

	m[nextPos.y][nextPos.x] = m[p.y][p.x]
	m[p.y][p.x] = '.'
}
