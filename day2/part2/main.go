package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

		handsPicked := strings.Split(picks, ";")

		topBlue := 0
		topGreen := 0
		topRed := 0

		for i := 0; i < len(handsPicked); i++ {

			cubesPicked := strings.Split(strings.TrimSpace(handsPicked[i]), ",")
			// fmt.Println(cubesPicked)

			for x := 0; x < len(cubesPicked); x++ {

				colors := strings.TrimSpace(cubesPicked[x])

				cubes := strings.Split(colors, " ")

				// fmt.Println(cubes[0])

				num, _ := strconv.Atoi(cubes[0])
				color := cubes[1]

				switch color {
				case "blue":
					// fmt.Printf("%d > %d", num, maxBlue)
					if num > topBlue {
						topBlue = num
					}
				case "green":
					// fmt.Printf("%d > %d", num, maxGreen)
					if num > topGreen {
						topGreen = num
					}
				case "red":
					// fmt.Printf("%d > %d", num, maxRed)
					if num > topRed {
						topRed = num
					}
				}

			}
		}

		result += (topBlue * topGreen * topRed)

		fmt.Printf("\n")

	}

	fmt.Printf("%d", result)
}
