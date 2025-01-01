package main

import (
	"fmt"
	_ "fmt"

	_ "github.com/codewars/method"
	_ "github.com/codewars/nodeTree"
	_ "github.com/codewars/object"
	cacke "github.com/codewars/peteTheBaker"
	_ "github.com/codewars/searchJSON"
	_ "github.com/codewars/searchWord"
	_ "github.com/codewars/toWeirdCase"
)

func main() {
	//output := toWeirdCase.ToWeirdCaseFix("This is a test Looks like you passed")
	//lastElementArray := method.LastElement([]string{"hola","chau","como_estas"});
	//deleteLastElementArray := method.PopElement([]string{"hola","chau","como_estas"})
	//unshiftElement := method.Unshift([]string{"hola","chau","como_estas"}, "holaaaaa")
	//encodeURI := method.EncodeURI("http://www.codewars.com/kata/test/ejemplo/te#t/$ola/ch#u/como+estas/");

	// a1 := method.StringSlice{"a", "b", "c"}
	// a2 := method.StringSlice{"d", "e", "f"}
	// result := a1.ConcatArray(a2)

	// handler := ProxyValue{
	// 	delete: func(prop []string, valueDeleted string) ([]string,error){
	// 		var newArray []string
	// 		for index,value := range prop {
	// 			if prop[index] != valueDeleted {
	// 				newArray = append(newArray, value)
	//             }
	// 		}
	// 		return newArray, nil
	// 	},
	// 	set: func(prop []string, valueAdd string) ([]string,error){
	// 		return append(prop, valueAdd),nil
	// 	},
	// }

	//Objetos de funciones
	// proxy := method.NewProxyValue()

	// resultHandler, err := proxy.Delete([]string{"hola", "chau", "como_estas"}, "chau")
	// if err != nil {
	// 	fmt.Println("Error en Delete:", err)
	// }

	// resultSet, err := proxy.Set(resultHandler, "holaaaaa")
	// if err != nil {
	// 	fmt.Println("Error en Set:", err)
	// }
	//Fin de Objetos de funciones

	//value := method.At([]string{"hola","chau","como_estas","fin"},-1)
	// Mostrar resultados
	// value := method.CopyWithin([]string{"hola","chau","como_estas","fin"},0,3)
	// fmt.Println("value:", value)

	// line := searchWord.SearchWord()
	// fmt.Println("line:",line)

	//Buscar alguna valor de una propiedad en un array de objetos
	// folderPath := "./tablasTemporadas" // Ruta a la carpeta donde están los archivos JSON
	// keyword := "Olympics Men"
	// searchJSON.SearchObjectTeam(keyword, folderPath)

	// Buscada de arbol inOrder
	// nodeTree.RecorrerEjemplo()

	// capitalized := alternativeCapitalization.Capitalize("Hola")
	// fmt.Println(capitalized)

	// ******** NO PUDE RESOLVERLO ********************************* //
	// ******* consecutiveFibNumbers.ConsecutiveFibonacciFR() ****** //
	// ************************************************************* //

	// https://www.codewars.com/kata/56f78a42f749ba513b00037f

	/*objRecipe := map[string]int{
		"flour": 30, "sugar": 200, "eggs": 5,
	}*/

	/*objRecipe := map[string]int{
		"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100,
	}*/

	/*available := map[string]int{
		"flour": 100, "sugar": 400, "eggs": 15, "milk": 200,
	}*/

	objRecipe := map[string]int{
		"apples": 3, "flour": 300, "milk": 100, "oil": 100, "sugar": 150,
	}

	available := map[string]int{
		"apples": 15, "flour": 2000, "milk": 2000, "oil": 20, "sugar": 500,
	}

	/*available := map[string]int{
		"sugar": 500, "flour": 2000, "milk": 2000,
	}*/
	output := cacke.Cakes(objRecipe, available)
	fmt.Print("Output: ", output, "\n")
	// FIN del reto: // https://www.codewars.com/kata/56f78a42f749ba513b00037f
}
