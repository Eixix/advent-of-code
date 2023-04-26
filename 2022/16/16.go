package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	name        string
	flowRate    int
	childValves map[string]int
	isOpen      bool
}

func main() {
	first()
}

func parseValves(filename string) map[string]Valve {
	file, _ := os.Open(filename)
	defer file.Close()
	valves := make(map[string]Valve)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		valveName := s[1]
		flowRate, _ := strconv.Atoi(strings.ReplaceAll(s[4][5:], ";", ""))
		childValves := s[9:]
		editedChildValves := make(map[string]int)
		for _, v := range childValves {
			editedChildValves[strings.ReplaceAll(v, ",", "")] = 1
		}

		valves[valveName] = Valve{valveName, flowRate, editedChildValves, false}
	}

	return valves
}

func first() {
	valves := parseValves("small.txt")

	fmt.Println(valves)
}
