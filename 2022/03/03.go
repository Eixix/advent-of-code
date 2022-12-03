package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)


func main() {
	first()
	second()
}

func first() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	toReturn := 0

	for scanner.Scan() {
		text := scanner.Text()
		length := len(text)
		half := length / 2
		firstHalf := text[:half]
		secondHalf := text[half:]

		alreadyIn := []rune{}

		for _, firstCharacter := range firstHalf {
			if contains(alreadyIn, firstCharacter) {
				break
			}
			for _, secondCharacter := range secondHalf {
				if firstCharacter == secondCharacter {
					if contains(alreadyIn, secondCharacter) {
						break
					}
					if unicode.IsLower(firstCharacter) {
						toReturn += smallLetterToInt(firstCharacter)
						alreadyIn = append(alreadyIn, firstCharacter)
					} else {
						toReturn += capitalLetterToInt(firstCharacter)
						alreadyIn = append(alreadyIn, firstCharacter)
					}
					break
				}
			}
		}
	}

	fmt.Println(toReturn)
}

func second() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	first := ""
	second := ""
	third := ""

	toReturn := 0
	found := false

	for scanner.Scan() {
		text := scanner.Text()

		if first == "" {
			first = text
		} else if second == "" {
			second = text
		} else if third == "" {
			third = text
			for _, firstCharacter := range first {
				for _, secondCharacter := range second {
					if found {
						break
					}
					if firstCharacter == secondCharacter {
						for _, thirdCharacter := range third {
							if firstCharacter == thirdCharacter {
								if unicode.IsLower(firstCharacter) {
									toReturn += smallLetterToInt(firstCharacter)
									found = true
								} else {
									toReturn += capitalLetterToInt(firstCharacter)
									found = true
								}
								break
							}
						}
					}
				}
				if found {
					found = false
					first = ""
					second = ""
					third = ""
					break
				}
			}
		}
	}

	fmt.Println(toReturn)
}

func contains(s []rune, str rune) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func smallLetterToInt(letter rune) int {
	return int(letter) - 96
}

func capitalLetterToInt(letter rune) int {
	return int(letter) - 38
}