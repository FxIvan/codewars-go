Bienvenido a la Edición de Desafío de Números al Revés
En el kata original de @KenKamau, estabas limitado a números enteros menores que 2^17. Aquí, se te darán cadenas de dígitos de hasta 42 caracteres de longitud (el límite superior es 10^42 - 1).

Tu tarea es esencialmente la misma, pero el desafío adicional consiste en crear una solución rápida y eficiente.

Entrada:
Tu función recibirá dos cadenas, cada una compuesta por dígitos que representan un número entero positivo. Estos dos valores representarán los límites inferior y superior de un rango.

Salida:
Tu función debe devolver la cantidad de números válidos que son números al revés dentro del rango de los dos argumentos de entrada, incluyendo ambos límites.

¿Qué es un Número al Revés?
Un número al revés es un número entero que se ve igual cuando se rota 180 grados, como se ilustra a continuación.

1961 → Válido
88 → Válido
101 → Válido
25 → No válido

Ejemplo:
```
var x, y string = "0", "25";
upsideDown(x, y); // 4
// Los números válidos en el rango son 0, 1, 8 y 11.
```

Notas adicionales:
Todas las entradas serán válidas.
El primer argumento siempre será menor que el segundo (es decir, el rango siempre será válido).
Si disfrutaste este kata, asegúrate de revisar mis otros katas.

Borrador
```
package faberger

import (
	"fmt"
	"strconv"
	"strings"
)

func contains(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ValidateInvertedNumber(numberText int) bool {
	validNumbers := []string{"0", "1", "6", "8", "9"}
	numberTextStr := strconv.Itoa(numberText)
	lenNumberTextStr := len(numberTextStr)
	numberSeparated := strings.Split(numberTextStr, "")
	// mountIteration := 0
	validationStatus := false

	for i := 0; i < lenNumberTextStr; i++ {
		validationStatus = contains(validNumbers, numberSeparated[i])
		// Si el primero ya es false entonces devolver false o El ultimo
		if !validationStatus {
			return validationStatus
		}
	}

	return validationStatus
}

func UpsideDown(n1, n2 string) uint64 {
	n1Numer, err := strconv.Atoi(n1)
	n2Number, err := strconv.Atoi(n2)
	outputValidateNumbers := []uint64{}

	if err != nil || n1Numer > n2Number {
		if err != nil {
			panic(err)
		}
		panic("El primer numero el Mayor al Segundo y eso no puede hacerse")
	}

	for i := n1Numer; i < n2Number; i++ {
		isValid := ValidateInvertedNumber(i)
		/*fmt.Print("i -->", i, "\n")
		fmt.Print("isValid -->", isValid, "\n")
		fmt.Print("\n---------------------------------\n")*/
		if isValid {
			outputValidateNumbers = append(outputValidateNumbers, uint64(i))
		}
	}
	fmt.Print("outputValidateNumbers -->", outputValidateNumbers, "\n")
	return uint64(len(outputValidateNumbers))
}
```

```
// Borrador sin optimizacion
package faberger

import (
	"fmt"
	"strconv"
)

func searchValuePalindrome(num int, palindrome [10]int) (string, bool, string) {
	for _, v := range palindrome {
		if v == num {
			numStr := strconv.Itoa(palindrome[num])
			unturned := strconv.Itoa(num)
			return numStr, true, unturned
		}
	}
	return "", false, ""
}

func revertNumber(txt string) (result string) {
	for _, v := range txt {
		result = string(v) + result
	}
	return
}

func NumberOfRepetition(arrayOne []string, arrayTwo []string) uint64 {
	var amountRepeat int

	for iOne, vOne := range arrayOne {
		for iTwo, vTwo := range arrayTwo {
			if vOne == vTwo && iOne == iTwo && vOne != "" && vTwo != "" {
				amountRepeat += 1
			}
		}
	}

	return uint64(amountRepeat)
}

func UpsideDown(n1, n2 string) uint64 {
	min, _ := strconv.Atoi(n1)
	max, _ := strconv.Atoi(n2)

	palindrome := [10]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}

	invertedNumbers := make([]string, len(n2))
	numberPalindrome := make([]string, len(n2))

	reverseNumbers := make([]string, len(n2))

	for i := min; i <= max; i++ {

		// intoInvertNumb
		var valueNumber string
		var unturnedTxt string

		// Convertir a string el numero
		strNumber := strconv.Itoa(i)

		valid := true
		for _, a := range strNumber {
			intA, _ := strconv.Atoi(string(a))
			str, exist, unturned := searchValuePalindrome(intA, palindrome)
			if !exist {
				valid = false
				break
			}

			unturnedTxt += unturned
			valueNumber += str

		}

		if !valid {
			continue
		}

		invertedNumbers = append(invertedNumbers, valueNumber)
		numberPalindrome = append(numberPalindrome, unturnedTxt)
	}

	for _, v := range invertedNumbers {
		if v == "" { // Si v está vacío, lo saltamos
			continue
		}
		reverseNumbers = append(reverseNumbers, revertNumber(v))
	}
	fmt.Print(reverseNumbers)
	fmt.Print(numberPalindrome)
	return NumberOfRepetition(reverseNumbers, numberPalindrome)
}

```

Tiempo de Ejecucion Sin ejecucion en paralelo
8.79
7.96
7.97
8.48
8.14
8.28
Promedio: 8.27 segundos

Tiempo de Ejecucion En paralelo
7.57
7.91
7.79
7.55
7.56
7.78
Promedio:  7.69 segundos
