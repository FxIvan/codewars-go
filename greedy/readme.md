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

Ejemplo donde se podria implementar un algoritmo greedy:
```
El algoritmo greedy es útil en problemas donde tomar decisiones óptimas localmente también conduce a una solución globalmente óptima. Esto ocurre generalmente en problemas que tienen una estructura de optimalidad de subproblemas o cumplen la propiedad greedy (es decir, las decisiones locales óptimas siempre llevan a la mejor solución global).

Ejemplos comunes de uso:
1. Problemas de optimización combinatoria:
Problema de la mochila fraccionaria (Fractional Knapsack): Se seleccionan objetos para maximizar el valor total dentro de un límite de peso, eligiendo primero los objetos con mayor valor por unidad de peso.
Problema de cobertura de intervalos: Seleccionar la mínima cantidad de intervalos que cubren una línea de números.
2. Teoría de grafos:
Árbol de expansión mínima (MST): Algoritmos como Kruskal o Prim usan greedy para construir el MST seleccionando las aristas de menor peso.
Algoritmo de Dijkstra: Encuentra el camino más corto desde un nodo a todos los demás, seleccionando el nodo con la menor distancia acumulada en cada paso.
Algoritmo de coloración de grafos: Asignar colores a los nodos de un grafo con el mínimo número de colores.
3. Problemas de programación de tareas:
Problema de la selección de actividades (Activity Selection Problem): Seleccionar el máximo número de actividades que no se solapen, ordenándolas por su tiempo de finalización.
Planificación de tareas con plazos: Maximizar las tareas completadas antes de su fecha límite, eligiendo aquellas con mayor beneficio.
4. Compresión y codificación:
Codificación de Huffman: Utiliza greedy para construir un árbol binario óptimo para la codificación de datos, basado en la frecuencia de los caracteres.
5. Optimización financiera:
Cambio de monedas (Coin Change): Encontrar el mínimo número de monedas para una cantidad dada (solo funciona si las denominaciones cumplen ciertas propiedades).
6. Sistemas de redes y telecomunicaciones:
Maximización del ancho de banda: Seleccionar rutas óptimas para maximizar la capacidad de transmisión.
Enrutamiento en redes: Encontrar caminos más cortos o eficientes para la transmisión de datos.
Ventajas:
Simplicidad y eficiencia: Generalmente tiene tiempos de ejecución más rápidos porque toma decisiones locales inmediatas sin explorar todo el espacio de soluciones.
Aplicable a problemas en tiempo real: Por ejemplo, encontrar rutas más cortas en un mapa GPS mientras conduces.
Limitaciones:
No siempre garantiza una solución global óptima.
Solo funciona bien si el problema cumple las propiedades greedy mencionadas.
En resumen, el enfoque greedy es poderoso y práctico en problemas bien estructurados, pero es importante analizar si las decisiones locales conducen a la mejor solución global antes de aplicarlo.
```