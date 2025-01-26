Un algoritmo greedy es aquel que tomando exclusivamente la solución óptima local puede generar una solución óptima global.

Se utiliza en problemas donde se busca encontrar el mínimo o el máximo de algo.

### Ejemplo

Un ejemplo un poco más de la vida real podría ser el de organizar tareas en rangos de tiempo: imaginemos que
tenemos muchísimas cosas que hacer en nuestro día a día, pero queremos maximizar la cantidad de tareas que realizamos
diariamente. Nosotros no podemos dividir nuestra atención en dos tareas a la vez, así que una acción lógica es organizarnos
de manera que ejecutemos las tareas más pequeñas primero y, luego, las que nos consuman más tiempo.

Algo interesante es que podemos incluso añadir importancia en forma de “peso”, es decir, quizás una tarea muy larga vale más
que todas las pequeñas juntas. Esto también puede ser solucionado agregando un par de condiciones al ordenamiento, haciendo
una especie de proporción o “ratio” entre el tiempo que toma realizar una tarea y la importancia que tiene en escala numérica.


### Pregunta
```
Una consulta, se saltan el nodo 2, esta bien? Asi seria el grafo que le paso: 	grafo := greedy.Grafo{
		0: {1: 4, 2: 1}, // 1 vecino : 4 Peso | Siguiente Nodo = 1
		1: {3: 8, 4: 1}, // 3 vecino : 8 Peso | Siguiente Nodo = 3
		2: {1: 15, 8: 5},
		3: {6: 7},
	}
```

Respuesta de GPT:
```
Con el grafo que pasaste, sí es posible que se salte el nodo 2, y eso está correcto en un algoritmo greedy, ya que este enfoque toma decisiones basándose únicamente en la opción más óptima en cada paso, sin tener en cuenta cómo podría afectar el resultado global.

Explicación paso a paso:
Inicio en el nodo 0:

Vecinos: {1: 4, 2: 1}.
Selecciona el vecino con mayor peso (nodo 1, peso 4).
Total acumulado: 4.
Nodo actual: 1:

Vecinos: {3: 8, 4: 1}.
Selecciona el vecino con mayor peso (nodo 3, peso 8).
Total acumulado: 4 + 8 = 12.
Nodo actual: 3:

Vecinos: {6: 7}.
Solo hay un vecino (nodo 6, peso 7).
Total acumulado: 12 + 7 = 19.
Nodo actual: 6:

No tiene vecinos, se detiene.
¿Por qué se salta el nodo 2?
El nodo 2 nunca es visitado porque no hay un camino hacia él con un peso mayor que los caminos que ya tomaste:

Desde el nodo 0, el camino hacia el nodo 1 tiene un peso mayor que el camino hacia el nodo 2.
Una vez que decides ir al nodo 1, el algoritmo no vuelve a considerar el nodo 2 porque los caminos se basan en vecinos directos desde el nodo actual.
Conclusión:
Esto es esperado en un algoritmo greedy. Si quieres asegurarte de visitar todos los nodos o encontrar el camino más corto a todos ellos, necesitas un algoritmo diferente, como Dijkstra o BFS/DFS modificados para ese propósito.
```
