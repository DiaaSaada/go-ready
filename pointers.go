package main

import (
	"fmt"
)

func main() {

	var arr1 = [5]float32{1, 2, 3, 4, 5}
	var res = sqrArrByValue(arr1)
	fmt.Println(arr1) //[1 2 3 4 5]  arr1 not changed => a copy of arr1 was passed not the refrence of arr2
	// no side effect but double memory
	fmt.Println(res) //[1 4 9 16 25]

	//--- by refrence (pointer)
	var res2 = sqrArrByPointer(&arr1)
	fmt.Println(arr1)  //[1 4 9 16 25] changed since passed by ref this time
	fmt.Println(*res2) //[1 4 9 16 25]
}

func sqrArrByValue(arr2 [5]float32) [5]float32 {
	fmt.Println("arr2 address", &arr2)
	for i := range arr2 {
		fmt.Println("arr2 address", i, arr2[i])
		arr2[i] = arr2[i] * arr2[i]
	}

	return arr2
}

func sqrArrByPointer(arr2 *[5]float32) *[5]float32 {
	fmt.Println("arr2 address", &arr2)
	for i := range arr2 {
		fmt.Println("arr2 address", i, arr2[i])
		arr2[i] = arr2[i] * arr2[i]
	}

	return arr2
}
