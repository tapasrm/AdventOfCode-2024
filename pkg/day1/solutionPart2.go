package day1

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

	distanceMap := make(map[int]int)
	var right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		leftVal, err := strconv.Atoi(parts[0])
		rightVal, err := strconv.Atoi(parts[1])

		if err != nil {
			fmt.Println("Err ", err)
		}

		distanceMap[leftVal] = 0
		right = append(right, rightVal)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for i := 0; i < len(right); i++ {
		value, exists := distanceMap[right[i]]
		if exists {
			distanceMap[right[i]] = value + 1
		}
	}

	distanceSum := 0
	for key, value := range distanceMap {
		distanceSum += key * value
	}
	print(distanceSum)
}
