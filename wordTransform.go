package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Goal create a function that takes strings and return  a slice of strings

func WordSplitter(text string) ([]string, error) {
	// handle edge case (1. if the input parameter is empty)
	if text == "" {
		return nil, fmt.Errorf("cannot split empty string")
	}
	// spilting string.
	var store []string
	store = strings.Fields(text)
	return store, fmt.Errorf("cannot split empty string")
}

// Case Manipulators (ToUpper,ToLower,Capitalize)
func ToLower(text string) string {
	// handle edge case (1. if the input parameter is empty)
	if text == "" {
		return "cannot split convert empty string"
	}
	return strings.ToLower(text)
}
func ToUpper(text string) string {
	// handle edge case (1. if the input parameter is empty)
	if text == "" {
		return "cannot split convert empty string"
	}
	return strings.ToUpper(text)
}

func Capitalize(text string) string {
	// handle edge case (1. if the input parameter is empty)
	if text == "" {
		return "cannot split convert empty string"
	}
	return strings.ToTitle(strings.ToLower(text))
}

// Base Coverter in Go(binary, and Hexdecimal)

func Binary(text string) string {
	// handle edge case (1. if the input parameter is empty)
	if text == "" {
		return "cannot convert empty string"
	}
	textLen := len(text)

	store := ""
	for i := 0; i < textLen; i++ {
		set, _ := strconv.Atoi(string(int(text[i]) * int(math.Pow(2.0, float64(textLen-i)))))
		store += strconv.Itoa((set))
	}
	return store
}
