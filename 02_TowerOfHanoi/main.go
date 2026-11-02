package main

import "fmt"

func toh(n int, from int, to int, temp int) {
	if n == 1 {
		fmt.Printf("%d -> %d\n", from, to)
		return
	}
	toh(n-1, from, temp, to)
	toh(1, from, to, temp)
	toh(n-1, temp, to, from)
}

func main() {
	toh(3, 1, 3, 2)
}
