package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//first()
	second()
}

type Sensor struct {
	sensorX           int
	sensorY           int
	beaconX           int
	beaconY           int
	manhattanDistance int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseFile(filename string) ([]Sensor, int, int) {
	file, _ := os.Open(filename)
	defer file.Close()

	sensors := []Sensor{}
	var minX int
	var maxX int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		sensorX, _ := strconv.Atoi(s[2][2 : len(s[2])-1])
		sensorY, _ := strconv.Atoi(s[3][2 : len(s[3])-1])
		beaconX, _ := strconv.Atoi(s[8][2 : len(s[8])-1])
		beaconY, _ := strconv.Atoi(s[9][2:len(s[9])])

		manhattanDistance := abs(sensorX-beaconX) + abs(sensorY-beaconY)

		if beaconX-manhattanDistance < minX {
			minX = beaconX - manhattanDistance
		}
		if maxX < beaconX+manhattanDistance {
			maxX = beaconX + manhattanDistance
		}

		sensors = append(sensors, Sensor{sensorX, sensorY, beaconX, beaconY, manhattanDistance})
	}

	return sensors, minX, maxX
}

func first() {
	sensors, minX, maxX := parseFile("challenge.txt")
	y := 2000000
	toReturn := 0

	for x := minX; x <= maxX; x++ {
		for _, sensor := range sensors {
			if abs(sensor.sensorX-x)+abs(sensor.sensorY-y) <= sensor.manhattanDistance {
				for _, innerSensor := range sensors {
					if x == innerSensor.beaconX && y == innerSensor.beaconY {
						toReturn--
						break
					}
				}
				toReturn++
				break
			}
		}
	}

	fmt.Println(toReturn)
}

func parallelSearch(channel chan int, max int, sensor Sensor, sensors []Sensor) {
	for x := sensor.sensorX - sensor.manhattanDistance - 1; x <= sensor.sensorX; x++ {
		y := sensor.sensorY - (x - (sensor.sensorX - sensor.manhattanDistance - 1))
		notInAnySensor := true
		for _, innerSensor := range sensors {
			if (abs(innerSensor.sensorX-x)+abs(innerSensor.sensorY-y) <= innerSensor.manhattanDistance) && x >= 0 && x <= max && y >= 0 && y <= max {
				notInAnySensor = false
			}
		}
		if notInAnySensor && x >= 0 && x <= max && y >= 0 && y <= max {
			channel <- x*4000000 + y
		}
		y = sensor.sensorY + (x - (sensor.sensorX - sensor.manhattanDistance - 1))
		for _, innerSensor := range sensors {
			if (abs(innerSensor.sensorX-x)+abs(innerSensor.sensorY-y) <= innerSensor.manhattanDistance) && x >= 0 && x <= max && y >= 0 && y <= max {
				notInAnySensor = false
			}
		}
		if notInAnySensor && x >= 0 && x <= max && y >= 0 && y <= max {
			channel <- x*4000000 + y
		}
	}
	for x := sensor.sensorX; x <= sensor.sensorX+sensor.manhattanDistance+1; x++ {
		y := (sensor.sensorY + sensor.manhattanDistance + 1) - (x - sensor.sensorX)
		notInAnySensor := true
		for _, innerSensor := range sensors {
			if (abs(innerSensor.sensorX-x)+abs(innerSensor.sensorY-y) <= innerSensor.manhattanDistance) && x >= 0 && x <= max && y >= 0 && y <= max {
				notInAnySensor = false
			}
		}
		if notInAnySensor && x >= 0 && x <= max && y >= 0 && y <= max {
			channel <- x*4000000 + y
		}
		y = (sensor.sensorY - sensor.manhattanDistance - 1) + (x - sensor.sensorX)
		for _, innerSensor := range sensors {
			if (abs(innerSensor.sensorX-x)+abs(innerSensor.sensorY-y) <= innerSensor.manhattanDistance) && x >= 0 && x <= max && y >= 0 && y <= max {
				notInAnySensor = false
			}
		}
		if notInAnySensor && x >= 0 && x <= max && y >= 0 && y <= max {
			channel <- x*4000000 + y
		}
	}
}

func second() {
	sensors, _, _ := parseFile("challenge.txt")
	max := 4000000

	resultChannel := make(chan int)

	for _, sensor := range sensors {
		go parallelSearch(resultChannel, max, sensor, sensors)
	}

	fmt.Println(<-resultChannel)
}
