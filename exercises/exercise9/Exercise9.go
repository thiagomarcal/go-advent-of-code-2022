package exercise9

import (
	"log"
	"os"
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

func (e Resolution) PartOne() {

	//fmt.Printf("Resolution part1: %d\n", counter)

}

func (e Resolution) PartTwo() {

	//fmt.Printf("Resolution part2: %d\n", counter)
}
