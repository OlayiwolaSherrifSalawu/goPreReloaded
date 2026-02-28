package core

import (
	"reloaded/helper"
	"strconv"
	"strings"
)

var CaseToLook = map[string]func(string) string{
	"bin": helper.Binary,
	"up":  helper.ToUpper,
	"low": helper.ToLower,
	"hex": helper.HexDecimal,
	"cap": helper.Capitalize,
}

func Modify(text string) string {
	sliceOfWords, err := helper.WordSplitter(text)
	if err != nil {
		return ""
	}
	skipNext := false //this flag is needed for skiping numbers in (lower,<number>)and to avoid just appending it to my result.
	var result []string
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
		val, ok := CaseToLook[newWord]
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
