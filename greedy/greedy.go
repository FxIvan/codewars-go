package greedy

// Grafo representado como lista de adyacencia
type Grafo map[int]map[int]int

// Encuentra la mejor solucion del optimo local actual para encontrar el siguiente mejor optimo local

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
