package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize"
	"math"
)

func pain1() {
	data := utils.ReadFile("day_13/input.txt")

	matches := r.FindAllStringSubmatch(data, -1)

	tokens := 0
	for _, match := range matches {
		ax := utils.ParseFloat(match[1])
		ay := utils.ParseFloat(match[2])
		bx := utils.ParseFloat(match[3])
		by := utils.ParseFloat(match[4])
		px := utils.ParseFloat(match[5])
		py := utils.ParseFloat(match[6])

		buttons := mat.NewDense(2, 2, []float64{
			ax, bx,
			ay, by,
		})
		prize := mat.NewVecDense(2, []float64{
			px, py,
		})

		var res mat.VecDense
		err := res.SolveVec(buttons, prize)
		if err != nil {
			fmt.Println(err)
			continue
		}

		a := res.AtVec(0)
		b := res.AtVec(1)

		if utils.IsInt(a) && utils.IsInt(b) {
			tokens += int(math.Round(a))*3 + int(math.Round(b))
		}
	}

	fmt.Println(tokens)
}

func pain2() {
	data := utils.ReadFile("day_13/input.txt")

	matches := r.FindAllStringSubmatch(data, -1)

	tokens := 0
	for _, match := range matches {
		ax := utils.ParseFloat(match[1])
		ay := utils.ParseFloat(match[2])
		bx := utils.ParseFloat(match[3])
		by := utils.ParseFloat(match[4])
		px := utils.ParseFloat(match[5])
		py := utils.ParseFloat(match[6])

		problem := optimize.Problem{
			Func: func(x []float64) float64 {
				a := x[0]
				b := x[1]

				const1 := ax*a + bx*b - px
				const2 := ay*a + by*b - py

				penalty := 1e6 * (const1*const1 + const2*const2)

				return a + penalty
			},
		}

		initX := []float64{0, 0}

		settings := optimize.Settings{}
		result, err := optimize.Minimize(problem, initX, &settings, nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		a := result.X[0]
		b := result.X[1]

		if utils.IsInt(a) && utils.IsInt(b) {
			fmt.Println(a, b)
			tokens += int(math.Round(a))*3 + int(math.Round(b))
		}
	}

	fmt.Println(tokens)
}
