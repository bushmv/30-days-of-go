package main

func knapsackV3(weights []int, values []int, limit int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range len(dp) {
		dp[i] = make([]int, limit+1)
	}
	for y := 1; y <= n; y++ {
		for x := 1; x <= limit; x++ {
			noPick := dp[y-1][x]
			if x < weights[y-1] {
				dp[y][x] = noPick
			} else {
				pick := dp[y-1][x-weights[y-1]] + values[y-1]
				dp[y][x] = max(noPick, pick)
			}
		}
	}
	return dp[n][limit]
}
