package main

import (
	"fmt"
	"time"
)

func main() {
	// var intArrr [3]int32 = [3]int32{1, 2, 3}
	// fmt.Println(intArrr[0])

	// intArr2 := [3]int32{2, 3, 4}
	// fmt.Println(intArr2[1])

	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Println(myMap)

	// var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
	// fmt.Println(myMap2)

	// for name := range myMap2 {
	// 	fmt.Printf("Name: %v\n", name)
	// }

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation %v\n", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
