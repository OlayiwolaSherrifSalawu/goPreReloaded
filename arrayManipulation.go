package main

import (
	"strconv"
	"strings"
)

var caseToLook = map[string]func(string) string{
	"bin": Binary,
	"up":  ToUpper,
	"low": ToLower,
	"hex": HexDecimal,
	"cap": Capitalize,
}
var checkStrings = "aeiouhAEIOUH"

func LookAtWord(text string) string {
	sliceOfWords, err := WordSplitter(text)
	if err != nil {
		return ""
	}
	var result []string
	skipNext := false //this falg is used to check and no when to skip numbers in (lower,<number>)and to avoid just appending it to my result.
	isQuote := false
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
		} else if strings.Trim(words, ".,!?:;") == "" {
			lastResultIndex := len(result) - 1
			if lastResultIndex >= 0 {
				result[lastResultIndex] = result[lastResultIndex] + words
			}
		} else if !isQuote && words == "'" { //this means  it is just finding the first quote.
			isQuote = true
			if i+1 < len(sliceOfWords) {
				result = append(result, words+sliceOfWords[i+1])
				skipNext = true
			}
		} else if isQuote && words == "'" { //this means  it is just finding the last quote.
			if len(result)-1 >= 0 {
				result[len(result)-1] = result[len(result)-1] + words
				isQuote = false
			}
		} else if (words == "a" || words == "A") && (i+1 < len(sliceOfWords)) && strings.Contains(checkStrings, string(sliceOfWords[i+1][0])) {
			switch words {
			case "a":
				result = append(result, "an")
			case "A":
				result = append(result, "An")
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

 