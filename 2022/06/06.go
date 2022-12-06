package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	first()
	second()
}

func first() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		toReturn := 0
		first := ""
		second := ""
		third := ""
		fourth := ""

		for _, v := range text {
			if first == "" {
				first = string(v)
			} else if second == "" {
				second = string(v)
			} else if third == "" {
				third = string(v)
			} else if fourth == "" {
				fourth = string(v)
			} else {
				first, second, third, fourth = second, third, fourth, string(v)
				if first != second && first != third && first != fourth && second != third && second != fourth && third != fourth {
					toReturn++
					break
				}
			}
			toReturn++
		}
		fmt.Println(toReturn)

	}
}

func second() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		toReturn := 0
		isOk := false

		for i := 13; i < len(text); i++ {
			m := make(map[rune]bool)
			for _, j := range text[i-13 : i+1] {
				_, ok := m[j]
				if ok {
					break
				}
				m[j] = true
				if len(m) == 14 {
					isOk = true
					toReturn = i + 1
					break
				}

			}

			if isOk {
				break
			}
		}

		fmt.Println(toReturn)

	}
}
