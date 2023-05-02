package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	mas := []string{"dog", "cat", "bird", "seal", "cow", "platypus"}
	var goSlice []string

	for {
		for i, el := range mas {
			go Append(&goSlice, fmt.Sprintf("line %d,", i+1)+" goroutine 1, animal \""+el+"\"")
			go Append(&goSlice, fmt.Sprintf("line %d,", i+1)+" goroutine 2, animal \""+el+"\"")
			go Append(&goSlice, fmt.Sprintf("line %d,", i+1)+" goroutine 3, animal \""+el+"\"")
		}

		time.Sleep(time.Millisecond * 1)
		//Иногда не все горутины запускаются, в таком случае мы повторяем цикл до тех пор пока не сработают все горутины
		if len(goSlice) < len(mas)*3 {
			continue
		}

		for i := 1; i <= len(mas); i++ {
			for j := 0; j < len(goSlice); j++ {
				if strings.Contains(goSlice[j], fmt.Sprintf("line %d", i)) {
					fmt.Println(goSlice[j])
					j = len(goSlice)
				}
			}
		}
		fmt.Println("done")

		break
	}
}

func Append(goSlice *[]string, text string) {
	*goSlice = append(*goSlice, text)
}

/*line 1, goroutine 1, animal "dog"
line 2, goroutine 2, animal "cat"
line 3, goroutine 1, animal "seal"
line 4, goroutine 3, animal "bird"
...
done*/

/*for i, el := range mas {
	for j := 1; j <= 3; j++ {
		go func() {
			strCh <- fmt.Sprintf("line %d,", i+1) + fmt.Sprintf(" goroutine %d, animal \"", j) + el + "\""
		}()
		fmt.Println(<-strCh)
	}
}
time.Sleep(1 * time.Second)*/

/*for i, el := range mas {
	go func() {
		strCh <- fmt.Sprintf("line %d,", i+1) + " goroutine 1, animal \"" + el + "\""
		mutex.Lock()
	}()
	go func() {
		strCh <- fmt.Sprintf("line %d,", i+1) + " goroutine 2, animal \"" + el + "\""
		mutex.Lock()
	}()
	go func() {
		strCh <- fmt.Sprintf("line %d,", i+1) + " goroutine 3, animal \"" + el + "\""
		mutex.Lock()
	}()
	fmt.Println(<-strCh)
	mutex.Unlock()
}*/

/*go func() {
	fmt.Println(fmt.Sprintf("line %d,", i+1) + " goroutine 1, animal \"" + el + "\"")
}()
go func() {
	fmt.Println(fmt.Sprintf("line %d,", i+1) + " goroutine 2, animal \"" + el + "\"")
}()
go func() {
	fmt.Println(fmt.Sprintf("line %d,", i+1) + " goroutine 3, animal \"" + el + "\"")
}()*/
