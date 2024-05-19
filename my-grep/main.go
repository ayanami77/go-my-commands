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
	Content string
	Number  int
}

func main() {
	isLineNumberCountingEnabled := flag.Bool("n", false, "Count lines")
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
	currentLineNumber := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), targetString) {
			matchedLines = append(matchedLines, Line{Content: scanner.Text(), Number: currentLineNumber})
		}
		currentLineNumber++
	}

	if *isLineNumberCountingEnabled {
		listLinesWithLineNumber(matchedLines, targetString)
		return
	}

	listLines(matchedLines, targetString)
}

func listLines(lines []Line, target string) {
	for _, line := range lines {
		fmt.Println(reconstructLine(line.Content, target))
	}
}

func listLinesWithLineNumber(lines []Line, target string) {
	for _, line := range lines {
		fmt.Printf("%d: %s\n", line.Number, reconstructLine(line.Content, target))
	}
}

func reconstructLine(line string, target string) string {
	green := color.New(color.FgGreen).SprintFunc()
	return strings.ReplaceAll(line, target, green(target))
}
