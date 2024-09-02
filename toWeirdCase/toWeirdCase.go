package toWeirdCase

/**
*https://www.codewars.com/kata/52b757663a95b11b3d00062d
*
*
*
*/

import (
	"strings"
)

func stringToArray(text string) []string{
	arrayString := make([]string , len(text))
	for i, r := range text{
		arrayString[i] = string(r)
	}
	return arrayString
}

func ToWeirdCase(str string) string{

	arrayConvert := stringToArray(str)

	var newArray []string
	for i := 0; i < len(arrayConvert); i++ {
		if i % 2 == 0 {
			if(arrayConvert[i] != " "){
				newArray = append(newArray, strings.ToUpper(arrayConvert[i]))
			}else{
				newArray = append(newArray, " ")
			}
		}else{
			if(arrayConvert[i] != " "){
				newArray = append(newArray, strings.ToLower(arrayConvert[i]))
			}else{
				newArray = append(newArray, " ")
			}
		}
	}
	justString := strings.Join(newArray, "")

	return justString
}