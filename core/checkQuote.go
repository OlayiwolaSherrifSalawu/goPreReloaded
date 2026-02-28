package core

import (
	"reloaded/helper"
	"strings"
)

func CheckQuote(text string) string {
	isQuote := false
	isNext := false
	var result []string
	sliceOfWords, err := helper.WordSplitter(text)
	if err != nil {
		return ""
	}
	for i, word := range sliceOfWords {
		if isNext {
			isNext = false
			continue
		}
		lengthOfLast := len(result) - 1
		if word == "'" && !isQuote && i+1 < len(sliceOfWords) {
			isQuote = true
			result = append(result, word+sliceOfWords[i+1])
			isNext = true
		} else if word == "'" && isQuote && lengthOfLast > 1 {
			result[lengthOfLast] = result[lengthOfLast] + word
			isQuote = false
		} else {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}
