package exercise8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type TileVisibility struct {
	positionX int
	positionY int
	element   int

	visLeft  bool
	visRight bool
	visUp    bool
	visDown  bool
}

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

func extractGridSize(e Resolution) int {
	input := e.readInput()
	defer func() {
		if err := input.Close(); err != nil {
			log.Fatalf("Problems closing the file")
		}
	}()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	gridSize := 0

	for scanner.Scan() {
		gridSize++
	}

	return gridSize
}

func extractInfo(e Resolution, gridSize int) [][]TileVisibility {

	input := e.readInput()

	defer func() {
		if err := input.Close(); err != nil {
			log.Fatalf("Problems closing the file")
		}
	}()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	treeMatrix := make([][]TileVisibility, gridSize)

	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()

		treeMatrix[lineNumber] = make([]TileVisibility, len(line))

		for i := 0; i < len(line); i++ {
			tree, _ := strconv.Atoi(string(line[i]))
			treeMatrix[lineNumber][i] = TileVisibility{
				positionX: lineNumber,
				positionY: i,
				element:   tree,
				visLeft:   false,
				visRight:  false,
				visUp:     false,
				visDown:   false,
			}
		}

		lineNumber++
	}

	// UP
	for col := 0; col < gridSize; col++ {
		maxUp := -1
		for row := 0; row < gridSize; row++ {
			currentTree := &treeMatrix[row][col]
			if currentTree.element > maxUp {
				currentTree.visUp = true
				maxUp = currentTree.element
			}
		}
	}

	// DOWN
	for col := 0; col < gridSize; col++ {
		maxDown := -1
		for row := gridSize - 1; row >= 0; row-- {
			currentTree := &treeMatrix[row][col]
			if currentTree.element > maxDown {
				currentTree.visDown = true
				maxDown = currentTree.element
			}
		}
	}

	// LEFT
	for row := 0; row < gridSize; row++ {
		maxLeft := -1
		for col := 0; col < gridSize; col++ {
			currentTree := &treeMatrix[row][col]
			if currentTree.element > maxLeft {
				currentTree.visLeft = true
				maxLeft = currentTree.element
			}
		}
	}

	// RIGHT
	for row := 0; row < gridSize; row++ {
		maxRight := -1
		for col := gridSize - 1; col >= 0; col-- {
			currentTree := &treeMatrix[row][col]
			if currentTree.element > maxRight {
				currentTree.visRight = true
				maxRight = currentTree.element
			}
		}
	}

	return treeMatrix

}

func (e Resolution) PartOne() {

	gridSize := extractGridSize(e)

	treeMatrix := extractInfo(e, gridSize)

	counter := 0
	for row := 0; row < len(treeMatrix); row++ {
		for col := 0; col < len(treeMatrix[row]); col++ {
			currentTree := &treeMatrix[row][col]

			if currentTree.visUp {
				counter += 1
			} else if currentTree.visDown {
				counter += 1
			} else if currentTree.visLeft {
				counter += 1
			} else if currentTree.visRight {
				counter += 1
			}

		}
	}

	fmt.Printf("Resolution part1: %d\n", counter)

}

func (e Resolution) PartTwo() {

	gridSize := extractGridSize(e)

	treeMatrix := extractInfo(e, gridSize)

	highestScenicScore := 0
	for row := 0; row < len(treeMatrix); row++ {
		for col := 0; col < len(treeMatrix[row]); col++ {
			currentTree := &treeMatrix[row][col]

			upDistanceBlock := 0
			for i := row; i > 0; i-- {
				upDistanceBlock += 1
				if treeMatrix[i-1][col].element >= currentTree.element {
					break
				}
			}

			downDistanceBlock := 0
			for i := row; i+1 < gridSize; i++ {
				downDistanceBlock += 1
				if treeMatrix[i+1][col].element >= currentTree.element {
					break
				}
			}

			leftDistanceBlock := 0
			for i := col; i > 0; i-- {
				leftDistanceBlock += 1
				if treeMatrix[row][i-1].element >= currentTree.element {
					break
				}
			}

			rightDistanceBlock := 0
			for i := col; i+1 < gridSize; i++ {
				rightDistanceBlock += 1
				if treeMatrix[row][i+1].element >= currentTree.element {
					break
				}
			}

			score := upDistanceBlock * downDistanceBlock * leftDistanceBlock * rightDistanceBlock
			if score > highestScenicScore {
				highestScenicScore = score
			}
		}
	}

	fmt.Printf("Resolution part2: %d\n", highestScenicScore)

}
