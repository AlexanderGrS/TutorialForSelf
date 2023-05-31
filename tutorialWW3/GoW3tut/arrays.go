package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("Type: %T, value: %v \n", arr1, arr1)
	fmt.Printf("Type: %T, value: %v \n", arr2, arr2)

	var arr3 = [...]int64{3, 2, 1}
	arr4 := [...]int32{5, 4, 3, 2, 1}

	fmt.Printf("Type: %T, value: %v \n", arr3, arr3)
	fmt.Printf("Type: %T, value: %v \n", arr4, arr4)

	fmt.Printf("Type: %T, value: %v \n", arr3[1], arr3[1])

	arr3[1] = 50
	fmt.Printf("Type: %T, value: %v \n", arr3[1], arr3[1])

	arr5 := [5]int16{1: 40, 3: 120}
	fmt.Printf("Type: %T, value: %v \n", arr5, arr5)

	fmt.Printf("Type: %T, value: %v \n", len(arr5), len(arr5))
	fmt.Printf("Type: %T, value: %v \n", cap(arr5), cap(arr5))
}
