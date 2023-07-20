package exercise6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Resolution struct {
	FileInput string
}

func (e Resolution) readInput() *os.File {
	input, err := os.Open(e.FileInput)
	if err != nil {
		log.Fatalf("Problems reading input file %e", err)
	}

	return input
}

func extractInfo(e Resolution) string {
	input := e.readInput()

	defer func() {
		if err := input.Close(); err != nil {
			log.Fatalf("Problems closing the file")
		}
	}()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	output := ""

	for scanner.Scan() {
		line := scanner.Text()
		output += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Problems capturing input line")
	}
	return output
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func extractDistinctAt(simpleSample string, accumulateUntil int) (string, int) {
	resultSample := ""
	resultAt := 0

	for i, _ := range simpleSample {
		calcInit := 0

		if i-accumulateUntil < 0 {
			calcInit = 0
		} else {
			calcInit = i - accumulateUntil
		}

		sample := simpleSample[calcInit:i]

		dup := removeDuplicateStr(strings.Split(sample, ""))

		if len(dup) == accumulateUntil {
			resultSample = sample
			resultAt = i
			break
		}
	}
	return resultSample, resultAt
}

func (e Resolution) PartOne() {
	simpleSample := extractInfo(e)

	resultSample, resultAt := extractDistinctAt(simpleSample, 4)

	fmt.Printf("Exercise6 part1: %s - %d\n", resultSample, resultAt)

}

func (e Resolution) PartTwo() {
	simpleSample := extractInfo(e)

	resultSample, resultAt := extractDistinctAt(simpleSample, 14)

	fmt.Printf("Exercise6 part2: %s - %d\n", resultSample, resultAt)
}
