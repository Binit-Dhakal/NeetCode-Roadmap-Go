package stack

import "fmt"

func LargestRectInHistogramBruteForce(heights []int) int {
	largestArea := 0
	for ptr := 0; ptr < len(heights); ptr++ {
		height := heights[ptr]

		leftIdx := ptr - 1
		for leftIdx >= 0 && heights[leftIdx] >= height {
			leftIdx--
		}

		rightIdx := ptr + 1
		for rightIdx < len(heights) && heights[rightIdx] >= height {
			rightIdx++
		}

		area := height * (rightIdx - leftIdx - 1)
		if area > largestArea {
			largestArea = area
		}
	}
	return largestArea
}

func LargestRectInHistogramStack(heights []int) int {
	heightStack := make([][2]int, 0) // (index, height)
	maxArea := 0

	for idx, height := range heights {
		start := idx
		for len(heightStack) > 0 && heightStack[len(heightStack)-1][1] > height {
			pair := heightStack[len(heightStack)-1]
			heightStack = heightStack[:len(heightStack)-1]
			maxArea = max(pair[1]*(idx-pair[0]), maxArea)
			start = pair[0]
		}
		heightStack = append(heightStack, [2]int{start, height})
		fmt.Println(heightStack, maxArea)
	}

	for _, pair := range heightStack {
		maxArea = max(maxArea, pair[1]*(len(heights)-pair[0]))
	}

	return maxArea
}
