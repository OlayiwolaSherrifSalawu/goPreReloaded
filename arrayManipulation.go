package main

import (
	"fmt"
	"strings"
)

func LookAtWord(text string) string {
	caseToLook := map[string]func(string) string{
		"(bin)": Binary,
		"(up)":  ToUpper,
		"(low)": ToLower,
		"(hex)": HexDecimal,
		"(cap)": Capitalize,
	}
	sliceOfWords, err := WordSplitter(text)
	if err != nil {
		return ""
	}
	var result []string
	for _, words := range sliceOfWords {
		lenOfResult := len(result) - 1
		if words != "(up)" || words != "(low)" || words != "(row)" || words != "(cap)" || words != "(bin)" || words != "(hex)" {
			result = append(result, words)
		} else if lenOfResult > 0 && words == "(up)" || words == "(low)" || words == "(row)" || words == "(cap)" || words == "(bin)" || words == "(hex)" {
			val, ok := caseToLook[words]
			if ok {
				result[lenOfResult] = fmt.Sprintf("%s", val(words))
			}
		}
	}
	return strings.Join(result," ")
}
