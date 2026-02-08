package rgb_to_hex

import (
	"strconv"
)

func GetColorHexadecimal(rgbaCode int) string {
	if rgbaCode < 0 {
		return "00"
	}

	if rgbaCode > 255 {
		return "FF"
	}

	valueDivision := rgbaCode / 16
	restoDivision := rgbaCode % 16

	hexMap := map[int]string{
		10: "A",
		11: "B",
		12: "C",
		13: "D",
		14: "E",
		15: "F",
	}

	valueOne, ok := hexMap[valueDivision]
	if !ok {
		valueOne = strconv.Itoa(valueDivision)
	}

	valueTwo, ok := hexMap[restoDivision]
	if !ok {
		valueTwo = strconv.Itoa(restoDivision)
	}

	return valueOne + valueTwo
}

/*
RGB To Hex Conversion
https://www.codewars.com/kata/513e08acc600c94f01000001/train/go
La función rgb está incompleta. Complétela para que, al pasar valores decimales RGB, se devuelva una representación hexadecimal.
Los valores decimales válidos para RGB son de 0 a 255. Cualquier valor fuera de este rango debe redondearse al valor válido más cercano.

Nota: Su respuesta siempre debe tener 6 caracteres; la abreviatura 3 no funcionará en este caso.

Ejemplos (entrada --> salida):


```
255, 255, 255 --> "FFFFFF"
255, 255, 300 --> "FFFFFF"
0, 0, 0       --> "000000"
148, 0, 211   --> "9400D3"
```

Decimal a Hexadecimal

Del 1 al 9 Son todo iguales
10 -> A
11 -> B
12 -> C
13 -> D
14 -> E
15 -> F

package kata

import (
	"strconv"
)

func getColorHexadecimal(rgbaCode int) string {
    if rgbaCode < 0 {
	    return "00"
    }

  if rgbaCode > 255 {
    return "FF"
  }

	valueDivision := rgbaCode / 16
	restoDivision := rgbaCode % 16

	hexMap := map[int]string{
		10: "A",
		11: "B",
		12: "C",
		13: "D",
		14: "E",
		15: "F",
	}

	valueOne, ok := hexMap[valueDivision]
	if !ok {
		valueOne = strconv.Itoa(valueDivision)
	}

	valueTwo, ok := hexMap[restoDivision]
	if !ok {
		valueTwo = strconv.Itoa(restoDivision)
	}

	return valueOne + valueTwo
}


func RGB(r, g, b int) string {
  // Your code here
  return getColorHexadecimal(r) + getColorHexadecimal(g) + getColorHexadecimal(b)
}
*/
