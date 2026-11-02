package main

import "fmt"

func main() {
	cb := NewCircularBuffer[int](5)
	cb.Put(1)
	cb.Put(2)
	cb.Put(3)
	cb.Put(4)
	_, _ = cb.Get()
	_, _ = cb.Get()
	_, _ = cb.Get()
	cb.Put(5)
	cb.Put(6)
	fmt.Println(cb.Dump())
	fmt.Println(cb.Values())
}
