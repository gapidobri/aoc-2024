package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	data := utils.ReadFile("day_04/input.txt")

	lines := strings.Split(data, "\n")

	sum := 0
	for _, line := range lines {
		for i := range len(line) - 3 {
			word := line[i : i+4]
			if word == "XMAS" || word == "SAMX" {
				sum++
			}
		}
	}

	for x := range len(lines[0]) {
		for y := range len(lines) - 3 {
			word := string(lines[y][x]) + string(lines[y+1][x]) + string(lines[y+2][x]) + string(lines[y+3][x])
			if word == "XMAS" || word == "SAMX" {
				sum++
			}
		}
	}

	for y := range len(lines) - 3 {
		for x := range len(lines[0]) - 3 {
			word := string(lines[y][x]) + string(lines[y+1][x+1]) + string(lines[y+2][x+2]) + string(lines[y+3][x+3])
			if word == "XMAS" || word == "SAMX" {
				sum++
			}
			word = string(lines[y][x+3]) + string(lines[y+1][x+2]) + string(lines[y+2][x+1]) + string(lines[y+3][x])
			if word == "XMAS" || word == "SAMX" {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func part2() {
	data := utils.ReadFile("day_04/input.txt")

	lines := strings.Split(data, "\n")

	sum := 0
	for y := range len(lines) - 2 {
		for x := range len(lines[0]) - 2 {
			word := string(lines[y][x]) + string(lines[y+1][x+1]) + string(lines[y+2][x+2])
			if !(word == "MAS" || word == "SAM") {
				continue
			}
			word = string(lines[y+2][x]) + string(lines[y+1][x+1]) + string(lines[y][x+2])
			if word == "MAS" || word == "SAM" {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
