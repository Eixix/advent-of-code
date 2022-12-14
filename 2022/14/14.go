package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	first()
}

func stringToCoordinates(coordinateString string) (int, int) {
	coordinates := strings.Split(coordinateString, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])

	return x, y
}

func coordinatesToString(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func parseMap(filename string) (map[string]string, int) {
	file, _ := os.Open(filename)
	defer file.Close()

	caveMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	floor := 0
	for scanner.Scan() {
		combinedCoordinates := strings.Split(scanner.Text(), " -> ")
		stone := [][]int{}

		for _, v := range combinedCoordinates {
			x, y := stringToCoordinates(v)
			stone = append(stone, []int{x, y})

			if y > floor {
				floor = y
			}
		}
		for i := 0; i < len(stone); i++ {
			x1 := stone[i][0]
			y1 := stone[i][1]

			if i+1 == len(stone) {
				caveMap[coordinatesToString(x1, y1)] = "#"
				break
			}

			x2 := stone[i+1][0]
			y2 := stone[i+1][1]

			for x1 < x2 {
				caveMap[coordinatesToString(x1, y1)] = "#"
				x1++
			}
			for x1 > x2 {
				caveMap[coordinatesToString(x1, y1)] = "#"
				x1--
			}

			for y1 < y2 {
				caveMap[coordinatesToString(x1, y1)] = "#"
				y1++
			}
			for y1 > y2 {
				caveMap[coordinatesToString(x1, y1)] = "#"
				y1--
			}
		}
	}

	return caveMap, floor + 2
}

func dropSand(dropLocation string, caveMap map[string]string, floor int, isStart bool) (bool, map[string]string) {

	x, y := stringToCoordinates(dropLocation)
	straightDown := coordinatesToString(x, y+1)
	leftDiagonal := coordinatesToString(x-1, y+1)
	rightDiagonal := coordinatesToString(x+1, y+1)

	_, straightDownIsBlocked := caveMap[straightDown]
	_, leftDiagonalIsBlocked := caveMap[leftDiagonal]
	_, rightDiagonalIsBlocked := caveMap[rightDiagonal]

	if !straightDownIsBlocked && y+1 < floor {
		return dropSand(straightDown, caveMap, floor, false)
	} else if !leftDiagonalIsBlocked && y+1 < floor {
		return dropSand(leftDiagonal, caveMap, floor, false)
	} else if !rightDiagonalIsBlocked && y+1 < floor {
		return dropSand(rightDiagonal, caveMap, floor, false)
	} else if x == 500 && y == 0 {
		return true, caveMap
	}

	caveMap[dropLocation] = "+"
	return false, caveMap
}

func first() {
	caveMap, floor := parseMap("challenge.txt")
	clogged := false
	resting := 0

	for !clogged {
		clogged, caveMap = dropSand("500,0", caveMap, floor, true)
		resting++
	}
	fmt.Println(resting)
}
