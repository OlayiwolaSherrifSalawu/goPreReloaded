package core

var checkStrings = "aeiouhAEIOUH"

func ParseAndTransform(text string) string {
	functions := []func(string) string{Modify, Punctuation, CheckQuote, AToAN}
	var result string
	for _, funcs := range functions {
		result = funcs(text)
		text = result
	}
	return result
}
