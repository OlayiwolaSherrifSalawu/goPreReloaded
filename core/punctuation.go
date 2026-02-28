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
		lengthOfLast := len(result) - 1
		if strings.Trim(word, ".,!?:;") == "" {
			lengthOfLast := len(result) - 1
			if lengthOfLast >= 0 {
				result[lengthOfLast] = result[lengthOfLast] + word
			}
		} else if checkFront(word) {
			result[lengthOfLast] = result[lengthOfLast] + string(word[0])
			result = append(result, word[1:])
		} else if checkBack(word) {
			result[lengthOfLast] = result[lengthOfLast] + string(word[len(word)-1])
			result = append(result, word[:len(word)-1])
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func checkFront(text string) bool {
	word := false
	if strings.HasPrefix(text, ",") || strings.HasPrefix(text, "?") || strings.HasPrefix(text, ".") || strings.HasPrefix(text, "?") || strings.HasPrefix(text, ":") || strings.HasPrefix(text, ";") {
		word = true
	}
	return word
}
func checkBack(text string) bool {
	word := false
	if strings.HasSuffix(text, ",") || strings.HasSuffix(text, "?") || strings.HasSuffix(text, ".") || strings.HasSuffix(text, "?") || strings.HasSuffix(text, ":") || strings.HasSuffix(text, ";") {
		word = true
	}
	return word
}
