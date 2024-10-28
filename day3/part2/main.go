package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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
	regexSanitize := regexp.MustCompile(`(\.|\%|\/|\*|\$|\@|\&|\_|\+|\=|\-)\d{1,2}$|^\d{1,2}(\.|\%|\/|\*|\$|\@|\&|\_|\+|\=|\-)|^(\.|\%|\/|\*|\$|\@|\&|\_|\+|\=|\-)\d{1}(\.|\%|\/|\*|\$|\@|\&|\_|\+|\=|\-)`)
	regexNumber := regexp.MustCompile(`[0-9]+`)
	regexRight := regexp.MustCompile(`^\*\d{1,3}`)
	regexLeft := regexp.MustCompile(`\d{1,3}\*$`)

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

			upperLine := lines[i-1]
			lowerLine := lines[i+1]

			var numbers []int

			leftMatch := regexLeft.FindAllString(lines[i][init-3:end], -1)

			if len(leftMatch) > 0 {
				num, err := strconv.Atoi(re.ReplaceAllString(leftMatch[0], ""))

				if err == nil {
					numbers = append(numbers, num)
					// fmt.Printf("numero encontrado na esquerda: %s %d \n", lines[i][init-3:end], num)

				}
			}

			rightMatch := regexRight.FindAllString(lines[i][init:end+3], -1)

			if len(rightMatch) > 0 {

				num, err := strconv.Atoi(re.ReplaceAllString(rightMatch[0], ""))

				if err == nil {
					// fmt.Printf("numero encontrado na direita: %s %d \n", lines[i][init:end+3], num)
					numbers = append(numbers, num)
				}
			}

			sanitizeUpperMatch := regexSanitize.ReplaceAllString(upperLine[init-3:end+3], "")
			sanitizeLowerMatch := regexSanitize.ReplaceAllString(lowerLine[init-3:end+3], "")

			fmt.Printf("LINHA={%s} CIMA={%s} BAIXO={%s} \n", lines[i][init-3:end+3], upperLine[init-3:end+3], lowerLine[init-3:end+3])

			upperMatch := regexNumber.FindAllString(sanitizeUpperMatch, -1)
			lowerMatch := regexNumber.FindAllString(sanitizeLowerMatch, -1)

			if len(upperMatch) > 0 {
				for _, value := range upperMatch {
					// fmt.Println("Index:", index, "Value:", value)

					num, err := strconv.Atoi(value)

					if err == nil {
						// fmt.Printf("numero encontrado na linha de cima: %s %d \n", sanitizeUpperMatch, num)
						numbers = append(numbers, num)
					}
				}
			}

			if len(lowerMatch) > 0 {
				for _, value := range lowerMatch {
					// fmt.Println("Index:", index, "Value:", value)

					num, err := strconv.Atoi(value)

					if err == nil {
						// fmt.Printf("numero encontrado na linha de baixo: %s %d \n", sanitizeLowerMatch, num)
						numbers = append(numbers, num)
					}
				}
			}

			if len(numbers) > 2 {
				for i := 0; i < len(numbers); i++ {
					fmt.Printf("{%d}", numbers[i])
				}
				fmt.Printf("ERROR")
			}

			if len(numbers) == 2 {
				result += numbers[0] * numbers[1]

				fmt.Printf("multiplicando %d * %d \n", numbers[0], numbers[1])
			}

		}
	}
	println(result)
}
