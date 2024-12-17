package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

type Part1 struct {
	tiles [][]rune
	start pos
	goal  pos
	moves []move
}

func part1() {
	data := utils.ReadFile("day_16/input2.txt")

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

		generated := p1.generateMoves(nextMove.pos)

		for i, m := range p1.moves {
			for i2, m2 := range generated {
				if m.pos.noCost() != m2.pos.noCost() {
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

func (p1 *Part1) generateMoves(p pos) []move {
	var moves []move

	nextPos := p.move()
	if p1.getTile(nextPos) != '#' {
		moves = append(moves, move{
			pos: nextPos,
			f:   nextPos.cost + p1.heuristic(nextPos),
		})
	}

	cw := p.rotateCW().move()
	if p1.getTile(cw) != '#' {
		moves = append(moves, move{
			pos: cw,
			f:   cw.cost + p1.heuristic(cw),
		})
	}

	ccw := p.rotateCCW().move()
	if p1.getTile(ccw) != '#' {
		moves = append(moves, move{
			pos: ccw,
			f:   cw.cost + p1.heuristic(ccw),
		})
	}

	return moves
}

func (p1 *Part1) heuristic(p pos) int {
	return p.distance(p1.goal)
}

func (p1 *Part1) getTile(p pos) rune {
	return p1.tiles[p.y][p.x]
}
