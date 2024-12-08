package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/Rash419/AOC/input"
)

func main() {
	lines, err := input.ReadFileByLine("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt with error[%s]", err.Error())
	}

	locationList1 := make([]int, len(lines))
	locationList2 := make([]int, len(lines))

	for i, line := range lines {
		splitLine := strings.Split(line, "   ")
		if len(splitLine) != 2 {
			break
		}
		location1, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatalf("Failed to convert string[%s] to integer with error[%s]", splitLine[0], err)
		}
		location2, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatalf("Failed to convert string[%s] to integer with error[%s]", splitLine[1], err)
		}

		locationList1[i] = location1
		locationList2[i] = location2
	}

	findDistance(locationList1, locationList2)
  findSimilarityScore(locationList1, locationList2)
}

// soltion for part-2
func findSimilarityScore(locationList1, locationList2 []int) {
	locationList2Map := make(map[int]int)
	for _, location := range locationList2 {
		locationList2Map[location]++
	}

	similarityScore := 0
	for _, location := range locationList1 {
    if val, exist := locationList2Map[location]; exist {
		  similarityScore += location * val
    }
	}

	log.Printf("SimilarityScore: %d", similarityScore)
}

// solution for part-1
func findDistance(locationList1, locationList2 []int) {
	sort.Ints(locationList1)
	sort.Ints(locationList2)

	distance := 0
	for i := 0; i < len(locationList1); i++ {
		diff := locationList1[i] - locationList2[i]
		if diff < 0 {
			diff *= -1
		}
		distance += diff
	}

	log.Printf("Distance: %d", distance)
}
