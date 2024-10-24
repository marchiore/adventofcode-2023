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
	for !findChar {

		if forward {
			if isNumber(line[startPosition+1 : endPosition+3]) {
				num, err := strconv.Atoi(line[startPosition+1 : endPosition+3])
				if err == nil {
					findChar = true
					return num
				}
			} else if isNumber(line[startPosition+1 : endPosition+2]) {
				num, err := strconv.Atoi(line[startPosition+1 : endPosition+2])
				if err == nil {
					findChar = true
					return num
				}
			} else if isNumber(line[startPosition+1 : endPosition+1]) {
				num, err := strconv.Atoi(line[startPosition+1 : endPosition+1])
				if err == nil {
					findChar = true
					return num
				}
			}
		} else {
			if isNumber(line[startPosition-3 : endPosition-1]) {
				num, err := strconv.Atoi(line[startPosition-3 : endPosition-1])
				if err == nil {
					findChar = true
					return num
				}
			} else if isNumber(line[startPosition-2 : endPosition-1]) {
				num, err := strconv.Atoi(line[startPosition-2 : endPosition-1])
				if err == nil {
					findChar = true
					return num
				}
			} else if isNumber(line[startPosition-1 : endPosition-1]) {
				num, err := strconv.Atoi(line[startPosition-1 : endPosition-1])
				if err == nil {
					findChar = true
					return num
				}
			}
		}

	}
	return 0
}

