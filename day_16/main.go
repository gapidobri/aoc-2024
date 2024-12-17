package main

import (
	"github.com/gapidobri/aoc-2024/utils"
)

func main() {
	part1()
	//part2()
}

const (
	up = iota
	right
	down
	left
)

var rotRuneMap = map[int]rune{
	up:    '^',
	right: '>',
	down:  'v',
	left:  '<',
}

type pos struct {
	x, y int
	rot  int
	cost int
}

func (p pos) distance(other pos) int {
	return utils.IntAbs(p.x-other.x) + utils.IntAbs(p.y-other.y)
}

func (p pos) move() pos {
	switch p.rot {
	case up:
		return pos{p.x, p.y - 1, p.rot, p.cost + 1}
	case right:
		return pos{p.x + 1, p.y, p.rot, p.cost + 1}
	case down:
		return pos{p.x, p.y + 1, p.rot, p.cost + 1}
	case left:
		return pos{p.x - 1, p.y, p.rot, p.cost + 1}
	}
	return pos{}
}

func (p pos) rotateCW() pos {
	return pos{p.x, p.y, (p.rot + 1) % 4, p.cost + 1000}
}

func (p pos) rotateCCW() pos {
	return pos{p.x, p.y, (4 + p.rot - 1) % 4, p.cost + 1000}
}

func (p pos) noCost() pos {
	return pos{p.x, p.y, p.rot, 0}
}

func (p pos) pos() pos {
	return pos{x: p.x, y: p.y}
}

type move struct {
	pos  pos
	f    int
	prev *move
}
