package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Line struct {
	Content  string
	Position int
}

func main() {
	isLineCountingEnabled := flag.Bool("n", false, "Count lines")
	flag.Parse()

	targetString := flag.Arg(0)
	targetFile := flag.Arg(1)

	file, err := os.Open(targetFile)
	if err != nil {
		panic("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matchedLines := []Line{}
	currentPositions := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), targetString) {
			matchedLines = append(matchedLines, Line{Content: scanner.Text(), Position: currentPositions})
		}
		currentPositions++
	}

	if *isLineCountingEnabled {
		listLinesWithPosition(matchedLines, targetString)
		return
	}

	listLines(matchedLines, targetString)
}

func listLines(lines []Line, target string) {
	for _, line := range lines {
		fmt.Println(reconstructLine(line.Content, target))
	}
}

func listLinesWithPosition(lines []Line, target string) {
	for _, line := range lines {
		fmt.Printf("%d: %s\n", line.Position, reconstructLine(line.Content, target))
	}
}

func reconstructLine(line string, target string) string {
	green := color.New(color.FgGreen).SprintFunc()
	return strings.ReplaceAll(line, target, green(target))
}
