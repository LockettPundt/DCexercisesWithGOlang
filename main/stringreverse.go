package main

func stringreverse(s string) string {
	reversedString := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversedString += string(s[i])
	}
	return reversedString
}
