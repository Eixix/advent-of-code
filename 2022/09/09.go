package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ropeEnd struct {
	x int
	y int
}

func main() {
	first()
}

func first() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	head := ropeEnd{0, 0}
	tail := ropeEnd{0, 0}
	visitedFields := make(map[string]bool)
	visitedFields["0 0"] = true

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		switch s[0] {
		case "R":
			amount, _ := strconv.Atoi(s[1])
			head = moveHead(head, amount, 0)
			tail, visitedFields = moveTail(tail, head, visitedFields)
		case "U":
			amount, _ := strconv.Atoi(s[1])
			head = moveHead(head, 0, amount)
			tail, visitedFields = moveTail(tail, head, visitedFields)
		case "L":
			amount, _ := strconv.Atoi(s[1])
			head = moveHead(head, -amount, 0)
			tail, visitedFields = moveTail(tail, head, visitedFields)
		case "D":
			amount, _ := strconv.Atoi(s[1])
			head = moveHead(head, 0, -amount)
			tail, visitedFields = moveTail(tail, head, visitedFields)
		}
	}

	fmt.Println(len(visitedFields))

}

func moveHead(head ropeEnd, x int, y int) ropeEnd {
	return ropeEnd{head.x + x, head.y + y}
}

func moveTail(tail ropeEnd, head ropeEnd, visitedFields map[string]bool) (ropeEnd, map[string]bool) {
	tailX := tail.x
	tailY := tail.y

	for head.x > tailX+1 || head.x < tailX-1 || head.y > tailY+1 || head.y < tailY-1 {
		if head.x-tailX == 1 {
			if head.y > tailY+1 {
				tailY++
				tailX++
			}
			if head.y < tailY-1 {
				tailY--
				tailX++
			}
		} else if head.x-tailX == -1 {
			if head.y > tailY+1 {
				tailY++
				tailX--
			}
			if head.y < tailY-1 {
				tailY--
				tailX--
			}
		} else if head.y-tailY == 1 {
			if head.x > tailX+1 {
				tailX++
				tailY++
			}
			if head.x < tailX-1 {
				tailX--
				tailY++
			}
		} else if head.y-tailY == -1 {
			if head.x > tailX+1 {
				tailX++
				tailY--
			}
			if head.x < tailX-1 {
				tailX--
				tailY--
			}
		} else {
			if head.x > tailX+1 {
				tailX++
			}
			if head.x < tailX-1 {
				tailX--
			}
			if head.y > tailY+1 {
				tailY++
			}
			if head.y < tailY-1 {
				tailY--
			}
		}
		visitedFields[fmt.Sprint(tailX, tailY)] = true
	}
	return ropeEnd{tailX, tailY}, visitedFields
}
