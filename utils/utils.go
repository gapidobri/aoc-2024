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

func IntAbs(num int) int {
	return int(math.Abs(float64(num)))
}
