package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

type pos struct {
	x, y int
}

func (p pos) add(other pos) pos {
	return pos{p.x + other.x, p.y + other.y}
}

var nextPositions = []pos{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

var nextCheatPositions = []pos{
	{2, 0},
	{-2, 0},
	{0, 2},
	{0, -2},
}

func main() {
	part1()
	part2()
}

func getMap() (start pos, end pos, times map[pos]int) {
	data := utils.ReadFile("day_20/input.txt")

	var track [][]rune
	for y, row := range strings.Split(data, "\n") {
		track = append(track, []rune{})
		for x, char := range row {
			track[y] = append(track[y], char)
			switch char {
			case 'S':
				start = pos{x, y}
			case 'E':
				end = pos{x, y}
			}
		}
	}

	times = map[pos]int{}

	p := start
	prev := pos{0, 0}
	time := 0
	for {
		if p == end {
			break
		}
		for _, next := range nextPositions {
			newPos := p.add(next)
			if newPos != prev && track[newPos.y][newPos.x] != '#' {
				prev = p
				p = newPos
				break
			}
		}
		time++
		times[p] = time
	}

	return
}

func part1() {
	start, end, times := getMap()

	p := start
	time := 0
	count := 0
	for {
		if p == end {
			break
		}
		for _, next := range nextCheatPositions {
			newPos := p.add(next)
			if times[newPos] > time+2 {
				save := times[newPos] - time - 2
				if save >= 100 {
					count++
				}
			}
		}
		for _, next := range nextPositions {
			newPos := p.add(next)
			if times[newPos] == time+1 {
				p = newPos
				break
			}
		}
		time++
	}

	fmt.Println(count)
}

func part2() {
	start, end, times := getMap()

	p := start
	time := 0
	count := 0
	for {
		if p == end {
			break
		}
		for y := -20; y <= 20; y++ {
			for x := -20; x <= 20; x++ {
				dist := utils.IntAbs(x) + utils.IntAbs(y)
				if dist > 20 {
					continue
				}
				newPos := p.add(pos{x, y})
				save := times[newPos] - time - dist
				if save >= 100 {
					count++
				}
			}
		}
		for _, next := range nextPositions {
			newPos := p.add(next)
			if times[newPos] == time+1 {
				p = newPos
				break
			}
		}
		time++
	}

	fmt.Println(count)
}
