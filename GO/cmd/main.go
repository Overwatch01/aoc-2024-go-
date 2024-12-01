package main

import (
	"fmt"
	"github.com/Overwatch01/aoc-2024-go/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadInput()
	if err != nil {
		log.Fatalf("Invalid input :: %v", err)
	}
	leftArr := make([]int, 0)
	rightArr := make([]int, 0)

	for _, line := range input {
		line = strings.TrimSpace(line)
		vals := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' '
		})

		leftInput, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatal("Invalid input")
		}
		leftArr = append(leftArr, leftInput)

		rightInput, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Fatal("Invalid input")
		}
		rightArr = append(rightArr, rightInput)
	}
	slices.Sort(leftArr)
	slices.Sort(rightArr)

	totalDistance := 0
	totalSimilarityScore := 0
	for i, _ := range leftArr {
		totalDistance += utils.GetDifference(leftArr[i], rightArr[i])
		similarityScore := utils.CountOccurrences(rightArr, leftArr[i])
		totalSimilarityScore += similarityScore * leftArr[i]

	}
	fmt.Printf("Total distance to be covered is %v \n", totalDistance)
	fmt.Printf("Total similary score is  %v \n", totalSimilarityScore)
}
