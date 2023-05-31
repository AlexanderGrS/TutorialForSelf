package main

import "fmt"

func rec(x int) (y int) {
	if x > 0 {
		y = x * rec(x-1)
	}
	if x <= 1 {
		return 1
	}
	return
}

func main() {
	fmt.Println(rec(11))
}
