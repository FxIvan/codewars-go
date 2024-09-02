package toWeirdCase

import (
	"strings"
	"fmt"
)

func ToWeirdCaseGPT(str string) string {
	words := strings.Split(str, " ")
	var newArray []string
	fmt.Println("Array:",words)
	for _, word := range words {
		var newWord []string
		for i, r := range word {
			if i%2 == 0 {
				newWord = append(newWord, strings.ToUpper(string(r)))
			} else {
				newWord = append(newWord, strings.ToLower(string(r)))
			}
		}
		newArray = append(newArray, strings.Join(newWord, ""))
	}

	return strings.Join(newArray, " ")
}
