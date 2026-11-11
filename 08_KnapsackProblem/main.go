package main

import "fmt"

func main() {
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	limit := 7
	fmt.Println(knapsackV1(weights, values, limit))
	fmt.Println(knapsackV2(weights, values, limit))
	fmt.Println(knapsackV3(weights, values, limit))
	fmt.Println(knapsackV4(weights, values, limit))
}
