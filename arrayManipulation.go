	package main

	import (
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
			previousChar := len(result) - 1
			val, ok := caseToLook[words]
			if !ok {
				result = append(result, words)
			} else if len(result) > 0 && ok {
				result[previousChar] = val(result[previousChar])
			}
		}
		return strings.Join(result, " ")
	}
