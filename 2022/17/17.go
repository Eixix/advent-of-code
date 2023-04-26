package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	first()
}

type Rock struct {
	coordinates []string
}

func stringToCoordinates(coordinateString string) (int, int) {
	coordinateStringSlice := strings.Fields(coordinateString)
	x, _ := strconv.Atoi(coordinateStringSlice[0])
	y, _ := strconv.Atoi(coordinateStringSlice[1])
	return x, y
}

func coordinatesToString(x, y int) string {
	return strconv.Itoa(x) + strconv.Itoa(y)
}

func parseFile(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return scanner.Text()
}

func dropRock(rock Rock, tetrisMap map[string]string, jet rune) map[string]string {

}

func tick(totalTicks int, tetrisMap map[string]string, jets string) {
	minusRock := Rock{[]string{"0 0", "1 0", "2 0", "3 0"}}
	plusRock := Rock{[]string{"0 1", "1 1", "2 1", "1 0", "1 2"}}
	lRock := Rock{[]string{"0 0", "1 0", "2 0", "2 1", "2 2"}}
	columnRock := Rock{[]string{"0 0", "0 1", "0 2", "0 3"}}
	squareRock := Rock{[]string{"0 0", "0 1", "1 0", "2 0"}}

	for i := 1; i <= totalTicks; i++ {
		jet := rune(jets[(i-1)%len(jets)])
		rockType := i % 5
		if rockType == 0 {
			tetrisMap = dropRock(minusRock, tetrisMap, jet)
		} else if rockType == 1 {
			tetrisMap = dropRock(plusRock, tetrisMap, jet)
		} else if rockType == 2 {
			tetrisMap = dropRock(lRock, tetrisMap, jet)
		} else if rockType == 3 {
			tetrisMap = dropRock(columnRock, tetrisMap, jet)
		} else if rockType == 4 {
			tetrisMap = dropRock(squareRock, tetrisMap, jet)
		}
	}
}

func first() {
	jets := parseFile("first.txt")
	tick(2022, make(map[string]string), jets)
}
