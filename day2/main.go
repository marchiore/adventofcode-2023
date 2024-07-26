package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxBlue = 14
var maxRed = 12
var maxGreen = 13

func hasInvalidNumberOfCubes(color string, num int) bool {
	result := false

	switch color {
	case "blue":
		// fmt.Printf("%d > %d", num, maxBlue)
		if num > maxBlue {
			result = true
		}
	case "green":
		// fmt.Printf("%d > %d", num, maxGreen)
		if num > maxGreen {
			result = true
		}
	case "red":
		// fmt.Printf("%d > %d", num, maxRed)
		if num > maxRed {
			result = true
		}
	}
	return result
}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {

		lineSplitted := strings.Split(scanner.Text(), ":")

		picks := lineSplitted[1]

		gameNumber := strings.Replace(lineSplitted[0], "Game ", "", -1)

		handsPicked := strings.Split(picks, ";")

		gameIsPossible := true

		for i := 0; i < len(handsPicked); i++ {

			cubesPicked := strings.Split(strings.TrimSpace(handsPicked[i]), ",")
			// fmt.Println(cubesPicked)

			for x := 0; x < len(cubesPicked); x++ {

				colors := strings.TrimSpace(cubesPicked[x])

				cubes := strings.Split(colors, " ")

				// fmt.Println(cubes[0])

				num, _ := strconv.Atoi(cubes[0])
				color := cubes[1]

				hasInvalidNumberOfCubes := hasInvalidNumberOfCubes(color, num)

				// fmt.Printf(" %d - %s : %t", num, color, hasInvalidNumberOfCubes)

				if hasInvalidNumberOfCubes {
					gameIsPossible = false
				}
			}
		}

		if gameIsPossible {
			num, _ := strconv.Atoi(gameNumber)
			result += num
		}

		fmt.Printf("\n")

	}

	fmt.Printf("%d", result)
}
