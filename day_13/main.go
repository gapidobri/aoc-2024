package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"regexp"
)

var r = regexp.MustCompile("Button A: X\\+([0-9]+), Y\\+([0-9]+)\\nButton B: X\\+([0-9]+), Y\\+([0-9]+)\\nPrize: X=([0-9]+), Y=([0-9]+)")

func main() {
	part1()
	part2()
}

// i spent too much time on pain1() and pain2() in pain.go :sob::sob::sob:

func part1() {
	data := utils.ReadFile("day_13/input.txt")

	matches := r.FindAllStringSubmatch(data, -1)

	tokens := 0
	for _, match := range matches {
		xa := utils.ParseInt(match[1])
		ya := utils.ParseInt(match[2])
		xb := utils.ParseInt(match[3])
		yb := utils.ParseInt(match[4])
		xp := utils.ParseInt(match[5])
		yp := utils.ParseInt(match[6])

		b := (xa*yp - ya*xp) / (xa*yb - ya*xb)
		a := (yp - yb*b) / ya

		if xa*a+xb*b == xp && ya*a+yb*b == yp {
			tokens += a*3 + b
		}
	}

	fmt.Println(tokens)
}

func part2() {
	data := utils.ReadFile("day_13/input.txt")

	matches := r.FindAllStringSubmatch(data, -1)

	tokens := 0
	for _, match := range matches {
		xa := utils.ParseInt(match[1])
		ya := utils.ParseInt(match[2])
		xb := utils.ParseInt(match[3])
		yb := utils.ParseInt(match[4])
		xp := 10000000000000 + utils.ParseInt(match[5])
		yp := 10000000000000 + utils.ParseInt(match[6])

		b := (xa*yp - ya*xp) / (xa*yb - ya*xb)
		a := (yp - yb*b) / ya

		if xa*a+xb*b == xp && ya*a+yb*b == yp {
			tokens += a*3 + b
		}
	}

	fmt.Println(tokens)
}
