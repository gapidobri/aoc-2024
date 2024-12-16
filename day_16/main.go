package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

func main() {
	part1()
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

func (p pos) Distance(other pos) int {
	return utils.IntAbs(p.x-other.x) + utils.IntAbs(p.y-other.y)
}

func (p pos) Move() pos {
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

func (p pos) RotateCW() pos {
	return pos{p.x, p.y, (p.rot + 1) % 4, p.cost + 1000}
}

func (p pos) RotateCCW() pos {
	return pos{p.x, p.y, (4 + p.rot - 1) % 4, p.cost + 1000}
}

func (p pos) NoCost() pos {
	return pos{p.x, p.y, p.rot, 0}
}

type move struct {
	pos  pos
	f    int
	prev *move
}

type Part1 struct {
	tiles [][]rune
	start pos
	goal  pos
	moves []move
}

func part1() {
	data := utils.ReadFile("day_16/input.txt")

	var tiles [][]rune

	var start, end pos
	for y, row := range strings.Split(data, "\n") {
		tiles = append(tiles, []rune(row))
		for x, char := range row {
			if char == 'S' {
				start = pos{x, y, right, 0}
			} else if char == 'E' {
				end = pos{x, y, 0, 0}
			}
		}
	}

	p := Part1{
		tiles: tiles,
		goal:  end,
	}

	p.run(start)
}

func (p1 *Part1) run(start pos) {
	p1.moves = append(p1.moves, move{start, p1.heuristic(start), nil})

	for len(p1.moves) > 0 {
		nextMove := p1.moves[0]
		nextI := 0
		for i, m := range p1.moves {
			if m.f < nextMove.f {
				nextMove = m
				nextI = i
			}
		}

		if p1.goalCheck(nextMove.pos) {
			fmt.Println(nextMove.pos.cost)
			break
		}

		p1.moves = append(p1.moves[:nextI], p1.moves[nextI+1:]...)
		//
		//dist := nextMove.pos.Distance(p1.goal)
		//if dist < p1.minDist {
		//	p1.minDist = dist
		//	fmt.Println(p1.minDist)
		//}

		//p1.tiles[nextMove.pos.y][nextMove.pos.x] = rotRuneMap[nextMove.pos.rot]
		//p1.print()
		//time.Sleep(10 * time.Millisecond)

		generated := p1.generateMoves(&nextMove)

		for i, m := range p1.moves {
			for i2, m2 := range generated {
				if m.pos.NoCost() != m2.pos.NoCost() {
					continue
				}
				if m.f > m2.f {
					p1.moves = append(p1.moves[:i], p1.moves[i+1:]...)
				} else {
					generated = append(generated[:i2], generated[i2+1:]...)
				}
			}
		}

		p1.moves = append(p1.moves, generated...)
	}
}

func (p1 *Part1) goalCheck(p pos) bool {
	return p.x == p1.goal.x && p.y == p1.goal.y
}

func (p1 *Part1) generateMoves(m *move) []move {
	var moves []move

	nextPos := m.pos.Move()
	if p1.getTile(nextPos) != '#' {
		moves = append(moves, move{
			pos:  nextPos,
			f:    nextPos.cost + p1.heuristic(nextPos),
			prev: m,
		})
	}

	cw := m.pos.RotateCW().Move()
	if p1.getTile(cw) != '#' {
		moves = append(moves, move{
			pos:  cw,
			f:    cw.cost + p1.heuristic(cw),
			prev: m,
		})
	}

	ccw := m.pos.RotateCCW().Move()
	if p1.getTile(ccw) != '#' {
		moves = append(moves, move{
			pos:  ccw,
			f:    cw.cost + p1.heuristic(ccw),
			prev: m,
		})
	}

	return moves
}

func (p1 *Part1) heuristic(p pos) int {
	return p.Distance(p1.goal)
}

func (p1 *Part1) getTile(p pos) rune {
	return p1.tiles[p.y][p.x]
}

func (p1 *Part1) print() {
	//fmt.Print("\033[H\033[2J")
	for _, row := range p1.tiles {
		fmt.Println(string(row))
	}
}
