package main

import "fmt"

func returnChange2Num(intA int, intB int) (int, int) {
	if intA < intB {
		return intB, intA
	} else {
		return intA, intB
	}
}

// func main() {
// 	fmt.Println(returnChange2Num(returnChange2Num(1, 2)))
// 	// fmt.Println(returnChange2Num(1, 2), " d" ) // error
// }
