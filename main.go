package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sentence := "It has been 10 (bin) years"

	ans := binary(sentence)
	fmt.Println(ans)
}

func binary(sentence string) string {
	newSlice := strings.Split(sentence, " ")

	for i := 0; i < len(newSlice); i++ {
		if i > 0 && newSlice[i] == "(bin)" {

			num, err := strconv.ParseInt(newSlice[i-1], 2, 64)
			if err != nil {
				fmt.Println(err)
				break
			}
			newSlice[i-1] = fmt.Sprintf("%d", num)
			newSlice = append(newSlice[:i], newSlice[i+1:]...)
		}

	}
	return strings.Join(newSlice, " ")
}
