package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//             [M] [S] [S]
//         [M] [N] [L] [T] [Q]
// [G]     [P] [C] [F] [G] [T]
// [B]     [J] [D] [P] [V] [F] [F]
// [D]     [D] [G] [C] [Z] [H] [B] [G]
// [C] [G] [Q] [L] [N] [D] [M] [D] [Q]
// [P] [V] [S] [S] [B] [B] [Z] [M] [C]
// [R] [H] [N] [P] [J] [Q] [B] [C] [F]
//  1   2   3   4   5   6   7   8   9


func main() {
	cargo := [][]string{
		{"G", "B", "D", "C", "P", "R"},
		{"G", "V", "H"},
		{"M", "P", "J", "D", "Q" ,"S", "N"},
		{"M", "N", "C", "D", "G" ,"L", "S", "P"},
		{"S", "L", "F", "P", "C" ,"N", "B", "J"},
		{"S", "T", "G", "V", "Z" ,"D", "B", "Q"},
		{"Q", "T", "F", "H", "M" ,"Z", "B"},
		{"F", "B", "D", "M", "C"},
		{"G", "Q", "C", "F"},
	}


	//first(cargo)
	second(cargo)
}

func first(cargo [][]string) {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")

		cargoAmount, _ := strconv.Atoi(s[1])
		startStack, _ := strconv.Atoi(s[3])
		startStack -= 1
		endStack, _ := strconv.Atoi(s[5])
		endStack -= 1

		for i := 1; i <= cargoAmount; i++ {
			var container string
			container, cargo[startStack] = cargo[startStack][0], cargo[startStack][1:]
			cargo[endStack] = append([]string{container}, cargo[endStack]...)
		}
	}

	for _, v := range cargo {
		fmt.Print(v[0])
	}

}

func second(cargo [][]string) {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")

		cargoAmount, _ := strconv.Atoi(s[1])
		startStack, _ := strconv.Atoi(s[3])
		startStack -= 1
		endStack, _ := strconv.Atoi(s[5])
		endStack -= 1

		var containers []string
		
		// I hate call by reference
		toMove := append([]string(nil), cargo[startStack][0:cargoAmount]...)
		rest := append([]string(nil), cargo[startStack][cargoAmount:]...)

		containers, cargo[startStack] = toMove, rest
		cargo[endStack] = append(containers, cargo[endStack]...)
	}

	for _, v := range cargo {
		fmt.Print(v[0])
	}

}