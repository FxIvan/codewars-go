package alternativeCapitalization

import (
	"strings"
)

func Capitalize(text string) []string {

	elementArray := []string{}
	concatLetter := ""
	concatLetterOther := ""

	arrayText := strings.Split(text, "")
	for i := 0; i < len(arrayText); i++ {
		if i%2 == 0 {
			concatLetter += strings.ToUpper(arrayText[i])
			concatLetterOther += strings.ToLower(arrayText[i])
		} else {
			concatLetter += strings.ToLower(arrayText[i])
			concatLetterOther += strings.ToUpper(arrayText[i])
		}
	}

	elementArray = append(elementArray, concatLetter)
	elementArray = append(elementArray, concatLetterOther)

	return elementArray
}
