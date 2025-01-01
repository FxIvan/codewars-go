package cacke

/*
https://www.codewars.com/kata/525c65e51bf619685c000059/train/go
Pete le gusta hornear pasteles. Tiene algunas recetas y los ingredientes necesarios. Desafortunadamente, no es bueno con las matemáticas.
¿Puedes ayudarlo a averiguar cuántos pasteles podría hornear considerando sus recetas?
Escribe una función llamada cakes(), que

tome la receta (un objeto) y los -> recipe
ingredientes disponibles (también un objeto)  available
y devuelva el número máximo de pasteles que Pete puede hornear (un número entero).

Para simplificar, no hay unidades para las cantidades (por ejemplo, 1 lb de harina o 200 g de azúcar se representan simplemente como 1 o 200).
Los ingredientes que no estén presentes en los objetos pueden considerarse como 0.

// must return 2
cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200});
// must return 0
cakes({apples: 3, flour: 300, sugar: 150, milk: 100, oil: 100}, {sugar: 500, flour: 2000, milk: 2000});
*/

/*
func Cakes(recipe map[string]int, available map[string]int) int {
	//Existen todos los ingredientes de recipe?
	value_temp := new(int)
	*value_temp = 1
	cantRecipe := new(int)
	*cantRecipe = 1
	var newArray []int
	for key, value := range recipe {
		valueFind := available[key] / value
		if valueFind == 0 {
			value_temp = &valueFind
			break
		}
		newArray = append(newArray, valueFind)
	}

	if value_temp == nil {
		return 0
	}
	if newArray == nil {
		return 0
	}
	sort.Ints(newArray)
	minimum := newArray[0]
	return minimum
}
*/

func Cakes(recipe map[string]int, available map[string]int) int {
	cantRecipe := -1
	for key, value := range recipe {
		valueFind := available[key] / value
		if valueFind == 0 {
			return 0
		}
		if cantRecipe == -1 || valueFind < cantRecipe {
			cantRecipe = valueFind
		}
	}
	return cantRecipe
}
