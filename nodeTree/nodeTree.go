package nodeTree

import "fmt"

type Node struct{
	Valor int
	Izquierda *Node
	Derecha *Node
}


func RouteInOrder(nodo *Node, recorrido *[]int){
	if nodo == nil{
		return
	}

	RouteInOrder(nodo.Izquierda, recorrido);
	*recorrido = append(*recorrido, nodo.Valor)
	RouteInOrder(nodo.Derecha, recorrido);
}

func RecorrerEjemplo(){
	    // Crear los nodos
		raiz := &Node{Valor: 10}
		raiz.Izquierda = &Node{Valor: 5}
		raiz.Derecha = &Node{Valor: 20}
		raiz.Izquierda.Izquierda = &Node{Valor: 3}
		raiz.Izquierda.Derecha = &Node{Valor: 7}
	
		// Crear un array para almacenar el recorrido
		var recorrido []int
	
		// Llamar a la función RouteInOrder para recorrer el árbol
		RouteInOrder(raiz, &recorrido)
	
		// Mostrar el recorrido en orden
		fmt.Println("Recorrido InOrder:", recorrido)
}