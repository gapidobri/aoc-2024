package utils

import (
	"math"
	"os"
	"strconv"
)

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ParseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func ParseFloat(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func IntAbs(num int) int {
	return int(math.Abs(float64(num)))
}

func IsInt(num float64) bool {
	epsilon := 1e-6
	_, frac := math.Modf(math.Abs(num))
	return frac < epsilon || frac > 1.0-epsilon
}
