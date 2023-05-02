package main

import "runtime"

func main() {
	var n = 27000000
	var newMap = make(map[int]*int)
	for i := 0; i < n; i++ {
		value := int(i)
		newMap[value] = &value
	}
	runtime.GC()
	_ = newMap[0]
}
