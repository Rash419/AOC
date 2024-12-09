package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/Rash419/AOC/input"
)

const (
	Increasing = iota
	Decreasing
	Stagnant
)

func main() {
	lines, err := input.ReadFileByLine("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt with error[%s]", err.Error())
	}

	report := make([][]int, len(lines))
	parseReport(report, lines)

	safeReports := 0
	for i := 0; i < len(report); i++ {
		// check first two elements for to check if report is increasing or decreasing
		if len(report[i]) < 2 {
			continue
		}

		var reportType int
		if report[i][0] < report[i][1] {
			reportType = Increasing
		} else if report[i][0] > report[i][1] {
			reportType = Decreasing
		} else {
			continue
		}

		unsafeReport := false
		for j := 0; j < len(report[i])-1; j++ {
			diff := 0
			switch reportType {
			case Increasing:
				diff = report[i][j+1] - report[i][j]
			case Decreasing:
				diff = report[i][j] - report[i][j+1]
			default:
			}
			if diff <= 0 || diff > 3 {
				unsafeReport = true
				break
			}
		}

		if !unsafeReport {
			safeReports++
		}
	}

	log.Printf("SafeReports: %d", safeReports)
}

func parseReport(fullReport [][]int, lines []string) {
	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		report := make([]int, len(splitLine))
		for j := 0; j < len(splitLine); j++ {
			intEle, err := strconv.Atoi(splitLine[j])
			if err != nil {
				log.Fatalf("Failed to parseReport with error[%s]", err.Error())
			}
			report[j] = intEle
		}
		fullReport[i] = report
	}
}
