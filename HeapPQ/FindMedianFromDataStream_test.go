package heappq_test

import (
	"reflect"
	"testing"

	heappq "github.com/Binit-Dhakal/LeetCode-Go/HeapPQ"
)

func TestMedianFromDataStream(t *testing.T) {
	t.Run("Custom example", func(t *testing.T) {
		obj := heappq.ConstructMedian()

		addNumbers := []int{1, 3, 5, 4, 2, 6, 100, 50, 75, 6}
		medianEachStep := []float64{1.0, 2.0, 3.0, 3.5, 3.0, 3.5, 4.0, 4.5, 5, 5.5}

		for i, num := range addNumbers {
			obj.AddNum(num)
			got := obj.FindMedian()
			if medianEachStep[i] != got {
				t.Errorf("Got: %v, want: %v", got, medianEachStep[i])
			}
		}

		want := []int{1, 2, 3, 4, 5, 6, 6, 50, 75, 100}
		got := obj.ShowNums()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Leetcode example", func(t *testing.T) {
		obj := heappq.ConstructMedian()
		addNumbers := []int{12, 10, 13, 11, 5, 15, 1, 11, 6, 17, 14, 8, 17, 6, 4, 16, 8, 10, 2, 12, 0}
		for _, num := range addNumbers {
			obj.AddNum(num)
		}

		want := []int{0, 1, 2, 4, 5, 6, 6, 8, 8, 10, 10, 11, 11, 12, 12, 13, 14, 15, 16, 17, 17}
		got := obj.ShowNums()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})

	t.Run("Heap solution", func(t *testing.T) {
		obj := heappq.ConstructMedianHeap()
		addNumbers := []int{1, 3, 5, 4, 2, 6, 100, 50, 75, 6}
		medianEachStep := []float64{1.0, 2.0, 3.0, 3.5, 3.0, 3.5, 4.0, 4.5, 5, 5.5}

		for i, num := range addNumbers {
			obj.AddNum(num)
			got := obj.FindMedian()
			if medianEachStep[i] != got {
				t.Errorf("Got: %v, want: %v", got, medianEachStep[i])
			}
		}
	})
}
