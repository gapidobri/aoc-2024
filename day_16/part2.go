package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

type Part2 struct {
	tiles [][]rune
	start pos
	goal  pos
	moves []move
}

func part2() {
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

	p := Part2{
		tiles: tiles,
		goal:  end,
	}

	p.run(start)
}

func (p2 *Part2) run(start pos) {
	p2.moves = append(p2.moves, move{start, p2.heuristic(start), nil})

	var bestCost int
	bestTiles := mapset.NewSet[pos]()

	for len(p2.moves) > 0 {
		nextMove := p2.moves[0]
		nextI := 0
		for i, m := range p2.moves {
			if m.f < nextMove.f {
				nextMove = m
				nextI = i
			}
		}

		if p2.goalCheck(nextMove.pos) {
			fmt.Println(nextMove.pos.cost)
			if bestCost == 0 {
				bestCost = nextMove.pos.cost
			}
			if nextMove.pos.cost < bestCost {
				break
			}
			m := nextMove
			for {
				bestTiles.Add(m.pos)
				if m.prev != nil {
					m = *m.prev
				}
			}
		}

		p2.moves = append(p2.moves[:nextI], p2.moves[nextI+1:]...)

		generated := p2.generateMoves(&nextMove)

		for i, m := range p2.moves {
			for i2, m2 := range generated {
				if m.pos.noCost() != m2.pos.noCost() {
					continue
				}
				if m.f > m2.f {
					p2.moves = append(p2.moves[:i], p2.moves[i+1:]...)
				} else {
					generated = append(generated[:i2], generated[i2+1:]...)
				}
			}
		}

		p2.moves = append(p2.moves, generated...)
	}

	fmt.Println(bestTiles.Cardinality())
}

func (p2 *Part2) goalCheck(p pos) bool {
	return p.x == p2.goal.x && p.y == p2.goal.y
}

func (p2 *Part2) generateMoves(m *move) []move {
	var moves []move

	nextPos := m.pos.move()
	if p2.getTile(nextPos) != '#' {
		moves = append(moves, move{
			pos:  nextPos,
			f:    nextPos.cost + p2.heuristic(nextPos),
			prev: m,
		})
	}

	cw := m.pos.rotateCW().move()
	if p2.getTile(cw) != '#' {
		moves = append(moves, move{
			pos:  cw,
			f:    cw.cost + p2.heuristic(cw),
			prev: m,
		})
	}

	ccw := m.pos.rotateCCW().move()
	if p2.getTile(ccw) != '#' {
		moves = append(moves, move{
			pos:  ccw,
			f:    cw.cost + p2.heuristic(ccw),
			prev: m,
		})
	}

	return moves
}

func (p2 *Part2) heuristic(p pos) int {
	return p.distance(p2.goal)
}

func (p2 *Part2) getTile(p pos) rune {
	return p2.tiles[p.y][p.x]
}

func (p2 *Part2) print() {
	//fmt.Print("\033[H\033[2J")
	for _, row := range p2.tiles {
		fmt.Println(string(row))
	}
}

func (p2 *Part2) printMoves(m move) {
	tiles := make([][]rune, len(p2.tiles))
	for i, row := range p2.tiles {
		tiles[i] = make([]rune, len(row))
		copy(tiles[i], row)
	}
	for {
		tiles[m.pos.y][m.pos.x] = rotRuneMap[m.pos.rot]
		if m.prev == nil {
			break
		}
		m = *m.prev
	}
	for _, row := range tiles {
		fmt.Println(string(row))
	}
}
