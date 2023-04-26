package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	first()
}

func first() {
	file, _ := os.Open("small.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
