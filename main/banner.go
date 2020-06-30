package main

import str "strings"

func banner(s string) string {
	length := len(s)
	var banner string
	for i := 1; i < 4; i++ {
		if i == 1 {
			banner += str.Repeat("*", length+2) + "\n"
		} else if i == 3 {
			banner += str.Repeat("*", length+2) + "\n"
		} else {
			banner += "*" + s + "*\n"
		}
	}
	return banner
}
