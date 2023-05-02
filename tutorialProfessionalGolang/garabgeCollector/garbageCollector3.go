package main

import "runtime"

func main() {
	var N = 27000000
	newMap := make(map[int]int)
	for i := 0; i < N; i++ {
		value := int(i)
		newMap[value] = value
	}
	runtime.GC()
	_ = newMap
}
