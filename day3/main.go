package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func checkLine(line string, init int, end int) string {

	startRange := 0
	endRange := 0

	if init == 0 {
		startRange = 1
	} else {
		startRange = init - 1
	}

	if end >= 139 {
		endRange = end
	} else {
		endRange = end + 1
	}

	return line[startRange:endRange]
}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines := make(map[int]string)
	lineNumber := 1

	re := regexp.MustCompile(`\d+`)
	regexSpecialCharacters := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-\=\[\]\{\}\\|;:'",<>/\?\~` + "`" + `]`)

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

			number := checkLine(line, init, end)

			if regexSpecialCharacters.MatchString(number) {

				num, _ := strconv.Atoi(line[init:end])

				result += num
			} else {

				// primeira linha
				if i == 1 {
					number := checkLine(lines[i+1], init, end)

					if regexSpecialCharacters.MatchString(number) {
						num, _ := strconv.Atoi(line[init:end])
						result += num
						fmt.Printf("number in={%d} line={%d} \n", num, i)
					}
					// fmt.Printf("line={%s} init={%d}, end={%d} \n", lines[2][startRange:endRange], init, end)
				} else if i == lineNumber-1 {

					number := checkLine(lines[i-1], init, end)

					if regexSpecialCharacters.MatchString(number) {
						num, _ := strconv.Atoi(line[init:end])
						result += num
						fmt.Printf("number in={%d} line={%d} \n", num, i)
					}
				} else {

					if regexSpecialCharacters.MatchString(checkLine(lines[i-1], init, end)) {
						num, _ := strconv.Atoi(line[init:end])
						result += num
						fmt.Printf("number in={%d} line={%d} \n", num, i)
					} else if regexSpecialCharacters.MatchString(checkLine(lines[i+1], init, end)) {
						num, _ := strconv.Atoi(line[init:end])
						result += num
						fmt.Printf("number in={%d} line={%d} \n", num, i)
					}

				}
			}

		}
	}
	// 539713
	println(result)
}
