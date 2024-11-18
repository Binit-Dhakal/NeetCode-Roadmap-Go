package binarysearch

import (
	"slices"
)

func KokoBananaBruteForce(piles []int, hours int) int {
	maxRate := slices.Max(piles)

	for rate := 1; rate <= maxRate; rate++ {
		totalHour := 0
		for _, pile := range piles {
			totalHour += pile / rate
			if pile%rate != 0 {
				totalHour++
			}
		}

		if totalHour <= hours {
			return rate
		}

	}
	return maxRate
}

func KokoBananaBinarySearch(piles []int, hours int) int {
	maxRate := slices.Max(piles)

	leftPtr, rightPtr := 1, maxRate
	res := rightPtr

	for leftPtr <= rightPtr {
		rate := leftPtr + (rightPtr-leftPtr)/2
		totalHour := 0
		for _, pile := range piles {
			totalHour += pile / rate
			if pile%rate != 0 {
				totalHour++
			}
		}

		if totalHour <= hours {
			res = rate
			rightPtr = rate - 1
		} else {
			leftPtr = rate + 1
		}
	}

	return res
}
