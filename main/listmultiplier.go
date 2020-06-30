package main

import (
	"math/rand"
)

func listmultiplier(num int) []int {
	var list = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}

	for j := range list {
		list[j] *= num
	}

	return list
}
