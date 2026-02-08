package main

import (
	"fmt"

	_ "github.com/codewars/faberger"
	_ "github.com/codewars/method"
	_ "github.com/codewars/nodeTree"
	_ "github.com/codewars/object"
	_ "github.com/codewars/peteTheBaker"
	_ "github.com/codewars/recursive"
	"github.com/codewars/rgb_to_hex"

	// recursive "github.com/codewars/recursive"
	_ "github.com/codewars/searchJSON"
	_ "github.com/codewars/searchWord"
	_ "github.com/codewars/toWeirdCase"
)

/*
	func timer(name string) func() {
		start := time.Now()
		return func() {
			fmt.Printf("\n<<<<<<<<<<<<<<<<<<<<<<< TIMER %s : %v >>>>>>>>>>>>>>>>>>>>>>>\n", name, time.Since(start))
		}
	}
*/
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

	// https://www.codewars.com/kata/56f78a42f749ba513b00037f INICIO
	/*objRecipe := map[string]int{
		"flour": 30, "sugar": 200, "eggs": 5,
	}*/
	/*objRecipe := map[string]int{
		"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100,
	}*/
	/*available := map[string]int{
		"flour": 100, "sugar": 400, "eggs": 15, "milk": 200,
	}*/
	// objRecipe := map[string]int{
	//	"apples": 3, "flour": 300, "milk": 100, "oil": 100, "sugar": 150,
	// }
	// available := map[string]int{
	//	"apples": 15, "flour": 2000, "milk": 2000, "oil": 20, "sugar": 500,
	// }
	/*available := map[string]int{
		"sugar": 500, "flour": 2000, "milk": 2000,
	}*/
	// output := cacke.Cakes(objRecipe, available)
	// fmt.Print("Output: ", output, "\n")
	// FIN del reto: // https://www.codewars.com/kata/56f78a42f749ba513b00037f

	// fmt.Print(recursive.FindSum(5), "\n")
	// n_veces := recursive.SerieFibonacci(3)
	// fmt.Print("Recursivo: ", n_veces, "\n")

	// Greedy
	// Grafo de prueba
	// https://chatgpt.com/c/678d095e-1b7c-8001-9099-df3814ee33aa
	// grafo := greedy.Grafo{
	// 	0: {1: 4, 2: 1}, // 1 vecino : 4 Peso | Siguiente Nodo = 1
	// 	1: {3: 8, 4: 1}, // 3 vecino : 8 Peso | Siguiente Nodo = 3
	//	2: {1: 15, 8: 5},
	//	3: {6: 7},
	// }
	// fmt.Println("Resultado de Greedy:", greedy.Greedy(grafo, 0))

	/////////////////////////////////////////////////////////////////////////////
	// MontecarloPi
	// https://chat.deepseek.com/a/chat/s/4e31b531-bb1e-4e30-820c-2b8d17629d77
	// numPoints := 100000
	// estimatedPi := montecarloPi.MontecarloPi(numPoints)
	// fmt.Printf("Estimación de π usando %d puntos: %f\n", numPoints, estimatedPi)

	/////////////////////////////////////////////////////////////////////////////
	//NO RESOLVIDO
	// https://www.codewars.com/kata/59f98052120be4abfa000304
	// response := faberger.UpsideDown("100000", "12345678900000000")
	// fmt.Print("\n<<<<<<<<<<<<<<<<<<<<<<< START FABERGER >>>>>>>>>>>>>>>>>>>>>>>\n")
	// response := faberger.UpsideDown("0", "12345678900000000")
	// fmt.Print("\n FINISH: ", response, "\n")
	// fmt.Print("\n<<<<<<<<<<<<<<<<<<<<<<< END FABERGER >>>>>>>>>>>>>>>>>>>>>>>\n")

	//////////////////////////////////////////////////////////////////////////////
	// array := []int{-9, -10, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 20, 19}
	// list := numberFormatter.Solution(array)
	// fmt.Print("\n String List", list, "\n")
	// fmt.Println(countIPAddresses.IpsBetween("10.0.0.0", "10.0.0.50"))
	// fmt.Println(countIPAddresses.IpsBetween("20.0.0.10", "20.0.1.0"))
	// fmt.Println(countIPAddresses.IpsBetween("10.11.12.13", "10.11.13.0"))

	//////////////////////////////////////////////////////////////////////////////
	// RGB To Hex Conversion
	// https://www.codewars.com/kata/513e08acc600c94f01000001/train/go
	// Color blanco en el ejemplo que seria #FFFFFF
	fmt.Print(rgb_to_hex.GetColorHexadecimal(255) + rgb_to_hex.GetColorHexadecimal(255) + rgb_to_hex.GetColorHexadecimal(255))
}
