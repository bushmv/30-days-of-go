package main

func knapsackV2(weights []int, values []int, limit int) int {
	return rknapsack(weights, values, limit, 0)
}

func rknapsack(w []int, v []int, limit int, i int) int {
	if i == len(w) {
		return 0
	}
	notPick := rknapsack(w, v, limit, i+1)
	if w[i] > limit {
		return notPick
	}
	pick := v[i] + rknapsack(w, v, limit-w[i], i+1)
	return max(pick, notPick)
}
