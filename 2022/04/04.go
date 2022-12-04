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
		text := scanner.Text()
		s := strings.Split(text, ",")

		firstPartLower := strings.Split(s[0], "-")[0]
		firstPartHigher := strings.Split(s[0], "-")[1]
		secondPartLower := strings.Split(s[1], "-")[0]
		secondPartHigher := strings.Split(s[1], "-")[1]

		if firstPartLower >= secondPartLower && firstPartHigher <= secondPartHigher {
			toReturn++
			fmt.Println(firstPartLower, firstPartHigher, secondPartLower, secondPartHigher)
		} else if secondPartLower >= firstPartLower && secondPartHigher <= firstPartHigher {
			toReturn++
			fmt.Println(firstPartLower, firstPartHigher, secondPartLower, secondPartHigher)
		}
	}

	fmt.Println(toReturn)

}