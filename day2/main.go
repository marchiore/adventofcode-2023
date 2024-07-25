package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	gameNumber int
	hands      []hand
}

type hand struct {
	size  int
	color string
}

func loadGames() {

}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineSplitted := strings.Split(scanner.Text(), ":")

		// gameNumber := strings.Replace(lineSplitted[0], "Game ", "", 1)

		picks := lineSplitted[1]

		handsPicked := strings.Split(picks, ";")

		for i := 0; i < len(handsPicked); i++ {
			cubesPicked := strings.Split(handsPicked[i], " ")

			num, err := strconv.Atoi(cubesPicked[1])

			if err != nil {
				fmt.Println("Erro ao converter string para int:", err)
			}

			fmt.Println(hand{size: num, color: strings.Replace(cubesPicked[2], ",", "", 1)})

		}
	}
}
