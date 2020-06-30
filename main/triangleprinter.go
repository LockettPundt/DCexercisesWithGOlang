package main

import str "strings"

func triangleprinter(num int) string {
	var triangle string
	i := 1

	for i < num {
		if i == 1 {
			triangle += str.Repeat(" ", num) + "*" + "\n"
		} else {
			triangle += str.Repeat(" ", num) + str.Repeat("*", i) + "\n"
		}
		i += 2
		num--
	}
	return triangle
}
