package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var input [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, num := range strings.Fields(line) {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Err: ", err)
			}
			row = append(row, val)
		}
		input = append(input, row)
	}
	validCount := 0
	for i := 0; i < len(input); i++ {
		prev := input[i][0]
		if len(input[i]) < 2 || input[i][0] == input[i][1] {
			continue
		}
		increasing := input[i][1] > input[i][0]
		for j := 1; j < len(input[i]); j++ {
			diff := input[i][j] - prev
			if diff == 0 || (diff > 0 && !increasing) {
				break
			}
			if diff < 0 {
				if increasing {
					break
				}
				diff = -diff
			}
			if diff > 3 {
				break
			}

			prev = input[i][j]

			if j == len(input[i])-1 {
				fmt.Println("Safe Arr: ", input[i])
				validCount++
			}
		}
	}
	fmt.Println("Valid Count: ", validCount)
}
