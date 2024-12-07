package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"slices"
	"strings"
	"time"
)

type path struct {
	start, end int
}

func overengineeredPart1() {
	start := time.Now()

	data := utils.ReadFile("day_06/input.txt")

	var guard pos
	var guardDir int
	verObst := map[int][]int{}
	horObst := map[int][]int{}

	rows := strings.Split(data, "\n")
	for y, row := range rows {
		for x, char := range row {
			switch char {
			case '#':
				verObst[x] = append(verObst[x], y)
				horObst[y] = append(horObst[y], x)
				break
			case '^':
				guardDir = up
				guard = pos{x, y}
				break
			case '>':
				guardDir = right
				guard = pos{x, y}
				break
			case 'v':
				guardDir = down
				guard = pos{x, y}
				break
			case '<':
				guardDir = left
				guard = pos{x, y}
				break
			}
		}
	}

	posCount := 1
	//verPaths := map[int][]path{}
	//horPaths := map[int][]path{}

outer:
	for {
		fmt.Println(posCount)
		switch guardDir {
		case up:
			for _, y := range slices.Backward(verObst[guard.x]) {
				if y < guard.y {
					posCount += guard.y - y - 1
					//verPaths[guard.x] = append(verPaths[guard.x], path{guard.y, y + 1})
					guard.y = y + 1
					guardDir = right
					continue outer
				}
			}
			posCount += guard.y
		case right:
			for _, x := range horObst[guard.y] {
				if x > guard.x {
					posCount += x - guard.x - 1
					//horPaths[guard.y] = append(horPaths[guard.y], path{guard.x, x - 1})
					guard.x = x - 1
					guardDir = down
					continue outer
				}
			}
			posCount += len(rows[0]) - guard.x - 1
		case down:
			for _, y := range verObst[guard.x] {
				if y > guard.y {
					posCount += y - guard.y - 1
					//verPaths[guard.x] = append(verPaths[guard.x], path{guard.y, y - 1})
					guard.y = y - 1
					guardDir = left
					continue outer
				}
			}
			posCount += len(rows) - guard.y - 1
		case left:
			for _, x := range slices.Backward(horObst[guard.y]) {
				if x < guard.x {
					posCount += guard.x - x - 1
					//horPaths[guard.y] = append(horPaths[guard.y], path{guard.x, x + 1})
					guard.x = x + 1
					guardDir = up
					continue outer
				}
			}
			posCount += guard.x
		}
		break
	}

	fmt.Println(posCount)

	fmt.Println("Time elapsed: ", time.Since(start))
}
