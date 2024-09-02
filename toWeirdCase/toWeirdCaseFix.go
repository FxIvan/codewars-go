package toWeirdCase

/**
*https://www.codewars.com/kata/52b757663a95b11b3d00062d
*
*
*
*/

import (
	"strings"
	"fmt"
)


func ToWeirdCaseFix(str string) []string{
	arrayString := strings.Split(str, " ")
	var arrayWords []string
	for _ , word := range arrayString{
		var newArray []string
		for i,r := range word {
			if i%2 == 0{
				newArray = append(newArray, strings.ToUpper(string(r)))
			}else{
				newArray = append(newArray, strings.ToLower(string(r)))
			}
		}
		fmt.Println("newArray:",newArray)
		arrayWords = append(arrayWords , strings.Join(newArray, ""))
		fmt.Println("arrayWords:",arrayWords)
	}
	return arrayWords
}