package main

import (
	"fmt"

	_"github.com/codewars/toWeirdCase"
	_"github.com/codewars/method"
)

type ProxyValue struct{
	delete func(prop []string, valueDeleted string) ([]string,error);
	set func(prop []string, valueAdd string) ([]string,error);
}


func main(){
	//output := toWeirdCase.ToWeirdCaseFix("This is a test Looks like you passed")
	//lastElementArray := method.LastElement([]string{"hola","chau","como_estas"});
	//deleteLastElementArray := method.PopElement([]string{"hola","chau","como_estas"})
	//unshiftElement := method.Unshift([]string{"hola","chau","como_estas"}, "holaaaaa")
	//encodeURI := method.EncodeURI("http://www.codewars.com/kata/test/ejemplo/te#t/$ola/ch#u/como+estas/");

	// a1 := method.StringSlice{"a", "b", "c"}
	// a2 := method.StringSlice{"d", "e", "f"}
	// result := a1.ConcatArray(a2)

	handler := ProxyValue{
		delete: func(prop []string, valueDeleted string) ([]string,error){
			var newArray []string
			for index,value := range prop {
				if prop[index] != valueDeleted {
					newArray = append(newArray, value)
                }
			}
			return newArray, nil
		},
		set: func(prop []string, valueAdd string) ([]string,error){
			return append(prop, valueAdd),nil
		},
	}

	resultHandler, err := handler.delete([]string{"hola","chau","como_estas"}, "chau")
	resultSet, err := handler.set(resultHandler, "holaaaaa")
	if err != nil {
        fmt.Println(err)
    }

	fmt.Println(resultHandler)
	fmt.Println(resultSet)
}