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

	a1 := method.StringSlice{"a", "b", "c"}
	a2 := method.StringSlice{"d", "e", "f"}

	result := a1.ConcatArray(a2)
	fmt.Println(result);
}