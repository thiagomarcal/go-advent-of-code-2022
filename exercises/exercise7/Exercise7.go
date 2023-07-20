package exercise7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type File struct {
	filename string
	size     int
}

type Dir struct {
	isRoot      bool
	name        string
	files       []File
	children    map[string]*Dir
	parent      *Dir
	currentSize int
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

func (e Resolution) mapDirectories() (*Dir, []*Dir) {

	input := e.readInput()

	defer func() {
		if err := input.Close(); err != nil {
			log.Fatalf("Problems closing the file")
		}
	}()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	currentDir := new(Dir)

	var rootDir *Dir

	allDirs := make([]*Dir, 0)
	for scanner.Scan() {
		line := scanner.Text()

		commandRegex := regexp.MustCompile("\\S+")

		command := commandRegex.FindAllString(line, -1)

		if command[0] == "$" {
			switch command[1] {
			case "cd":

				if command[2] == ".." {
					if currentDir.parent != nil {
						currentDir = currentDir.parent
					}

				} else {
					var dir *Dir
					if value, exists := currentDir.children[command[2]]; exists {
						dir = value
					} else {

						dir = &Dir{
							name:        command[2],
							isRoot:      true,
							children:    make(map[string]*Dir, 0),
							files:       make([]File, 0),
							parent:      nil,
							currentSize: 0,
						}

						allDirs = append(allDirs, dir)

						rootDir = dir

					}
					currentDir = dir
				}

				break
			case "ls":
				break
			}
		} else if command[0] == "dir" {
			if _, exists := currentDir.children[command[1]]; !exists {
				childDir := &Dir{
					name:        command[1],
					isRoot:      false,
					children:    make(map[string]*Dir, 0),
					files:       make([]File, 0),
					parent:      currentDir,
					currentSize: 0,
				}
				allDirs = append(allDirs, childDir)
				currentDir.children[command[1]] = childDir
			}
		} else {
			fileSizeRegex := regexp.MustCompile("\\d+")
			if fileSizeRegex.MatchString(line) {
				fileSize, _ := strconv.Atoi(fileSizeRegex.FindString(line))
				lineInfo := strings.Split(line, " ")
				file := File{filename: lineInfo[1], size: fileSize}
				currentDir.files = append(currentDir.files, file)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Problems capturing input line")
	}

	// calc size for each dir
	for _, dir := range allDirs {
		totalInDir := 0

		for _, file := range dir.files {
			totalInDir += file.size
		}

		dir.currentSize += totalInDir

		for dir.parent != nil {
			dir.parent.currentSize += totalInDir
			dir = dir.parent
		}
	}

	return rootDir, allDirs

}

func (e Resolution) PartOne() {

	_, allDirs := e.mapDirectories()

	total := 0
	for _, dir := range allDirs {
		maxDirSize := 100000
		if dir.currentSize <= maxDirSize {
			total += dir.currentSize
		}
	}

	fmt.Printf("Resolution part1: %d\n", total)

}

func (e Resolution) PartTwo() {

	rootDir, allDirs := e.mapDirectories()

	totalSpace := 70000000
	neededFreeSpace := 30000000

	currentFreeSpace := neededFreeSpace - (totalSpace - rootDir.currentSize)

	dirToDeleteCandidates := make([]*Dir, 0)
	for _, dir := range allDirs {
		if dir.currentSize >= currentFreeSpace {
			dirToDeleteCandidates = append(dirToDeleteCandidates, dir)
		}
	}

	var smallDirCandidateForDeletion *Dir

	minDir := math.MaxInt32
	for _, dir := range dirToDeleteCandidates {
		if dir.currentSize <= minDir {
			minDir = dir.currentSize
			smallDirCandidateForDeletion = dir
		}
	}

	fmt.Printf("Resolution part2: %d\n", smallDirCandidateForDeletion.currentSize)

}
