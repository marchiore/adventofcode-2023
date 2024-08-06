package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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
			// fmt.Printf("%d, %d", match[0], match[1])
			// fmt.Println("\n")

			// fmt.Println(line[match[0]:match[1]])

			number := ""

			if end >= 139 {
				// fmt.Println(line[init-1 : end])
				number = line[init-1 : end]
			} else {
				// fmt.Println(line[init-1 : end+1])
				number = line[init-1 : end+1]
			}

			if regexSpecialCharacters.MatchString(number) {
				// fmt.Println(number)

				num, _ := strconv.Atoi(line[init:end])

				result += num
			} else {
				// ler a linha de cima e de baixo para ver se na posição tem caracter especial

				fmt.Printf("line={%d} init={%d}, end={%d} \n", i, init, end)

				if line

			}

		}
	}

	println(result)
}
