package main

func knapsackV1(weights []int, values []int, target int) int {
	n := len(weights)
	best := 0
	for mask := 0; mask < (1 << n); mask++ {
		curW, curV := 0, 0
		for i := range n {
			if mask&(1<<i) > 0 {
				curW += weights[i]
				curV += values[i]
			}
		}
		if curW <= target && curV > best {
			best = curV
		}
	}
	return best
}
