package day1

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

	var left []int
	var right []int
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		leftVal, err := strconv.Atoi(parts[0])
		rightVal, err := strconv.Atoi(parts[1])

		if err != nil {
			fmt.Println("Err: ", err)
		}

		insertToSortedArray(&left, leftVal)
		insertToSortedArray(&right, rightVal)

		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	distanceSum := 0
	for i := 0; i < count; i++ {
		absDiff := left[i] - right[i]
		if absDiff < 0 {
			absDiff = -absDiff
		}
		distanceSum += absDiff
	}
	print(distanceSum)
}

func insertToSortedArray(arr *[]int, value int) {
	array := *arr

	left, right := 0, len(array)
	for left < right {
		mid := left + (right-left)/2
		if array[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}

	*arr = append(array[:left], append([]int{value}, array[left:]...)...)
}
