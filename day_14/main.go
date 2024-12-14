package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"math"
	"regexp"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	data := utils.ReadFile("day_14/input.txt")

	r := regexp.MustCompile("p=(-?[0-9]+,-?[0-9]+) v=(-?[0-9]+,-?[0-9]+)")

	robots := r.FindAllStringSubmatch(data, -1)

	width := 101
	height := 103
	xCenter := (width - 1) / 2
	yCenter := (height - 1) / 2

	var quad [4]int
	for _, robot := range robots {
		p := strings.Split(robot[1], ",")
		px := utils.ParseInt(p[0])
		py := utils.ParseInt(p[1])

		v := strings.Split(robot[2], ",")
		vx := utils.ParseInt(v[0])
		vy := utils.ParseInt(v[1])

		newX := (width*100 + px + vx*100) % width
		newY := (height*100 + py + vy*100) % height

		if newX < xCenter {
			if newY < yCenter {
				quad[0]++
			} else if newY > yCenter {
				quad[2]++
			}
		} else if newX > xCenter {
			if newY < yCenter {
				quad[1]++
			} else if newY > yCenter {
				quad[3]++
			}
		}
	}

	fmt.Println(quad[0] * quad[1] * quad[2] * quad[3])
}

type pos struct {
	x, y int
}

func part2() {
	data := utils.ReadFile("day_14/input.txt")

	r := regexp.MustCompile("p=(-?[0-9]+,-?[0-9]+) v=(-?[0-9]+,-?[0-9]+)")

	matches := r.FindAllStringSubmatch(data, -1)

	var robotPos, robotVel []pos
	for _, match := range matches {
		p := strings.Split(match[1], ",")
		robotPos = append(robotPos, pos{utils.ParseInt(p[0]), utils.ParseInt(p[1])})

		v := strings.Split(match[2], ",")
		robotVel = append(robotVel, pos{utils.ParseInt(v[0]), utils.ParseInt(v[1])})
	}

	width := 101
	height := 103

	maxSeconds := 10000

	minDist := 100.0
	minSec := 0

	for seconds := 1; seconds <= maxSeconds; seconds++ {
		for i, vel := range robotVel {
			robotPos[i].x = (width + robotPos[i].x + vel.x) % width
			robotPos[i].y = (height + robotPos[i].y + vel.y) % height
		}

		var sum float64
		var count float64
		for _, robot1 := range robotPos {
			for _, robot2 := range robotPos {
				count++
				dx := robot1.x - robot2.x
				dy := robot1.y - robot2.y
				sum += math.Sqrt(float64(dx*dx + dy*dy))
			}
		}

		avgDist := sum / count
		if avgDist < minDist {
			minDist = avgDist
			minSec = seconds
		}
	}

	fmt.Println(minSec)
}
