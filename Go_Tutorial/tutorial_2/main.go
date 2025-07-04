package main

import "fmt"

func main() {
	var intArrr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArrr[0])

	intArr2 := [3]int32{2, 3, 4}
	fmt.Println(intArr2[1])

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
	fmt.Println(myMap2)

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}
}
