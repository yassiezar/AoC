package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

	var nSafeSequences uint = 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) == 0 {
			break
		}

		if !isSafeSequence(words) {
			continue
		}
		nSafeSequences += 1
	}
	fmt.Printf("Safe sequences: %d", nSafeSequences)
}

func isSafeSequence(words []string) bool {
	var prevDiff int = 0
	var prevVal int = 0
	for idx, word := range words {
		num, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal("Conversion error ", err)
			break
		}
		// Need 2 vals for difference
		if idx == 0 {
			prevVal = num
			continue
		}

		// Check difference. If the product of the current and prev difference is positive,
		// they are both positive or negative (i.e. safe)
		diff := num - prevVal
		if idx == 1 && diff == 0 {
			prevDiff = diff
			return false
		} else if math.Abs(float64(diff)) > 3 {
			return false
		} else if idx > 1 && diff*prevDiff <= 0 {
			return false
		}
		prevDiff = diff
		prevVal = num
	}
	return true
}
