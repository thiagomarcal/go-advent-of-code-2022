package exercise5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

type Stack struct {
	elements []string
}

func (s *Stack) add(element string) {
	s.elements = append(s.elements, element)
}

func (s *Stack) peek() string {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *Stack) pop() string {
	var popElement string
	if len(s.elements) > 0 {
		lastIndex := len(s.elements) - 1
		popElement = s.elements[lastIndex]
		s.elements = s.elements[:lastIndex]
	}
	return popElement
}

func removeEmptyString(strSlice []string) []string {
	newStringSlice := make([]string, 0)
	for i := range strSlice {
		if strSlice[i] != "" {
			newStringSlice = append(newStringSlice, strSlice[i])
		}
	}

	return newStringSlice
}

func makeStackMap(header []string) map[int]*Stack {
	stackMap := make(map[int]*Stack)

	for i := len(header) - 1; i >= 0; i-- {
		if i == len(header)-1 {
			headerBegLine := removeEmptyString(strings.Split(header[i], " "))
			lastOneStack, _ := strconv.Atoi(headerBegLine[len(headerBegLine)-1])
			for i := 1; i <= lastOneStack; i++ {
				newStack := new(Stack)
				stackMap[i] = newStack
			}
		} else {
			stackLine := strings.Split(header[i], " ")
			for i := 0; i <= len(stackLine)-1; i++ {
				stack, _ := stackMap[i+1]
				if stackLine[i] != "[?]" {
					stack.add(stackLine[i])
				}
			}
		}
	}
	return stackMap
}

func extractInfo(e Resolution) ([]string, []string) {
	input := e.readInput()

	defer func() {
		if err := input.Close(); err != nil {
			log.Fatalf("Problems closing the file")
		}
	}()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	header := make([]string, 0)
	body := make([]string, 0)

	lineType := "h"

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) == 0 {
			lineType = "b"
			continue
		}

		if lineType == "h" {
			header = append(header, line)
		} else {
			body = append(body, line)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Problems capturing input line")
	}
	return header, body
}

func (e Resolution) PartOne() {
	header, body := extractInfo(e)
	stackMap := makeStackMap(header)

	for _, line := range body {

		re := regexp.MustCompile("\\d{1,2}")
		result := re.FindAllString(line, -1)

		move, _ := strconv.Atoi(result[0])
		from, _ := strconv.Atoi(result[1])
		to, _ := strconv.Atoi(result[2])

		for i := 0; i < move; i++ {
			element := stackMap[from].pop()
			stackMap[to].add(element)
		}

	}

	result := ""
	for i := 1; i <= 9; i++ {
		re := regexp.MustCompile("[A-Z]")
		result += re.FindString(stackMap[i].peek())
	}

	fmt.Printf("Exercise5 part1: %s\n", result)

}

func (e Resolution) PartTwo() {
	header, body := extractInfo(e)
	stackMap := makeStackMap(header)

	for _, line := range body {

		re := regexp.MustCompile("\\d{1,2}")
		result := re.FindAllString(line, -1)

		move, _ := strconv.Atoi(result[0])
		from, _ := strconv.Atoi(result[1])
		to, _ := strconv.Atoi(result[2])

		if move == 1 {
			element := stackMap[from].pop()
			stackMap[to].add(element)
		} else {
			moves := make([]string, 0)

			for i := 0; i < move; i++ {
				element := stackMap[from].pop()
				moves = append(moves, element)
			}

			for i := len(moves) - 1; i >= 0; i-- {
				stackMap[to].add(moves[i])
			}
		}

	}

	result := ""
	for i := 1; i <= 9; i++ {
		re := regexp.MustCompile("[A-Z]")
		result += re.FindString(stackMap[i].peek())
	}

	fmt.Printf("Exercise5 part2: %s\n", result)
}
