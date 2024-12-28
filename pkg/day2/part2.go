package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution2() {
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
		faultCount := 0
		for j := 1; j < len(input[i]); j++ {
			if faultCount > 1 {
				fmt.Println("Fault Index: ", i, "Arr: ", input[i])
				break
			}
			diff := input[i][j] - prev

			if !(1 <= abs(diff) && abs(diff) <= 3) || (diff > 0 && !increasing) || (diff < 0 && increasing) {
				faultCount++
				continue
			}

			prev = input[i][j]

			if j == len(input[i])-1 {
				// fmt.Println("Safe Arr: ", input[i])
				validCount++
			}
		}
	}
	fmt.Println("Valid Count: ", validCount)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
