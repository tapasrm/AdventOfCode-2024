package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Solution1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(mulPattern)

		matches := re.FindAllStringSubmatch(line, -1)

		if matches != nil {
			for _, match := range matches {
				operand1, _ := strconv.Atoi(match[1])
				operand2, _ := strconv.Atoi(match[2])

				result += operand1 * operand2

			}
		} else {
			fmt.Println("No matches found")
		}
	}
	fmt.Println("Result: ", result)
}
