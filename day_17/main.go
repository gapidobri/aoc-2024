package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"regexp"
	"strconv"
	"strings"
)

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func main() {
	part1()
	part2()
}

func part1() {
	data := utils.ReadFile("day_17/input.txt")

	matches := regexp.
		MustCompile("[0-9,]+").
		FindAllString(data, -1)

	a := utils.ParseInt(matches[0])
	b := utils.ParseInt(matches[1])
	c := utils.ParseInt(matches[2])
	var instr []int
	for _, i := range strings.Split(matches[3], ",") {
		instr = append(instr, utils.ParseInt(i))
	}

	var output []string
	for i := 0; i < len(instr); {
		literal := instr[i+1]
		combo := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: a, 5: b, 6: c}[literal]

		switch instr[i] {
		case adv:
			a = a / utils.IntPow(2, combo)
		case bxl:
			b = b ^ literal
		case bst:
			b = combo % 8
		case jnz:
			if a != 0 {
				i = literal
				continue
			}
		case bxc:
			b = b ^ c
		case out:
			output = append(output, strconv.Itoa(combo%8))
		case bdv:
			b = a / utils.IntPow(2, combo)
		case cdv:
			c = a / utils.IntPow(2, combo)
		}

		i += 2
	}

	fmt.Println(strings.Join(output, ","))
}

var (
	input  string
	instrs []int
	b, c   int
)

func part2() {
	data := utils.ReadFile("day_17/input.txt")

	matches := regexp.
		MustCompile("[0-9,]+").
		FindAllString(data, -1)

	b = utils.ParseInt(matches[1])
	c = utils.ParseInt(matches[2])
	input = matches[3]

	for _, i := range strings.Split(input, ",") {
		instrs = append(instrs, utils.ParseInt(i))
	}

	fmt.Println(bfs(0, 0))
}

func run(a int, exp []int) bool {
	b := b
	c := c
	var output []int
	for i := 0; i < len(instrs); {
		literal := instrs[i+1]
		combo := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: a, 5: b, 6: c}[literal]

		switch instrs[i] {
		case adv:
			a = a / utils.IntPow(2, combo)
		case bxl:
			b = b ^ literal
		case bst:
			b = combo % 8
		case jnz:
			if a != 0 {
				i = literal
				continue
			}
		case bxc:
			b = b ^ c
		case out:
			output = append(output, combo%8)
		case bdv:
			b = a / utils.IntPow(2, combo)
		case cdv:
			c = a / utils.IntPow(2, combo)
		}

		i += 2
	}
	if len(output) < len(exp) {
		return false
	}
	for i, v := range exp {
		if output[len(output)-len(exp)+i] != v {
			return false
		}
	}
	return true
}

func bfs(a int, i int) int {
	if i > len(instrs) {
		return a
	}

	for v := range 8 {
		newA := a*8 + v
		if run(newA, instrs[len(instrs)-i:]) {
			if o := bfs(newA, i+1); o != -1 {
				return o
			}
		}
	}

	return -1
}
