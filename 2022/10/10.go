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
	cycle := 0
	X := 1
	toReturn := 0

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		if s[0] == "noop" {
			cycle++
			if cycle%40 == 20 {
				toReturn += X * cycle
			}
		} else {
			amount, _ := strconv.Atoi(s[1])
			for i := 1; i <= 2; i++ {
				cycle++
				if cycle%40 == 20 {
					toReturn += X * cycle
				}
			}

			X += amount
		}
	}

	fmt.Println(toReturn)

}

func second() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	crt := ""

	for i := 0; i < 240; i++ {
		crt += "#"
	}

	scanner := bufio.NewScanner(file)
	cycle := 0
	X := 1

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		if s[0] == "noop" {

			if X == cycle%40 || X == (cycle-1)%40 || X == (cycle+1)%40 && cycle != 240 {
				crt = crt[:cycle] + "#" + crt[cycle+1:]
			} else if X == cycle%40 || X == (cycle-1)%40 || X == (cycle+1)%40 && cycle == 240 {
				crt = crt[:cycle] + "#"
			} else if cycle != 240 {
				crt = crt[:cycle] + "." + crt[cycle+1:]
			} else {
				crt = crt[:cycle] + "."
			}

			cycle++
		} else {
			amount, _ := strconv.Atoi(s[1])
			for i := 1; i <= 2; i++ {

				if X == cycle%40 || X == (cycle-1)%40 || X == (cycle+1)%40 && cycle != 240 {
					crt = crt[:cycle] + "#" + crt[cycle+1:]
				} else if X == cycle%40 || X == (cycle-1)%40 || X == (cycle+1)%40 && cycle == 240 {
					crt = crt[:cycle] + "#"
				} else if cycle != 240 {
					crt = crt[:cycle] + "." + crt[cycle+1:]
				} else {
					crt = crt[:cycle] + "."
				}
				cycle++
			}

			X += amount
		}
	}

	fmt.Println(crt[0:39])
	fmt.Println(crt[40:79])
	fmt.Println(crt[80:119])
	fmt.Println(crt[120:159])
	fmt.Println(crt[160:199])
	fmt.Println(crt[200:239])
}
