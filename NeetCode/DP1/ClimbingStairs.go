package dp1

import "math"

func ClimbingStairsTopDown(n int) int {
	cache := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = -1
	}
	var dfs func(sum int) int
	dfs = func(sum int) int {
		if sum > n {
			return 0
		} else if sum == n {
			return 1
		}

		if cache[sum] != -1 {
			return cache[sum]
		}

		cache[sum] = dfs(sum+1) + dfs(sum+2)
		return cache[sum]
	}

	return dfs(0)
}

func ClimbingStairsBottomUp(n int) int {
	recentInt := [2]int{1, 1}
	for i := n - 2; i >= 0; i-- {
		temp := recentInt[0]
		recentInt[0] = recentInt[0] + recentInt[1]
		recentInt[1] = temp
	}
	return recentInt[0]
}

func ClimbingStairsMath(n int) int {
	// https://en.wikipedia.org/wiki/Fibonacci_sequence#Closed-form_expression
	// fn = (phi^n-psi^n)/sqrt(5); phi = (1+sqrt(5))/2; psi=(1-sqrt(5))/2

	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	psi := (1 - sqrt5) / 2
	n++

	return int((math.Pow(phi, float64(n)) - math.Pow(psi, float64(n))) / sqrt5)
}
