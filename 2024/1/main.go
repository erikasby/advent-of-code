package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("values.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		if len(parts) != 2 {
			log.Fatal("More than 2 values per line.")
		}

		leftInt, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Unable to convert string to int.", err)
		}

		rightInt, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("Unable to convert string to int.", err)
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	if len(left) != len(right) {
		log.Fatal("Different lengths of left and right arrays.")
	}

	var distance int = 0
	for i := range left {
		var diff int = absDiffInt(left[i], right[i])
		distance += diff
	}

	fmt.Println("Answer:", distance)
}

func absDiffInt(x, y int) int {
	var diff int = x - y
	if diff < 0 {
		return diff * -1
	}
	return diff
}
