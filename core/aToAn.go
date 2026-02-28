package core

import (
	"reloaded/helper"
	"strings"
)

func AToAN(text string) string {
	sliceOfWord, err := helper.WordSplitter(text)
	if err != nil {
		return ""
	}
	var result []string
	lengthOfSlice := len(sliceOfWord)
	for i, word := range sliceOfWord {
		ahead := i + 1
		if (word == "a" || word == "A") && ahead < lengthOfSlice && strings.Contains(checkStrings, string(sliceOfWord[ahead][0])) {
			switch word {
			case "a":
				result = append(result, "an")
			case "A":
				result = append(result, "An")

			}
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
