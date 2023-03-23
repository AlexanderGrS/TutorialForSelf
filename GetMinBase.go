package main

import "fmt"

func GetMinBase(n uint64) uint64 {
	var base uint64 = 1
	for ; n != 0; base++ {
		for num := n; num > 0; num /= base {
			switch {
			case num%base != 1:
				num = 0
			case num == 1:
				n = 0
			}
		}
		if base > 10000000 {
			return n - 1
		}
	}
	return base - 1
}

func main() {
	fmt.Println(GetMinBase(125002500050001))
}
