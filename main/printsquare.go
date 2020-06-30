package main

import str "strings"

func printsquare(width, height int) string {
	var square string
	i := 1

	for i <= height {
		if i == 1 || i == height {
			square += str.Repeat("😅", width) + "\n"
		} else {
			square += "😅" + str.Repeat("  ", width-2) + "😅" + "\n"
		}
		i++
	}
	return square
}
