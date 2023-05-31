package main

import (
	"fmt"
	"math"
	"time"
)

func Hammer(n int) (answer uint) {
	compare := uint(math.Pow(30, 12))
	for i := uint(1); n != 0; i++ {
		if compare%i == 0 {
			n--
		}
		if n == 0 {
			answer = uint(i)
		}
	}
	return answer
}

func main() {
	checkSpeed := time.Now()
	fmt.Println(Hammer(1150))
	fmt.Println(time.Since(checkSpeed))
}
