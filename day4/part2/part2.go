package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ArrayList struct {
	data []int
}

func (al *ArrayList) incrementIndexInArray(index int, linePoints int) {

	for x := 1; x <= al.data[index]; x++ {
		for i := 1; i <= linePoints; i++ {
			al.data[index+i]++
		}
	}
}

func (al *ArrayList) SumAll() int {
	result := 0
	for _, value := range al.data {
		result += value
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

	line := 0

	arrayList := ArrayList{}

	for scanner.Scan() {
		arrayList.data = append(arrayList.data, 1)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {

		cardSplit := strings.Split(scanner.Text(), "|")
		numbers := strings.Split(cardSplit[0], ":")

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
						// fmt.Printf("points={%d} x={%s} y={%s} \n", points, value, winnerValue)
						points = 1
					} else {
						points++
						// fmt.Printf("points={%d} x={%s} y={%s} \n", points, value, winnerValue)
					}
				}
			}
		}
		// fmt.Printf("linha %d total de pontos %d \n", line, points)

		arrayList.incrementIndexInArray(line, points)

		// fmt.Printf("ESTADO DO ARRAY \n")

		// for index, value := range arrayList.data {
		// 	fmt.Printf("index %d value %d \n", index, value)
		// }
		line++
	}
	fmt.Printf("%d \n", arrayList.SumAll())

}
