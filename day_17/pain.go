package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gapidobri/aoc-2024/utils"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func part2pain() {
	data := utils.ReadFile("day_17/input.txt")

	matches := regexp.
		MustCompile("[0-9,]+").
		FindAllString(data, -1)

	a := []string{"x"}
	b := []string{matches[1]}
	c := []string{matches[2]}

	var instr []int
	for _, i := range strings.Split(matches[3], ",") {
		instr = append(instr, utils.ParseInt(i))
	}

	var output [][]string

outer:
	for i := 0; i < len(instr); {
		literal := instr[i+1]
		combo := map[int][]string{0: {"0"}, 1: {"1"}, 2: {"2"}, 3: {"3"}, 4: a, 5: b, 6: c}[literal]

		switch instr[i] {
		case adv:
			var o []string
			if ok, combVal := isConst(combo); ok {
				pow := utils.IntPow(2, combVal)
				if ok2, aVal := isConst(a); ok2 {
					o = append(o, strconv.Itoa(aVal/pow))
				} else {
					o = append(o, a...)
					o = append(o, strconv.Itoa(pow), "/")
				}
			} else {
				o = append(o, "2")
				o = append(o, combo...)
				o = append(o, "^", "/")
			}
			a = o
		case bxl:
			if ok, val := isConst(b); ok {
				b = []string{strconv.Itoa(val ^ literal)}
			} else {
				b = append(b, strconv.Itoa(literal), "|")
			}
		case bst:
			if ok, val := isConst(combo); ok {
				b = []string{strconv.Itoa(val % 8)}
			} else {
				var o []string
				o = append(o, combo...)
				o = append(o, "8", "%")
				b = o
			}
		case jnz:
			i = literal
			continue
		case bxc:
			ok, bVal := isConst(b)
			ok2, cVal := isConst(c)
			if ok && ok2 {
				b = []string{strconv.Itoa(bVal ^ cVal)}
			} else {
				b = append(b, c...)
				b = append(b, "|")
			}
		case out:
			var o []string
			o = append(o, combo...)
			o = append(o, "8")
			o = append(o, "%")
			output = append(output, o)
			if len(output) >= len(instr) {
				break outer
			}
		case bdv:
			var o []string
			if ok, combVal := isConst(combo); ok {
				pow := utils.IntPow(2, combVal)
				if ok2, aVal := isConst(a); ok2 {
					o = append(o, strconv.Itoa(aVal/pow))
				} else {
					o = append(o, a...)
					o = append(o, strconv.Itoa(pow), "/")
				}
			} else {
				o = append(o, "2")
				o = append(o, combo...)
				o = append(o, "^", "/")
			}
			b = o
		case cdv:
			var o []string
			if ok, combVal := isConst(combo); ok {
				pow := utils.IntPow(2, combVal)
				if ok2, aVal := isConst(a); ok2 {
					o = append(o, strconv.Itoa(aVal/pow))
				} else {
					o = append(o, a...)
					o = append(o, strconv.Itoa(pow), "/")
				}
			} else {
				o = append(o, "2")
				o = append(o, combo...)
				o = append(o, "^", "/")
			}
			c = o
		}

		i += 2
	}

	//for i, expr := range output {
	//	var vals []string
	//	for i2, s := range expr {
	//		switch s {
	//		case "/":
	//			fallthrough
	//		case "%":
	//			fallthrough
	//		case "^":
	//			fallthrough
	//		case "|":
	//			continue
	//		default:
	//			vals = append(vals, s)
	//		}
	//	}
	//}

	for _, o := range output {
		fmt.Println(o)
		fmt.Println(eval(o, 30878003))
	}

	return

	inters := mapset.NewSet[int]()

	for i, r := range instr {
		//fmt.Println(r)
		//fmt.Println(eval(output[i], 117440))
		set := mapset.NewSet[int]()
		for t := 0; t <= 0; t++ {
			res := float64(r)
			expr := output[i]

			ops := utils.NewStack[string]()
			vals := utils.NewStack[int]()
			for _, v := range expr {
				if v == "x" {
					continue
				}
				num, err := strconv.Atoi(v)
				if err == nil {
					vals.Push(num)
				} else {
					ops.Push(v)
				}
			}

			for ops.Len() > 0 {
				switch ops.Pop() {
				case "^":
					vals.Pop()
					res = rev2Pow(res)
				case "/":
					res = revDiv(float64(vals.Pop()), res)
				case "%":
					res = revMod(float64(vals.Pop()), res, 28) // todo 1835
				case "|":
					res = revXOR(float64(vals.Pop()), res)
				}
				//fmt.Println(res)
			}

			set.Add(int(res))
		}

		if inters.Cardinality() == 0 {
			inters = inters.Union(set)
		} else {
			inters = inters.Intersect(set)
			//fmt.Println(inters.Cardinality(), inters.ToSlice())
		}
		//fmt.Println()
	}

	//fmt.Println(inters.Cardinality())
	//fmt.Println(inters.ToSlice())
}

func isConst(reg []string) (bool, int) {
	if len(reg) == 1 && reg[0] != "x" {
		return true, utils.ParseInt(reg[0])
	}
	return false, 0
}

func eval(expr []string, x int) int {
	//fmt.Println(expr)
	s := utils.Stack[int]{}
	for _, v := range expr {
		switch v {
		case "x":
			s.Push(x)
		case "^":
			right := s.Pop()
			left := s.Pop()
			s.Push(utils.IntPow(left, right))
		case "/":
			right := s.Pop()
			left := s.Pop()
			s.Push(left / right)
		case "%":
			right := s.Pop()
			left := s.Pop()
			s.Push(left % right)
		case "|":
			right := s.Pop()
			left := s.Pop()
			s.Push(left ^ right)
		default:
			s.Push(utils.ParseInt(v))
		}
		fmt.Println(v, s.String())
	}
	return s.Pop()
}

func rev2Pow(res float64) float64 {
	return math.Log2(res)
}

func revDiv(b, res float64) float64 {
	return b * res
}

func revMod(b, res float64, t int) float64 {
	return float64(t)*b + res
}

func revXOR(b, res float64) float64 {
	return float64(int(b) ^ int(res))
}
