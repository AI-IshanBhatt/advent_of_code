package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func decodeLinesAndSum(lines [][]string) int {

	answer := 0
	for _, s_array := range lines {
		i, _ := strconv.Atoi(s_array[0])
		j, _ := strconv.Atoi(s_array[len(s_array)-1])
		answer += 10*i + j
	}
	return answer
}

func problemOne(fileLocation string) int {

	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	answer := make([][]string, 0)

	for sc.Scan() {
		digits := make([]string, 0)
		for _, c := range sc.Text() {
			//x := fmt.Sprintf("%c", c)
			if unicode.IsDigit(c) {
				digits = append(digits, fmt.Sprintf("%c", c))
			}
		}
		answer = append(answer, digits)
	}

	return decodeLinesAndSum(answer)
}

func problemTwo(fileLocation string) int {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	words := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8",
		"nine": "9",
	}

	sc := bufio.NewScanner(file)
	answer := make([][]string, 0)

	for sc.Scan() {
		digits := make([]string, 0)
		line := sc.Text()
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				digits = append(digits, fmt.Sprintf("%c", line[i]))
			} else {
				for k, v := range words {
					if line[i:min(i+len(k), len(line))] == k {
						digits = append(digits, v)
					}
				}
			}
		}
		answer = append(answer, digits)
	}

	return decodeLinesAndSum(answer)
}

func main() {
	fmt.Println(problemOne("day_one_input"))
	fmt.Println(problemTwo("day_one_input"))
}
