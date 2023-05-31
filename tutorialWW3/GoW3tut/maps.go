package main

import "fmt"

func main() {
	/*var a = map[string]string{"brand": "Ford", "model": "Mustang", "age": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)*/
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["age"] = "1964"

	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4
	/*delete(a, "brand")
	delete(b, "Bergen")
	fmt.Printf("a\t%v\nb\t%v\n", a, b)
	fmt.Println(b["Oslo"])*/

	val1, ok1 := a["brand"]
	val2, ok2 := a["color"]
	val3, ok3 := b["Oslo"]
	_, ok4 := a["model"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
}
