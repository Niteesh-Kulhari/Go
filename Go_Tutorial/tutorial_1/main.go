package main

import (
	"errors"
	"fmt"
)

func main() {

	var intNum int16 = 1
	fmt.Println(intNum)

	myVar := "test"
	fmt.Println(myVar)

	printMe("PrintME function")
	fmt.Println(additon(2, 3))

	var result, remainder, err = division(2, 0)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of division is %v", result)
	} else {
		fmt.Printf("The result of division is %v and remainder after division is %v ", result, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func additon(number1 int, number2 int) int {
	var ans int = number1 + number2
	return ans
}

func division(number1 int, number2 int) (int, int, error) {
	var err error
	if number2 == 0 {
		err = errors.New("cannot divide by Zero")
		return 0, 0, err
	}
	var res int = number1 / number2
	var remainder int = number1 % number2

	return res, remainder, err
}
