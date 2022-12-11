package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int
	operation   func(int) int
	test        func(int) int
	inspections int
}

func (m *Monkey) setItems(items []string) {
	var toSet []int

	for _, item := range items {
		convertedItem, _ := strconv.Atoi(item[:2])
		toSet = append(toSet, convertedItem)
	}

	m.items = toSet
}

func (m *Monkey) setOperation(function []string) {
	if function[0] == "+" {
		if function[1] == "old" {
			m.operation = func(i int) int { return i + i }
		} else {
			number, _ := strconv.Atoi(function[1])
			m.operation = func(i int) int { return i + number }
		}
	} else {
		if function[1] == "old" {
			m.operation = func(i int) int { return i * i }
		} else {
			number, _ := strconv.Atoi(function[1])
			m.operation = func(i int) int { return i * number }
		}
	}
}

func (m *Monkey) setTest(divisor int, trueMonkey int, falseMonkey int) {
	m.test = func(value int) int {
		if value%divisor == 0 {
			return trueMonkey
		} else {
			return falseMonkey
		}
	}
}

func parseMonkeys(filename string) map[int]Monkey {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	monkeys := make(map[int]Monkey)
	var id int
	var currentMonkey Monkey
	var divisor int
	var trueMonkeyId int
	var falseMonkeyId int

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		if len(s) > 0 {
			if s[0] == "Monkey" {
				id, _ = strconv.Atoi(s[1][:1])
				currentMonkey = Monkey{nil, nil, nil, 0}
			} else if s[0] == "Starting" {
				currentMonkey.setItems(s[2:])
			} else if s[0] == "Operation:" {
				currentMonkey.setOperation(s[4:])
			} else if s[0] == "Test:" {
				divisor, _ = strconv.Atoi(s[3])
			} else if s[1] == "true:" {
				trueMonkeyId, _ = strconv.Atoi(s[5])
			} else if s[1] == "false:" {
				falseMonkeyId, _ = strconv.Atoi(s[5])
				currentMonkey.setTest(divisor, trueMonkeyId, falseMonkeyId)
				monkeys[id] = currentMonkey
			}
		}
	}

	return monkeys
}

func main() {
	first()
}

func first() {
	monkeys := parseMonkeys("small.txt")
	fmt.Println(monkeys)
	rounds := 20

	for i := 1; i <= rounds; i++ {
		for j := range monkeys {
			currentMonkey := monkeys[j]

			for _, item := range currentMonkey.items {
				item := currentMonkey.operation(item) / 3
				targetMonkeyId := currentMonkey.test(item)
				currentMonkey.inspections++
				targetMonkey := monkeys[targetMonkeyId]
				targetMonkey.items = append(targetMonkey.items, item)
				monkeys[targetMonkeyId] = targetMonkey
			}

			currentMonkey.items = []int{}
			monkeys[j] = currentMonkey
		}
	}

	fmt.Println(monkeys)
}
