package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	useRegExp := false

	if useRegExp {
		addWithRegExp()
	} else {
		addWithoutRegExp()
	}
}

func addWithRegExp() {
	file, err := os.Open("values.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Drop RegExp if performance issues occur
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	addedResults := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			match = match[4 : len(match)-1]
			nums := strings.Split(match, ",")

			left, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal("Unable to convert string to int.", err)
			}

			right, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal("Unable to convert string to int.", err)
			}

			fmt.Println(match)

			addedResults += left * right
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(addedResults) // 182619815
}

func addWithoutRegExp() {
	// Open the file
	file, err := os.Open("values.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	addedResults := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		currentPosition := 0
		for {
			start := strings.Index(line[currentPosition:], "mul(")
			if start == -1 {
				break // No more "mul(" found
			}
			currentPosition += start

			end := strings.Index(line[currentPosition:], ")")
			if end == -1 {
				break // No more ")" found
			}
			end += currentPosition
			
			insideParens := line[currentPosition+4 : end] // Excluding "mul("

			// TODO: Refactor into something better? :D
			fmt.Println("Old: " + insideParens)
			for {
				insideParensAnotherMulIndex := strings.Index(line[currentPosition+4 : end], "mul(")
				if insideParensAnotherMulIndex == -1 || insideParensAnotherMulIndex == 0 {
					break // No more "mul(" found
				} else {
					currentPosition += insideParensAnotherMulIndex
					insideParens = line[currentPosition+4+4 : end]
					fmt.Println("Adj: " + insideParens)
				}
			}
			
			fmt.Println("End: " + insideParens)
			nums := strings.Split(insideParens, ",")
			
			left, err := strconv.Atoi(nums[0])
			if err != nil {
				currentPosition = end
				continue
			}
			
			right, err := strconv.Atoi(nums[1])
			if err != nil {
				currentPosition = end
				continue
			}


			addedResults += left * right
			currentPosition = end
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(addedResults)
}
