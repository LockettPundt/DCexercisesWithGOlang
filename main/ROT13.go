package main

import (
	"fmt"
	"strings"
)

func rot13(s string) string {
	var cypher string
	for i := 0; i < len(s); i++ {

		if int(s[i]) < 97 || int(s[i]) > 122 {
			if int(s[i]) < 65 || int(s[i]) > 90 {
				cypher += string(s[i])
			}
			if string(s[i]) != strings.ToLower(string(s[i])) {
				upperCaseShifted := int(s[i]) + 13
				if upperCaseShifted > 90 {
					upperCaseShifted -= 26
				}
				upperCaseToAdd := rune(upperCaseShifted)
				cypher += string(upperCaseToAdd)
			}
		} else {
			shiftedLetter := int(s[i]) + 13
			if shiftedLetter > 122 {
				shiftedLetter -= 26
			}
			letterToAdd := rune(shiftedLetter)
			cypher += string(letterToAdd)
		}

	}
	fmt.Println("here is the encrypted messaage: ", cypher)
	return cypher
}
