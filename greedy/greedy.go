package greedy

import "math"

// Grafo representado como lista de adyacencia
type Grafo map[int]map[int]int

// Encuentra la mejor solucion del optimo local actual para encontrar el siguiente mejor optimo local
/*
func Greedy(grafo Grafo, inicio int) int {
	visitados := make(map[int]bool)
	totalPeso := 0
	nodoActual := inicio

	for len(visitados) < len(grafo) {
		visitados[nodoActual] = true
		maxPeso := 0
		siguienteNodo := -1

		for vecino, peso := range grafo[nodoActual] {
			if peso > maxPeso && !visitados[vecino] {
				maxPeso = peso
				siguienteNodo = vecino
			}
		}
		if siguienteNodo == -1 {
			break
		}

		totalPeso += maxPeso
		nodoActual = siguienteNodo
	}

	return totalPeso
}
*/

func Greedy(grafo Grafo, inicio int) int {
	visitados := make(map[int]bool)
	totalPeso := 0
	nodoActual := inicio

	if len(grafo) == 0 {
		return totalPeso // Retorna 0 si el grafo está vacío
	}

	// Mantener un contador de nodos visitados
	nodosVisitados := 0

	for nodosVisitados < len(grafo) {
		visitados[nodoActual] = true
		maxPeso := math.MinInt // Inicializar como el valor más bajo posible
		siguienteNodo := -1

		// Buscar el vecino con el peso máximo no visitado
		for vecino, peso := range grafo[nodoActual] {
			if peso > maxPeso && !visitados[vecino] {
				maxPeso = peso
				siguienteNodo = vecino
			}
		}

		// Si no hay un siguiente nodo válido, terminamos
		if siguienteNodo == -1 {
			break
		}

		totalPeso += maxPeso
		nodoActual = siguienteNodo
		nodosVisitados++ // Aumentar el contador de nodos visitados
	}

	return totalPeso
}
