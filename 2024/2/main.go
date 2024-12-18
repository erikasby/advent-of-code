package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("values.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeCounter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) == 0 {
			log.Fatal("Invalid line; no entries")
		}

		if len(parts) == 1 {
			log.Fatal("Invalid line; single entry")
		}

		isIncreasing := false
		isSafe := false
		for i := range parts {
			if i == 0 {
				current, err := strconv.Atoi(parts[i])
				if err != nil {
					log.Fatal("Unable to convert string to int.", err)
				}

				next, err := strconv.Atoi(parts[i+1])
				if err != nil {
					log.Fatal("Unable to convert string to int.", err)
				}

				if next > current {
					isIncreasing = true
				}
				continue
			}

			previous, err := strconv.Atoi(parts[i-1])
			if err != nil {
				log.Fatal("Unable to convert string to int.", err)
			}

			current, err := strconv.Atoi(parts[i])
			if err != nil {
				log.Fatal("Unable to convert string to int.", err)
			}

			// Rule 1: Either is decreasing or increasing; never same
			isSameValue := previous == current
			if isSameValue {
				break
			}

			hasDecreasedOnIncrease := previous > current && isIncreasing
			if hasDecreasedOnIncrease {
				break
			}

			hasIncreasedOnDecrease := previous < current && !isIncreasing
			if hasIncreasedOnDecrease {
				break
			}

			// Rule 2: Max. difference of abs(3)
			hasOverDecreased := previous-current < -3
			hasOverIncreased := previous-current > 3
			if hasOverIncreased || hasOverDecreased {
				break
			}

			isLast := i == len(parts)-1
			if isLast {
				isSafe = true
			}
		}

		if isSafe {
			safeCounter++
		}
	}

	fmt.Println("Answer:", safeCounter)
}
