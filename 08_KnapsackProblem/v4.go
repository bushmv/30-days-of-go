package main

func knapsackV4(weights []int, values []int, limit int) int {
	dp := make([]int, limit+1)
	for y := range weights {
		for x := limit; x >= weights[y]; x-- {
			if v := dp[x-weights[y]] + values[y]; v > dp[x] {
				dp[x] = v
			}
		}
	}
	return dp[limit]
}
