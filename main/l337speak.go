package main

import str "strings"

func l337speak(s string) string {
	length := len(s)
	var leetResult string
	for i := 0; i < length; i++ {
		switch str.ToLower(string(s[i])) {
		case "a":
			leetResult += "4"
		case "e":
			leetResult += "3"
		case "g":
			leetResult += "6"
		case "i":
			leetResult += "1"
		case "o":
			leetResult += "0"
		case "s":
			leetResult += "5"
		case "t":
			leetResult += "7"
		default:
			leetResult += str.ToLower(string(s[i]))
		}
	}
	return leetResult
}
