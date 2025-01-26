package recursive

func SerieFibonacci(n int) int {
	new_values := []int{}

	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	result := SerieFibonacci(n-1) + SerieFibonacci(n-2)
	new_values = append(new_values, result)

	return new_values[0]
}

/*
	https://chatgpt.com/c/67756bbc-1940-8001-97bc-1ca9ad1b0888
	Ejercicio: Generar la Serie de Fibonacci
	Escribe una función recursiva en Go que tome un número entero n como entrada y devuelva el término n de la serie de Fibonacci.

	Reglas de la Serie de Fibonacci:
	El primer término (n = 0) es 0.
	El segundo término (n = 1) es 1.
	A partir de ahí, cada término es la suma de los dos anteriores.
	Ejemplo de entrada/salida:
	Entrada: 6
	Salida: 8 (La serie sería: 0, 1, 1, 2, 3, 5, 8)
	Requisitos adicionales:
	Asegúrate de manejar correctamente valores menores o iguales a 0 devolviendo 0.
	Escribe un programa que llame a esta función y muestre el resultado en la consola.
*/
