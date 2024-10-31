package main

import (
	"bufio"
	"fmt"
	"os"
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

		cardSplit := strings.Split(scanner.Text(), "|")
		numbers := strings.Split(cardSplit[0], ":")

		// fmt.Printf("%s\n", numbers[1])
		// fmt.Printf("%s | %s\n", strings.Replace(numbers[1], "  ", " ", -1), strings.Replace(cardSplit[1], "  ", " ", -1))

		myNumbers := strings.Split(strings.Replace(numbers[1], "  ", " ", -1), " ")
		winnerNumbers := strings.Split(strings.Replace(cardSplit[1], "  ", " ", -1), " ")
		points := 0

		for _, value := range myNumbers {

			for _, winnerValue := range winnerNumbers {
				if value == "" || winnerValue == "" {
					continue
				}

				// fmt.Printf("{%s}=={%s} \n", value, winnerValue)
				if value == winnerValue {
					if points == 0 {
						fmt.Printf("points={%d} x={%s} y={%s} \n", points, value, winnerValue)
						points = 1
					} else {
						points *= 2
						fmt.Printf("points={%d} x={%s} y={%s} \n", points, value, winnerValue)
					}
				}
			}
		}
		result += points
		fmt.Printf("%d \n", result)
	}
}
