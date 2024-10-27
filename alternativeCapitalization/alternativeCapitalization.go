package alternativeCapitalization

import (
	"strings"
)

// func Capitalize(text string) []string {

// 	elementArray := []string{}
// 	concatLetter := ""
// 	concatLetterOther := ""

// 	arrayText := strings.Split(text, "")
// 	for i := 0; i < len(arrayText); i++ {
// 		if i%2 == 0 {
// 			concatLetter += strings.ToUpper(arrayText[i])
// 			concatLetterOther += strings.ToLower(arrayText[i])
// 		} else {
// 			concatLetter += strings.ToLower(arrayText[i])
// 			concatLetterOther += strings.ToUpper(arrayText[i])
// 		}
// 	}

// 	elementArray = append(elementArray, concatLetter)
// 	elementArray = append(elementArray, concatLetterOther)

// 	return elementArray
// }

// Usar buffers

/*
Usar strings.Builder es una alternativa para reducir la cantidad de
operaciones de concatenación, especialmente cuando trabajas con muchas
letras, ya que este es más eficiente en la gestión de memoria.
*/
func Capitalize(text string) []string {
	var builder1, builder2 strings.Builder

	for i, char := range text {
		if i%2 == 0 {
			builder1.WriteString(strings.ToUpper(string(char)))
			builder2.WriteString(strings.ToLower(string(char)))
		} else {
			builder1.WriteString(strings.ToLower(string(char)))
			builder2.WriteString(strings.ToUpper(string(char)))
		}
	}

	return []string{builder1.String(), builder2.String()}
}
