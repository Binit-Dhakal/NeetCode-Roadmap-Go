package heappq

import "container/heap"

func TaskScheduler(tasks []byte, n int) int {
	count := make([][2]int, 26) // (repeatCount, nextTryIndex)
	for _, task := range tasks {
		count[task-'A'] = [2]int{count[task-'A'][0] + 1, 0}
	}

	totalCycle := 0
	maxInt := int(^uint(0) >> 1)
	for {
		allEmpty := true
		maxCntIdx := -1
		minCooldown := maxInt

		for i := 0; i < 26; i++ {
			c := count[i]

			if c[0] <= 0 {
				continue
			} else {
				allEmpty = false
			}

			if c[1] <= totalCycle {
				if maxCntIdx == -1 || count[maxCntIdx][0] <= c[0] {
					maxCntIdx = i
				}
			} else {
				minCooldown = min(c[1]-totalCycle, minCooldown)
			}
		}

		if allEmpty == true {
			break
		}
		if maxCntIdx != -1 {
			totalCycle++
			count[maxCntIdx] = [2]int{count[maxCntIdx][0] - 1, totalCycle + n}
		} else {
			// we are in idle state
			if minCooldown == maxInt {
				totalCycle++
			} else {
				totalCycle += minCooldown
			}
		}
	}
	return totalCycle
}

type TaskHeap []int

func (h TaskHeap) Len() int { return len(h) }

func (h TaskHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h TaskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TaskSchedulerHeap(tasks []byte, n int) int {
	h := TaskHeap{}
	count := [26]int{}

	for _, task := range tasks {
		count[task-'A']++
	}

	for _, c := range count {
		if c == 0 {
			continue
		}
		h = append(h, c)
	}

	queue := [][2]int{} // (count, nextTurnCount)
	heap.Init(&h)

	totalTime := 0
	for len(h) != 0 || len(queue) != 0 {
		totalTime++
		if len(h) == 0 {
			totalTime = queue[0][1]
		} else {
			x := heap.Pop(&h)
			if x.(int)-1 > 0 {
				queue = append(queue, [2]int{x.(int) - 1, totalTime + n})
			}
		}

		if len(queue) > 0 && queue[0][1] == totalTime {
			heap.Push(&h, queue[0][0])
			queue = queue[1:]
		}
	}
	return totalTime
}

func TaskSchedulerMath(tasks []byte, n int) int {
	count := [26]int{}
	for _, task := range tasks {
		count[task-'A']++
	}

	maxF := 0
	maxIdx := -1
	for idx, freq := range count {
		if freq > maxF {
			maxF = freq
			maxIdx = idx
		}
	}

	idleTime := (maxF - 1) * n
	for idx, freq := range count {
		if idx == maxIdx || freq == 0 {
			continue
		}

		idleTime -= min(maxF-1, freq)
	}

	return len(tasks) + max(0, idleTime)
}
