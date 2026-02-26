package main

import (
	"strconv"
	"strings"
)

func LookAtWord(text string) string {
	caseToLook := map[string]func(string) string{
		"bin": Binary,
		"up":  ToUpper,
		"low": ToLower,
		"hex": HexDecimal,
		"cap": Capitalize,
	}
	sliceOfWords, err := WordSplitter(text)
	if err != nil {
		return ""
	}
	var result []string
	skipNext := false
	for i, words := range sliceOfWords {
		var newWord = ""
		if skipNext {
			skipNext = false
			continue
		}
		n := 1
		if strings.HasPrefix(words, "(") {
			newWord = strings.Trim(words, "(),")
			if strings.HasSuffix(words, ",") && i+1 < len(sliceOfWords) {
				n, err = strconv.Atoi(strings.TrimSuffix(sliceOfWords[i+1], ")"))
				skipNext = true
				if err != nil {
					return ""
				}
			}
		} else {
			result = append(result, words)
		}

		val, ok := caseToLook[newWord]
		if ok {
			for step := 1; step <= n; step++ {
				targetIndex := len(result) - step
				if targetIndex >= 0 {
					result[targetIndex] = val(result[targetIndex])
				}
			}
		}

	}
	return strings.Join(result, " ")
}
