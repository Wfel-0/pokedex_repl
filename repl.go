package main

import "strings"

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	output := []string{}

	str := ""

	for i := range text {
		if text[i] == ' ' && str != "" {
			output = append(output, str)
			str = ""
		} else if text[i] == ' ' {
			continue
		} else {
			str += string(text[i])
		}
	}

	if str != "" {
		output = append(output, str)
	}

	return output
}
