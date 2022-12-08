package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	first()
	second()
}

func first() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trees := [][]int{}

	for scanner.Scan() {
		trees = append(trees, stringToIntSlice(scanner.Text()))
	}

	visibleTrees := 0

	// Visit the trees
	for i, column := range trees {
		for j, tree := range column {
			// Tree on the outside
			if i == 0 || j == 0 || i == len(trees)-1 || j == len(column)-1 {
				visibleTrees++
			} else {
				visible := false
				// Left to right
				for k := j + 1; k < len(column); k++ {
					if k == len(column)-1 && tree > column[k] {
						visible = true
						break
					} else if tree > column[k] {
						continue
					} else {
						break
					}
				}
				// Right to left
				for k := j - 1; k >= 0; k-- {
					if k == 0 && column[k] < tree {
						visible = true
						break
					} else if column[k] < tree {
						continue
					} else {
						break
					}
				}
				// Top to bottom
				for k := i + 1; k < len(trees); k++ {
					if k == len(trees)-1 && tree > trees[k][j] {
						visible = true
						break
					} else if tree > trees[k][j] {
						continue
					} else {
						break
					}
				}
				// Bottom to top
				for k := i - 1; k >= 0; k-- {
					if k == 0 && trees[k][j] < tree {
						visible = true
						break
					} else if trees[k][j] < tree {
						continue
					} else {
						break
					}
				}

				if visible {
					visibleTrees++
				}
			}
		}
	}

	fmt.Println(visibleTrees)
}

func second() {
	file, _ := os.Open("small.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trees := [][]int{}

	for scanner.Scan() {
		trees = append(trees, stringToIntSlice(scanner.Text()))
	}

	highestScenicScore := 0

	// Visit the trees
	for i, column := range trees {
		for j, tree := range column {
			scenicScore := 1
			// Left to right
			for k := j + 1; k < len(column); k++ {
				if k == len(column)-1 && tree > column[k] {
					scenicScore *= k - j
					break
				} else if tree > column[k] {
					continue
				} else {
					scenicScore *= k - j
					break
				}
			}
			// Right to left
			for k := j - 1; k >= 0; k-- {
				if k == 0 && column[k] < tree {
					scenicScore *= j - k
					break
				} else if column[k] < tree {
					continue
				} else {
					scenicScore *= j - k
					break
				}
			}
			// Top to bottom
			for k := i + 1; k < len(trees); k++ {
				if k == len(trees)-1 && tree > trees[k][j] {
					scenicScore *= k - i
					break
				} else if tree > trees[k][j] {
					continue
				} else {
					scenicScore *= k - i
					break
				}
			}
			// Bottom to top
			for k := i - 1; k >= 0; k-- {
				if k == 0 && trees[k][j] < tree {
					scenicScore *= i - k
					break
				} else if trees[k][j] < tree {
					continue
				} else {
					scenicScore *= i - k
					break
				}
			}
			if highestScenicScore < scenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	fmt.Println(highestScenicScore)
}

func stringToIntSlice(text string) []int {
	toReturn := []int{}

	for _, v := range text {
		integer, _ := strconv.Atoi(string(v))
		toReturn = append(toReturn, integer)
	}

	return toReturn
}
