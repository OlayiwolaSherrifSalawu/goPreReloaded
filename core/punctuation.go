package core

import (
	reloaded "reloaded/helper"
	"strings"
)

func Punctuation(text string) string {
	sliceOfWords, err := reloaded.WordSplitter(text)
	if err != nil {
		return ""
	}
	var result []string
	for _, word := range sliceOfWords {
		if strings.Trim(word, ".,!?:;") == "" {
			lengthOfLast := len(result) - 1
			if lengthOfLast >= 0 {
				result[lengthOfLast] = result[lengthOfLast] + word
			} else {
				result = append(result)
			}
		}
	}
	return strings.Join(result, " ")
}
