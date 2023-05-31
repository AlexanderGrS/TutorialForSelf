package main

import "fmt"

func main() {
	/*var myslice1 = []int{}
	fmt.Printf("Type: %T, value: %v \n", len(myslice1), len(myslice1))
	fmt.Printf("Type: %T, value: %v \n", cap(myslice1), cap(myslice1))

	var myslice2 = []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Printf("Type: %T, value: %v \n", len(myslice2), len(myslice2))
	fmt.Printf("Type: %T, value: %v \n", cap(myslice2), cap(myslice2))
	fmt.Printf("Type: %T, value: %v \n", myslice2, myslice2)

	var arr1 = [7]int16{23, 45, 67, 41, 45}
	var myslice3 = arr1[2:4]

	fmt.Printf("Type: %T, value: %v \n", myslice3, myslice3)
	fmt.Printf("Type: %T, value: %v \n", cap(myslice3), cap(myslice3))
	fmt.Printf("Type: %T, value: %v \n", len(myslice3), len(myslice3))*/

	/*myslice4 := make([]int, 5, 10)
	fmt.Printf("mysclice4 = %v\n", myslice4)
	fmt.Printf("lenght = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	myslice5 := make([]int, 5)
	fmt.Printf("mysclice4 = %v\n", myslice5)
	fmt.Printf("lenght = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))*/

	/*mysclice6 := []int{1, 2, 3, 4}
	fmt.Printf("mysclice6 = %v\n", mysclice6)
	fmt.Printf("lenght = %v\n", len(mysclice6))
	fmt.Printf("capacity = %v\n", cap(mysclice6))

	mysclice6 = append(mysclice6, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("mysclice6 = %v\n", mysclice6)
	fmt.Printf("lenght = %v\n", len(mysclice6))
	fmt.Printf("capacity = %v\n", cap(mysclice6))

	mysclice6 = append(mysclice6, 11, 12)
	fmt.Printf("mysclice6 = %v\n", mysclice6)
	fmt.Printf("lenght = %v\n", len(mysclice6))
	fmt.Printf("capacity = %v\n", cap(mysclice6))*/

	/*mysclice7 := []int{1, 2, 3}
	mysclice8 := []int{4, 5, 6}
	mysclice8 = append(mysclice8, mysclice7...)
	mysclice9 := append(mysclice7, mysclice8...)

	fmt.Printf("mysclice6 = %v\n", mysclice9)
	fmt.Printf("lenght = %v\n", len(mysclice9))
	fmt.Printf("capacity = %v\n", cap(mysclice9))

	arr2 := [6]int{9, 10, 11, 12, 13, 14}
	mysclice10 := arr2[0:2]
	fmt.Printf("mysclice10 = %v\n", mysclice10)
	fmt.Printf("lenght = %v\n", len(mysclice10))
	fmt.Printf("capacity = %v\n", cap(mysclice10))

	fmt.Printf("mysclice6 = %v\n", arr2)
	fmt.Printf("lenght = %v\n", len(arr2))
	fmt.Printf("capacity = %v\n", cap(arr2))

	mysclice10 = append(mysclice10, 15, 16, 17, 18)
	fmt.Printf("mysclice10 = %v\n", mysclice10)
	fmt.Printf("lenght = %v\n", len(mysclice10))
	fmt.Printf("capacity = %v\n", cap(mysclice10))*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("lenght = %v\n", len(numbers))
	fmt.Printf("capacity = %v\n", cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]

	fmt.Printf("numberCopy = %v\n", neededNumbers)
	fmt.Printf("lenght = %v\n", len(neededNumbers))
	fmt.Printf("capacity = %v\n", cap(neededNumbers))

	var numberCopy = make([]int, len(neededNumbers))
	copy(numberCopy, neededNumbers)

	fmt.Printf("numberCopy = %v\n", numberCopy)
	fmt.Printf("lenght = %v\n", len(numberCopy))
	fmt.Printf("capacity = %v\n", cap(numberCopy))

}
