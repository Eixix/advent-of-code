package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	first()
}

func first() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	toReturn := 0

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		switch s[0] {
		case "A":
			toReturn += A(s[1])
		case "B":
			toReturn += B(s[1])
		case "C":
			toReturn += C(s[1])
		}
	}

	fmt.Println(toReturn)
}

func A(input string) int {
	switch input {
	case "X":
		return 1 + 3
	case "Y":
		return 1 + 0
	case "Z":
		return 1 + 6
	default:
		panic("Can't happen")
	}
}

func B(input string) int {
	switch input {
	case "X":
		return 2 + 6
	case "Y":
		return 2 + 3
	case "Z":
		return 2 + 0
	default:
		panic("Can't happen")
	}
}

func C(input string) int {
	switch input {
	case "X":
		return 3 + 0
	case "Y":
		return 3 + 6
	case "Z":
		return 3 + 3
	default:
		panic("Can't happen")
	}
}