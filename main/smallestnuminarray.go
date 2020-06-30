package main

import "fmt"

func smallestnuminarray(arr []int) int {

	temp := arr[0]

	for index, element := range arr {
		fmt.Printf("Here is the index %d and here is the num %d \n", index, element)
		if temp > element {
			temp = element
		}
	}
	return temp
}
