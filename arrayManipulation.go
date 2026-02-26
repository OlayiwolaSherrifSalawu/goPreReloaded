package main

func LookAtWord(text string) string {
	caseToLook := map[string]func(string) string{
		"(bin)": Binary,
		"(up)":  ToUpper,
		"(low)": ToLower,
		"(hex)": HexDecimal,
		"(cap)": Capitalize,
	}
}
