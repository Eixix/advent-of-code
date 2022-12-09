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

type ropeKnot struct {
	child *ropeKnot
	x     int
	y     int
}

func main() {
	first()
	second()
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

func second() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tail := ropeKnot{nil, 0, 0}
	knot8 := ropeKnot{&tail, 0, 0}
	knot7 := ropeKnot{&knot8, 0, 0}
	knot6 := ropeKnot{&knot7, 0, 0}
	knot5 := ropeKnot{&knot6, 0, 0}
	knot4 := ropeKnot{&knot5, 0, 0}
	knot3 := ropeKnot{&knot4, 0, 0}
	knot2 := ropeKnot{&knot3, 0, 0}
	knot1 := ropeKnot{&knot2, 0, 0}
	head := ropeKnot{&knot1, 0, 0}

	visitedFields := make(map[string]bool)
	visitedFields["0 0"] = true

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		switch s[0] {
		case "R":
			amount, _ := strconv.Atoi(s[1])
			head = ropeKnot{&knot1, head.x + amount, head.y}
			tail, visitedFields = moveRope(head, visitedFields)
		case "U":
			amount, _ := strconv.Atoi(s[1])
			head = ropeKnot{&knot1, head.x, head.y + amount}
			tail, visitedFields = moveRope(head, visitedFields)
		case "L":
			amount, _ := strconv.Atoi(s[1])
			head = ropeKnot{&knot1, head.x - amount, head.y}
			tail, visitedFields = moveRope(head, visitedFields)
		case "D":
			amount, _ := strconv.Atoi(s[1])
			head = ropeKnot{&knot1, head.x, head.y - amount}
			tail, visitedFields = moveRope(head, visitedFields)
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

func moveRope(knot ropeKnot, visitedFields map[string]bool) (ropeKnot, map[string]bool) {
	tailX := knot.child.x
	tailY := knot.child.y

	for knot.x > tailX+1 || knot.x < tailX-1 || knot.y > tailY+1 || knot.y < tailY-1 {
		if knot.x-tailX == 1 {
			if knot.y > tailY+1 {
				tailY++
				tailX++
			}
			if knot.y < tailY-1 {
				tailY--
				tailX++
			}
		} else if knot.x-tailX == -1 {
			if knot.y > tailY+1 {
				tailY++
				tailX--
			}
			if knot.y < tailY-1 {
				tailY--
				tailX--
			}
		} else if knot.y-tailY == 1 {
			if knot.x > tailX+1 {
				tailX++
				tailY++
			}
			if knot.x < tailX-1 {
				tailX--
				tailY++
			}
		} else if knot.y-tailY == -1 {
			if knot.x > tailX+1 {
				tailX++
				tailY--
			}
			if knot.x < tailX-1 {
				tailX--
				tailY--
			}
		} else {
			if knot.x > tailX+1 {
				tailX++
			}
			if knot.x < tailX-1 {
				tailX--
			}
			if knot.y > tailY+1 {
				tailY++
			}
			if knot.y < tailY-1 {
				tailY--
			}
		}
		if knot.child.child == nil {
			visitedFields[fmt.Sprint(tailX, tailY)] = true
		}
	}
	knot.child.x = tailX
	knot.child.y = tailY

	if knot.child.child == nil {
		return *knot.child, visitedFields
	}

	return moveRope(*knot.child, visitedFields)
}
