package exercise4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (e Resolution) PartOne() {
	input := e.readInput()

	defer func() {
		if err := input.Close(); err != nil {
			log.Fatalf("Problems closing the file")
		}
	}()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	totalCommons := 0

	for scanner.Scan() {

		commons := make([]string, 0)

		line := scanner.Text()
		ranges := strings.Split(line, ",")

		range1 := ranges[0]
		range2 := ranges[1]

		comp1 := strings.Split(range1, "-")
		comp2 := strings.Split(range2, "-")

		comp1Ini, _ := strconv.Atoi(comp1[0])
		comp1End, _ := strconv.Atoi(comp1[1])

		comp2Ini, _ := strconv.Atoi(comp2[0])
		comp2End, _ := strconv.Atoi(comp2[1])

		if comp1Ini <= comp2Ini && comp2End <= comp1End {
			commons = append(commons, range1)
		}

		if comp2Ini <= comp1Ini && comp1End <= comp2End {
			commons = append(commons, range2)
		}

		// removing duplicates
		distinctCommons := removeDuplicateStr(commons)
		totalCommons += len(distinctCommons)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Problems capturing input line")
	}

	fmt.Printf("Resolution part1: %d\n", totalCommons)
}

func (e Resolution) PartTwo() {
	input := e.readInput()

	defer func() {
		if err := input.Close(); err != nil {
			log.Fatalf("Problems closing the file")
		}
	}()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	totalCommons := 0

	for scanner.Scan() {

		commons := make([]string, 0)

		line := scanner.Text()
		ranges := strings.Split(line, ",")

		range1 := ranges[0]
		range2 := ranges[1]

		comp1 := strings.Split(range1, "-")
		comp2 := strings.Split(range2, "-")

		comp1Ini, _ := strconv.Atoi(comp1[0])
		comp1End, _ := strconv.Atoi(comp1[1])

		comp2Ini, _ := strconv.Atoi(comp2[0])
		comp2End, _ := strconv.Atoi(comp2[1])

		if comp1End >= comp2Ini && comp2End >= comp1End {
			fmt.Printf("line cond1: %s\n", line)
			commons = append(commons, line)
		} else {
			if comp2End >= comp1Ini && comp1End >= comp2End {
				fmt.Printf("line cond2: %s\n", line)
				commons = append(commons, line)
			} else {
				fmt.Printf("not overlapping: %s\n", line)
			}
		}

		// removing duplicates
		distinctCommons := removeDuplicateStr(commons)
		totalCommons += len(distinctCommons)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Problems capturing input line")
	}

	fmt.Printf("Resolution part2: %d\n", totalCommons)
}