func checkAsteriskConnectedNumbers(line string, lineAbove string, lineBelow string, init int, end int) int {

	var arr []int

	charBefore := line[init-1 : end-1]
	if isNumber(charBefore) {
		arr = append(arr, findTheWholeNumber(line, init, end, false))
	}

	charAfter := line[init+1 : end+1]
	if isNumber(charAfter) {
		arr = append(arr, findTheWholeNumber(line, init, end, true))
	}

	// check line below
	if lineBelow != "" {
		shouldIgnoreSpacedNumberBelow := false
		if isNumber(lineBelow[init-1:end-1]) && isNumber(lineBelow[init+1:end+1]) {
			// .123.
			//  .*.
			num, err := strconv.Atoi(lineBelow[init-1 : end+1])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberBelow = true
			}
		} else if isNumber(lineBelow[init-2 : end]) {
			// .123...
			// ...*...
			num, err := strconv.Atoi(lineBelow[init-2 : end])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberBelow = true
			}
		} else if isNumber(lineBelow[init : end+2]) {
			// ...123.
			// ...*...
			num, err := strconv.Atoi(lineBelow[init : end+2])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberBelow = true
			}
		} else if isNumber(lineBelow[init-1 : end]) {
			// ..12...
			// ...*...
			num, err := strconv.Atoi(lineBelow[init-1 : end])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberBelow = true
			}
		} else if isNumber(lineBelow[init : end+1]) {
			// ...12..
			// ...*...
			num, err := strconv.Atoi(lineBelow[init : end+1])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberBelow = true
			}
		} else if isNumber(lineBelow[init:end]) {
			// ...1...
			// ...*...
			num, err := strconv.Atoi(lineBelow[init:end])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberBelow = true
			}
		}

		if len(arr) == 0 && !shouldIgnoreSpacedNumberBelow {
			if isNumber(lineBelow[init-3 : end-1]) {
				// 123....
				// ...*...
				num, err := strconv.Atoi(lineBelow[init-3 : end-1])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineBelow[init-2 : end-1]) {
				// .12....
				// ...*...
				num, err := strconv.Atoi(lineBelow[init-2 : end-1])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineBelow[init-1 : end-1]) {
				// ..1....
				// ...*...
				num, err := strconv.Atoi(lineBelow[init-1 : end-1])

				if err == nil {
					arr = append(arr, num)
				}
			}

			if isNumber(lineBelow[init+1 : end+3]) {
				// ....123
				// ...*...
				num, err := strconv.Atoi(lineBelow[init+1 : end+3])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineBelow[init+1 : end+2]) {
				// ....12.
				// ...*...
				num, err := strconv.Atoi(lineBelow[init+1 : end+3])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineBelow[init+1 : end+1]) {
				// ....1..
				// ...*...
				num, err := strconv.Atoi(lineBelow[init+1 : end+1])

				if err == nil {
					arr = append(arr, num)
				}
			}
		}

	}

	if lineAbove != "" && len(arr) < 2 {
		shouldIgnoreSpacedNumberAbove := false

		if isNumber(lineAbove[init-1:end-1]) && isNumber(lineAbove[init+1:end+1]) {
			// .123.
			//  .*.
			num, err := strconv.Atoi(lineAbove[init-1 : end+1])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberAbove = true
			}
		} else if isNumber(lineAbove[init-2 : end]) {
			// .123...
			// ...*...
			num, err := strconv.Atoi(lineAbove[init-2 : end])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberAbove = true
			}
		} else if isNumber(lineAbove[init : end+2]) {
			// ...123.
			// ...*...
			num, err := strconv.Atoi(lineAbove[init : end+2])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberAbove = true
			}
		} else if isNumber(lineAbove[init-1 : end]) {
			// ..12...
			// ...*...
			num, err := strconv.Atoi(lineAbove[init-1 : end])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberAbove = true
			}
		} else if isNumber(lineAbove[init : end+1]) {
			// ...12..
			// ...*...
			num, err := strconv.Atoi(lineAbove[init : end+1])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberAbove = true
			}
		} else if isNumber(lineAbove[init:end]) {
			// ...1...
			// ...*...
			num, err := strconv.Atoi(lineAbove[init:end])

			if err == nil {
				arr = append(arr, num)
				shouldIgnoreSpacedNumberAbove = true
			}
		}

		if len(arr) <= 1 && !shouldIgnoreSpacedNumberAbove {
			if isNumber(lineAbove[init-3 : end-1]) {
				// 123....
				// ...*...
				num, err := strconv.Atoi(lineAbove[init-3 : end-1])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineAbove[init-2 : end-1]) {
				// .12....
				// ...*...
				num, err := strconv.Atoi(lineAbove[init-2 : end-1])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineAbove[init-1 : end-1]) {
				// ..1....
				// ...*...
				num, err := strconv.Atoi(lineAbove[init-1 : end-1])

				if err == nil {
					arr = append(arr, num)
				}
			}

			if isNumber(lineAbove[init+1 : end+3]) {
				// ....123
				// ...*...
				num, err := strconv.Atoi(lineAbove[init+1 : end+3])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineAbove[init+1 : end+2]) {
				// ....12.
				// ...*...
				num, err := strconv.Atoi(lineAbove[init+1 : end+3])

				if err == nil {
					arr = append(arr, num)
				}
			} else if isNumber(lineAbove[init+1 : end+1]) {
				// ....1..
				// ...*...
				num, err := strconv.Atoi(lineAbove[init+1 : end+1])

				if err == nil {
					arr = append(arr, num)
				}
			}
		}

	}

	if len(arr) > 2 {
		for _, value := range arr {
			fmt.Println(value)
		}
		fmt.Printf("BUG")
	}

	if len(arr) > 1 {
		fmt.Printf("multiplicando os numeros {%d} * {%d} \n", arr[0], arr[1])
		return arr[0] * arr[1]
	}

	return 0
}

func main() {

	result := 0

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

	for i, line := range lines {
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {

			init := match[0]
			end := match[1]

			upperLine := ""
			lowerLine := ""

			if i > 1 {
				upperLine = lines[i-1]
			}

			if i < len(lines) {
				lowerLine = lines[i+1]
			}

			num := checkAsteriskConnectedNumbers(line, upperLine, lowerLine, init, end)

			if num > 0 {
				result += num
			}
		}
	}
	println(result)
}
