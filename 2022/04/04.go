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
	second()
}

func first() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	toReturn := 0

	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, ",")

		firstPartLower, _ := strconv.Atoi(strings.Split(s[0], "-")[0])
		firstPartHigher, _ := strconv.Atoi(strings.Split(s[0], "-")[1])
		secondPartLower, _ := strconv.Atoi(strings.Split(s[1], "-")[0])
		secondPartHigher, _ := strconv.Atoi(strings.Split(s[1], "-")[1])

		firstInSecond := firstPartLower >= secondPartLower && firstPartHigher <= secondPartHigher
		secondInFirst := secondPartLower >= firstPartLower && secondPartHigher <= firstPartHigher

		if firstInSecond || secondInFirst {
			toReturn++
		} 
	}

	fmt.Println(toReturn)

}

func second() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	toReturn := 0

	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, ",")

		firstPartLower, _ := strconv.Atoi(strings.Split(s[0], "-")[0])
		firstPartHigher, _ := strconv.Atoi(strings.Split(s[0], "-")[1])
		secondPartLower, _ := strconv.Atoi(strings.Split(s[1], "-")[0])
		secondPartHigher, _ := strconv.Atoi(strings.Split(s[1], "-")[1])

		overlapping := firstPartLower <= secondPartHigher && firstPartLower >= secondPartHigher || firstPartHigher <= secondPartHigher && firstPartHigher >= secondPartLower || secondPartLower <= firstPartHigher && secondPartLower >= firstPartHigher || secondPartHigher <= firstPartHigher && secondPartHigher >= firstPartLower

		if overlapping {
			toReturn++
		} 
	}

	fmt.Println(toReturn)

}