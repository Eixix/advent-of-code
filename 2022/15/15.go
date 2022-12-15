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
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func parseFile(filename string) map[string]string {
	file, _ := os.Open(filename)
	defer file.Close()

	beaconSensorMap := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		sensorX, _ := strconv.Atoi(s[2][2 : len(s[2])-1])
		sensorY, _ := strconv.Atoi(s[3][2 : len(s[3])-1])
		beaconX, _ := strconv.Atoi(s[8][2 : len(s[8])-1])
		beaconY, _ := strconv.Atoi(s[9][2:len(s[9])])
		beaconSensorMap[coordinatesToString(sensorX, sensorY)] = "S"
		beaconSensorMap[coordinatesToString(beaconX, beaconY)] = "B"

		manhattanDistance := abs(sensorX-beaconX) + abs(sensorY-beaconY)

		for i := 1; i < manhattanDistance; i++ {
			if sensorY == 2000000 && beaconSensorMap[coordinatesToString(sensorX+i, sensorY)] != "B" {
				beaconSensorMap[coordinatesToString(sensorX+i, sensorY)] = "#"
			} else if sensorY == 2000000 && beaconSensorMap[coordinatesToString(sensorX-i, sensorY)] != "B" {
				beaconSensorMap[coordinatesToString(sensorX-i, sensorY)] = "#"
			} else if sensorY+i == 2000000 && beaconSensorMap[coordinatesToString(sensorX, sensorY+i)] != "B" {
				beaconSensorMap[coordinatesToString(sensorX, sensorY+i)] = "#"
			} else if sensorY-i == 2000000 && beaconSensorMap[coordinatesToString(sensorX, sensorY-i)] != "B" {
				beaconSensorMap[coordinatesToString(sensorX, sensorY-i)] = "#"
			}
			for j := 1; j <= manhattanDistance-i; j++ {
				if sensorY+j == 2000000 && beaconSensorMap[coordinatesToString(sensorX+i, sensorY+j)] != "B" {
					beaconSensorMap[coordinatesToString(sensorX+i, sensorY+j)] = "#"
				} else if sensorY-j == 2000000 && beaconSensorMap[coordinatesToString(sensorX+i, sensorY-j)] != "B" {
					beaconSensorMap[coordinatesToString(sensorX+i, sensorY-j)] = "#"
				}
			}
			for j := 1; j <= manhattanDistance-i; j++ {
				if sensorY+j == 2000000 && beaconSensorMap[coordinatesToString(sensorX-i, sensorY+j)] != "B" {
					beaconSensorMap[coordinatesToString(sensorX-i, sensorY+j)] = "#"
				} else if sensorY-j == 2000000 && beaconSensorMap[coordinatesToString(sensorX-i, sensorY-j)] != "B" {
					beaconSensorMap[coordinatesToString(sensorX-i, sensorY-j)] = "#"
				}
			}
		}
	}

	return beaconSensorMap
}

func first() {
	beaconSensorMap := parseFile("challenge.txt")
	toReturn := 0

	for i, v := range beaconSensorMap {
		_, y := stringToCoordinates(i)
		if y == 2000000 && v == "#" {
			toReturn++
		}
	}

	fmt.Println(toReturn, beaconSensorMap)
}
