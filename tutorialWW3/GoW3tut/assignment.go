package main

import "fmt"

func main() {
	var x = 5
	fmt.Println(x % 3)
	fmt.Println(x >> 3)
	fmt.Println(x << 3)
	fmt.Println(x & 3)
	fmt.Println(((x | 3) > 2))
	fmt.Println(x ^ 3)
}
