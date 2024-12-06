package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fHandle, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fHandle.Close()

	scanner := bufio.NewScanner(fHandle)

	var leftList []int
	var rightList []int

	// Build the right and left list from input
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		if len(words) == 0 {
			continue
		}

		leftList = appendValToList(leftList, words[0])
		rightList = appendValToList(rightList, words[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort the imported lists
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	var diffArr []int

	// Build the difference vector
	for idx, _ := range leftList {
		diffArr = append(diffArr, absDiff(leftList[idx], rightList[idx]))
	}

	fmt.Println(sumList(diffArr))
}

func appendValToList(list []int, val string) []int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
	}
	return append(list, intVal)
}

func absDiff(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func sumList(list []int) int {
	sum := 0
	for _, elem := range list {
		sum += elem
	}

	return sum
}
