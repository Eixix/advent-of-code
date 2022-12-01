package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	highestSum := 0.0
	sum := 0.0

	for scanner.Scan() {
		if (scanner.Text() == "") {
			highestSum = math.Max(highestSum, sum)
			sum = 0.0
		} else {
			tmp, _ := strconv.ParseFloat(scanner.Text(), 32)
			sum += tmp
		}
	}

	fmt.Println(highestSum)
}

func second() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var sums []float64
	sum := 0.0

	for scanner.Scan() {
		if (scanner.Text() == "") {
			sums = append(sums, sum)
			sum = 0.0
		} else {
			tmp, _ := strconv.ParseFloat(scanner.Text(), 32)
			sum += tmp
		}
	}

	sort.Slice(sums, func(i, j int) bool {
    	return sums[i] > sums[j]
	})
	toReturn := 0.0

	for i := 0; i < 3; i++ {
		toReturn += sums[i]
	}

	fmt.Println(toReturn)
}