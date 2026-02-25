package main

import (
	"fmt"
	"log"
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
	store, err := strconv.ParseInt(text, 2, 64)
	if err != nil {
		log.Fatal("couldnt convert error")
	}
	return strconv.Itoa(int(store))
}

func HexDecimal(text string) string {
	// handle edge case (1. if the input parameter is empty)
	if text == "" {
		return "cannot convert empty string"
	}
	store, err := strconv.ParseInt(text, 16, 64)
	if err != nil {
		log.Fatal("couldnt convert error")
	}
	return strconv.Itoa(int(store))
}
