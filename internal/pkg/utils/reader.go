package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadIntegersInFile(filePath string) []int {
	file, _ := os.Open(filePath)

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, int(integer))
	}

	return lines
}

func ReadMultipleIntegersPerLineInFile(filePath string) [][]int {
	file, _ := os.Open(filePath)

	var lines [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var intLine []int

		for _, part := range parts {
			if part != "" {
				integer, _ := strconv.ParseInt(part, 10, 64)
				intLine = append(intLine, int(integer))
			}
		}

		lines = append(lines, intLine)
	}

	return lines
}

func ReadStringsInFile(filePath string) []string {
	file, _ := os.Open(filePath)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
