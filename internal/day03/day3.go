package day03

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	r, _ := regexp.Compile("mul" + regexp.QuoteMeta("(") + "[0-9]{1,3},[0-9]{1,3}" + regexp.QuoteMeta(")"))
	instructions := r.FindAllString(input, -1)
	instructionsSum := 0

	for _, instruction := range instructions {
		instructionsSum += ComputeMultiplyInstruction(instruction)
	}

	return instructionsSum
}

func Part2(input string) int {
	r, _ := regexp.Compile("mul" + regexp.QuoteMeta("(") + "[0-9]{1,3},[0-9]{1,3}" + regexp.QuoteMeta(")") + "|do" + regexp.QuoteMeta("()") + "|don't" + regexp.QuoteMeta("()"))
	instructions := r.FindAllString(input, -1)
	instructionsSum := 0
	multiplyEnabled := true

	for _, instruction := range instructions {
		switch instruction {
		case "don't()":
			multiplyEnabled = false
		case "do()":
			multiplyEnabled = true
		default:
			if multiplyEnabled {
				instructionsSum += ComputeMultiplyInstruction(instruction)
			}
		}
	}
	return instructionsSum
}

func ComputeMultiplyInstruction(instruction string) int {
	parsedInstruction := strings.Split(instruction[4:len(instruction)-1], ",")
	first, _ := strconv.Atoi(parsedInstruction[0])
	second, _ := strconv.Atoi(parsedInstruction[1])
	return first * second
}
