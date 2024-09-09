package main

import (
	"fmt"

	_"github.com/codewars/toWeirdCase"
	"github.com/codewars/method"
)

func main(){
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
	value := method.CopyWithin([]string{"hola","chau","como_estas","fin"},0,3)
	fmt.Println("value:", value)
}