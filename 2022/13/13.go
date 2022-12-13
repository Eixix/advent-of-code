package main

import (
	"bufio"
	"os"
)

func main() {
	first()
}

func parseCharacters(line string, smaller bool, writeChannel chan rune, readChannel chan rune, returnChannel chan bool) bool {

	inRightOrder := true

	inList := 0
	// json.Unmarshal(line)
	for _, v := range line {
		if v == '[' {
			inList++
		} else if v == ']' {
			inList--
		} else if v == ',' {
			continue
		} else {

		}

		writeChannel <- v
		if smaller {
			if <-readChannel < v {
				return false
			} else if <-readChannel > v {
				return true
			}
		} else {
			if <-readChannel > v {
				return false
			} else if <-readChannel < v {
				return true
			}
		}
	}
	return inRightOrder
}

func first() {
	file, _ := os.Open("small.txt")
	defer file.Close()

	lines := []string{}

	validPairs := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	channel1 := make(chan rune)
	channel2 := make(chan rune)

	returnChannel1 := make(chan bool)
	returnChannel2 := make(chan bool)

	for i := 0; i < len(lines); i += 3 {
		firstLine := lines[i]
		secondLine := lines[i+1]

		// Start at second character as the first always opens up a list
		go parseCharacters(firstLine, true, channel1, channel2, returnChannel1)
		go parseCharacters(secondLine, false, channel2, channel1, returnChannel2)

		if <-returnChannel1 && <-returnChannel2 {
			validPairs++
		}
	}
}
