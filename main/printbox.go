package main

import str "strings"

func printbox(width, height int) string {
	var box string
	row := str.Repeat("*", width)
	i := 1
	// quasi while loop since GO doesn't technically have a while loop
	for i < height {
		box += str.Repeat(row, 1) + "\n"
		i++
	}
	return box
}
