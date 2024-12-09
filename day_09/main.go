package main

import (
	"fmt"
	"github.com/gapidobri/aoc-2024/utils"
)

func getDiskMap() []rune {
	data := utils.ReadFile("day_09/input.txt")

	var disk []rune
	for i, char := range data {
		if i%2 == 0 {
			id := i / 2
			for range char - '0' {
				disk = append(disk, rune(id)+'0')
			}
		} else {
			for range char - '0' {
				disk = append(disk, '.')
			}
		}
	}

	return disk
}

func main() {
	part1()
	part2()
}

func part1() {
	disk := getDiskMap()

	freeI := 0
	lastI := len(disk) - 1
	for {
		for disk[freeI] != '.' {
			freeI++
		}
		for disk[lastI] == '.' {
			lastI--
		}
		if freeI >= lastI {
			break
		}
		disk[freeI] = disk[lastI]
		disk[lastI] = '.'
	}

	sum := 0
	for i, char := range disk {
		if char == '.' {
			break
		}
		sum += i * int(char-'0')
	}

	fmt.Println(sum)
}

func part2() {
	disk := getDiskMap()

	startI := len(disk) - 1

outer:
	for {
		// skip empty space
		for disk[startI] == '.' {
			startI--
		}
		lastId := disk[startI]
		endI := startI
		// get to the start of the file
		for disk[startI-1] == lastId {
			startI--
			if startI == 0 {
				break outer
			}
		}
		fileSize := endI - startI + 1

		freeStartI := 0
	findFree:
		for {
			for disk[freeStartI] != '.' {
				freeStartI++
			}

			freeEndI := freeStartI
			if freeEndI >= startI {
				startI--
				continue outer
			}
			for freeEndI-freeStartI+1 < fileSize {
				freeEndI++
				if disk[freeEndI] != '.' {
					// not enough space to move file
					freeStartI = freeEndI
					continue findFree
				}
			}
			break
		}

		// move the file
		for i := range fileSize {
			disk[freeStartI+i] = disk[startI+i]
			disk[startI+i] = '.'
		}
	}

	sum := 0
	for i, char := range disk {
		if char == '.' {
			continue
		}
		sum += i * int(char-'0')
	}

	fmt.Println(sum)
}
