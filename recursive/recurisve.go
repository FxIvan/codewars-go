package recursive

func FindSum(n int) int {
	if n == 0 {
		return 0
	}

	return n + FindSum(n-1)
}

/*
	Una función recursiva debe tener al menos un caso base
	y un caso recursivo, de lo contrario se ejecutará indefinidamente
	o provocará un error de desbordamiento de pila.

	La recursividad puede hacer que algunos algoritmos sean más fáciles de entender
	e implementar, especialmente cuando el problema tiene una estructura recursiva
	o se puede dividir en subproblemas similares.

	Por ejemplo, la recursividad se utiliza a menudo para atravesar árboles, grafos
	y otras estructuras de datos que tienen una naturaleza jerárquica o ramificada.

	¿Cómo usar la recursividad de manera efectiva?
	Sebe seguir algunas prácticas recomendadas para asegurarse de que funciona de manera
	correcta y eficiente.

	1. Defina un caso base y un caso recursivo, asegurándose de que el caso base sea accesible
	y que el caso recursivo reduzca el tamaño del problema.

	2. Compruebe el tamaño de la pila y el uso de memoria de la función recursiva, evitando
	llamadas recursivas innecesarias o excesivas que puedan causar un error de
	desbordamiento de pila o una pérdida de memoria

	3. Busque formas de optimizar la función recursiva mediante la recursividad de cola,
	la memoización, la poda u otras técnicas que puedan mejorar el rendimiento o reducir el consumo de memoria.
*/
