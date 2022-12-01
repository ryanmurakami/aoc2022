package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	elfCalories := aggCalories(lines)

	fmt.Println(elfCalories)

	largest := int64(0)
	topThree := make([]int64, 3)
	for _, val := range elfCalories {
		if val > largest {
			largest = val
		}
		topThree = addIfInTopThree(topThree, val)
	}

	fmt.Printf("Largest val is %d\n", largest)
	fmt.Printf("Top three are %v\n", topThree)
	fmt.Printf("Top three total is %v\n", topThree[0]+topThree[1]+topThree[2])

	return
}

func aggCalories(lines []string) []int64 {
	elves := make([]int64, 1)
	count := 0

	for _, val := range lines {
		if val == "" {
			count++
			elves = append(elves, 0)
		} else {
			intVal, err := strconv.ParseInt(val, 0, 64)
			if err != nil {
				panic(err)
			}
			elves[count] += intVal
		}
	}

	return elves
}

func addIfInTopThree(topThree []int64, val int64) []int64 {
	if val > topThree[0] {
		return []int64{val, topThree[0], topThree[1]}
	}

	if val > topThree[1] {
		return []int64{topThree[0], val, topThree[1]}
	}

	if val > topThree[2] {
		return []int64{topThree[0], topThree[1], val}
	}

	return topThree
}
