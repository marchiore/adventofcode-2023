package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func findTheWholeNumber(line string, startPosition int, endPosition int, forward bool) int {

	findChar := false
	for findChar == false {
		if isNumber(line[startPosition:endPosition]) {
			if forward {
				endPosition += 1
			} else {
				startPosition -= 1
			}
		} else {
			findChar = true

			if forward {
				endPosition -= 1
			} else {
				startPosition += 1
			}

			// fmt.Printf("number i={%s} \n", line[startPosition:endPosition])

			num, err := strconv.Atoi(line[startPosition:endPosition])
			if err == nil {
				fmt.Printf("number i={%d} \n", num)
				return num
			}
		}
	}
	return 0
}

func checkAsteriskConnectedNumbers(line string, lineAbove string, lineBelow string, init int, end int) string {

	// check same line
	charBefore := line[init-1 : end-1]
	if isNumber(charBefore) {
		findTheWholeNumber(line, init-1, end-1, false)
	}

	charAfter := line[init+1 : end+1]
	if isNumber(charAfter) {
		findTheWholeNumber(line, init+1, end+1, true)
	}

	return ""
}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines := make(map[int]string)
	lineNumber := 1

	re := regexp.MustCompile(`\*`)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines[lineNumber] = scanner.Text()
		lineNumber++
	}

	result := 0

	for i, line := range lines {
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {

			init := match[0]
			end := match[1]

			// fmt.Printf("number i={%d} init={%d} end={%d} \n", i, init, end)

			checkAsteriskConnectedNumbers(line, lines[i], lines[i], init, end)

		}
	}
	println(result)
}
