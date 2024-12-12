package heappq_test

import (
	"testing"

	heappq "github.com/Binit-Dhakal/LeetCode-Go/HeapPQ"
)

func TestKthLargestElement(t *testing.T) {
	t.Run("Ex-1", func(t *testing.T) {
		obj := heappq.Constructor(3, []int{4, 5, 8, 2})
		nums := []int{3, 5, 10, 9, 4}
		res := []int{4, 5, 5, 8, 8}

		for i, num := range nums {
			got := obj.Add(num)
			want := res[i]
			if want != got {
				t.Errorf("Got: %v, want: %v", got, want)
			}
		}
	})
	t.Run("Ex-2", func(t *testing.T) {
		obj := heappq.Constructor(4, []int{7, 7, 7, 7, 8, 3})
		nums := []int{2, 10, 9, 9}
		res := []int{7, 7, 7, 8}
		for i, num := range nums {
			got := obj.Add(num)

			want := res[i]
			if want != got {
				t.Errorf("Got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("Ex-3 Negative numbers", func(t *testing.T) {
		obj := heappq.Constructor(4, []int{-1, -2, -3, 0, -10})
		nums := []int{-1, -1, -9, -10, -11}
		res := []int{-2, -1, -1, -1, -1}
		for i, num := range nums {
			got := obj.Add(num)

			want := res[i]
			if want != got {
				t.Errorf("Got: %v, want: %v", got, want)
			}
		}
	})
}
