package numberFormatter

import (
	"fmt"
	"strconv"
	"strings"
)

// Kata https://www.codewars.com/kata/51ba717bb08c1cd60f00002f/train/go

/*
Un formato para expresar una lista ordenada de números enteros es utilizar una lista separada por comas que contenga:

*Números enteros individuales
*O un rango de números enteros, representado por el número inicial y el número final del rango, separados por un guion '-'.
*El rango incluye todos los números enteros dentro del intervalo, incluidos ambos extremos. No se considera un rango a menos que abarque al menos 3 números.

Por ejemplo: "12,13,15-17"

Completa la solución para que tome una lista de números enteros en orden creciente y devuelva una cadena correctamente formateada en el formato de rango.

Ejemplo:
-----------------------------------------------------------------------------------------------------
solution([-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-10--8,-6,-3-1,3-5,7-11,14,15,17-20"
-----------------------------------------------------------------------------------------------------
*/

func merge(a []int, b []int) []int {

	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {

		if a[i] < b[j] {

			final = append(final, a[i])

			i++

		} else {

			final = append(final, b[j])

			j++

		}

	}

	for ; i < len(a); i++ {

		final = append(final, a[i])

	}

	for ; j < len(b); j++ {

		final = append(final, b[j])

	}

	return final
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func searIntoArrayString(value string, element []string) bool {
	for i := 0; i < len(element); i++ {
		if element[i] == value {
			return true
		} else {
			return false
		}
	}

	return false
}

func convertStringArray(list []int) []string {

	var newArrayString []string

	for i := 0; i < len(list); i++ {
		newArrayString = append(newArrayString, strconv.Itoa(list[i]))
	}

	return newArrayString
}

func searchCeroIndex(list []int, value int) (int, bool) {
	for index, _ := range list {
		if list[index] == value {
			return index, true
		}
	}
	return value, false
}

func searchMinMax(list []int) (int, int) {
	min := list[0]
	max := list[len(list)]

	for index, _ := range list {
		if index < min {
			return min, max
		}

		if index > max {
			return min, max
		}
	}

	return min, max
}

func searchWithNumberNegative(list []int) string {
	var number string

	for i := list[len(list)]; i == 1; i-- {
		fmt.Print("searchOnlyNumberPositive: ", i, "\n")
	}
	return number
}

func searchOnlyNumberPositive(list []int, max int) string {
	var number string

	for i := 0; i < max; i++ {
		fmt.Print("searchOnlyNumberPositive: ", i, "\n")
	}
	return number
}

func Solution(list []int) string {
	orderList := mergeSort(list)
	fmt.Print("Order List --->", orderList, "\n")

	var listString strings.Builder
	_, max := searchMinMax(orderList)
	NumberNegativeIndex, _ := searchCeroIndex(orderList, 0)
	// fmt.Print("Numero Negativos: ", orderList[:NumberNegativeIndex], "\n")
	valueOne := searchWithNumberNegative(orderList[:NumberNegativeIndex])
	valueTwo := searchOnlyNumberPositive(orderList[NumberNegativeIndex:], max)
	// fmt.Print("Numero Positivo: ", orderList[NumberNegativeIndex:], "\n")
	fmt.Print(valueOne, valueTwo)

	return listString.String()
}
