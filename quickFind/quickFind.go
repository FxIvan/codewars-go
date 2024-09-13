package quickFind

/*********
cnt   = conectar
vrfc  = verificar
uni   = unir

Inicio Programa
	Inicio de formulario
		variable ctdElemento  // número total de elementos
		variable cntUno       // primer elemento para conectar
		variable cntDos       // segundo elemento para conectar
		variable vrfcUno      // primer elemento para verificar conexión
		variable vrfcDos      // segundo elemento para verificar conexión
		variable uniUno       // primer elemento para unir
		variable uniDos       // segundo elemento para unir
	Fin formulario

	Inicio procedimiento QuickFind
		// Función para inicializar el QuickFindUF con una cantidad de elementos
		funcion NewQuickFindUF(ctdElemento numero)
			crear un array de tamaño ctdElemento
			asignar a cada índice su valor inicial desde 0 hasta ctdElemento - 1
			retornar el array con los valores inicializados
		fin funcion

		// Función para verificar si dos elementos están conectados
		funcion VerificarConectados(vrfcUno, vrfcDos)
			retornar si el valor en el índice vrfcUno es igual al valor en el índice vrfcDos
		fin funcion

		// Función para unir dos elementos
		funcion UnirElementos(uniUno, uniDos)
			guardar el valor del índice uniUno en variable pid
			guardar el valor del índice uniDos en variable qid

			// Reemplazar todos los elementos con el valor pid por el valor qid
			para cada índice en el array
				si el valor en el índice es igual a pid
					asignar a ese índice el valor de qid
				fin si
			fin para
		fin funcion
	Fin procedimiento
Fin programa
**********/

import(
	"fmt"
)

type QuickFindUF  struct{
	id	[]int
}

func NewQuickFindUF(N int) *QuickFindUF{
	uf := &QuickFindUF{
		id: make([]int,N),
	}
	for i := 0 ; i < N ; i++{
		uf.id[i] = i
	}

	return uf
}

func (qf *QuickFindUF) Unir(p,q int){
	pid := qf.id[p]
	qid := qf.id[q]

	for i := 0; i < len(qf.id); i++{
		if qf.id[i] == pid{
			qf.id[i] = qid
		}
	}

	fmt.Println("value:",qf,"qf")
}

func (qf *QuickFindUF) VerifyJoint(p, q int) bool{
	fmt.Println("Es verdadero la union:",qf.id[p] == qf.id[q])
	return qf.id[p] == qf.id[q]
}
