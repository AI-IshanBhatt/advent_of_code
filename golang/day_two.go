package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var masterMap = map[string]int{
	"red": 12, "green": 13, "blue": 14,
}

func partOne(fileLocation string) int {

	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	answer := 0

	for sc.Scan() {
		flag := true
		line := sc.Text()

		gameConfig := strings.Split(line, ":")
		gameId, _ := strconv.Atoi(strings.Split(gameConfig[0], " ")[1])

		for _, game := range strings.Split(gameConfig[1], ";") {
			for _, colorConfig := range strings.Split(game, ",") {
				countColor := strings.Split(colorConfig, " ")
				i, _ := strconv.Atoi(countColor[1])
				if i > masterMap[countColor[2]] {
					flag = false
				}
			}
		}
		if flag {
			answer += gameId
		}
	}

	return answer
}

func partTwo(fileLocation string) int {

	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	answer := 0

	for sc.Scan() {
		line := sc.Text()

		maxColor := map[string]int{}

		gameConfig := strings.Split(line, ":")

		for _, game := range strings.Split(gameConfig[1], ";") {
			for _, colorConfig := range strings.Split(game, ",") {
				countColor := strings.Split(colorConfig, " ")
				i, _ := strconv.Atoi(countColor[1])
				maxColor[countColor[2]] = max(i, maxColor[countColor[2]])
			}
		}

		currentMult := 1
		for _, v := range maxColor {
			currentMult *= v
		}

		answer += currentMult
	}

	return answer
}

func main() {
	fmt.Println(partOne("day_two_input"))
	fmt.Println(partTwo("day_two_input"))
}
