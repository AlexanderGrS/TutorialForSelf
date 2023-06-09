package main

import "runtime"

type data struct {
	i, j int
}

func main() {
	var N = 27000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()
	_ = structure[0]

}
