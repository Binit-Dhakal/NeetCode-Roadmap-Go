package stack

import (
	"slices"
)

func CarFleet(target int, position []int, speed []int) int {
	n := len(position)
	pair := make([][2]int, n)

	for i := range n {
		pair[i] = [2]int{position[i], speed[i]}
	}

	slices.SortFunc(pair, func(car1 [2]int, car2 [2]int) int {
		return car1[0] - car2[0]
	})

	fleetCount := 0
	var currentTime float64 = 0

	for i := n - 1; i >= 0; i-- {
		t := float64(target-pair[i][0]) / float64(pair[i][1])
		if t > currentTime {
			fleetCount++
			currentTime = t
		}
	}

	return fleetCount
}
