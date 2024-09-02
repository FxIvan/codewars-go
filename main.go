package main

import (
	"fmt"

	_"github.com/codewars/toWeirdCase"
	"github.com/codewars/method"
)

func main(){
	//output := toWeirdCase.ToWeirdCaseFix("This is a test Looks like you passed")
	lastElementArray := method.LastElement([]string{"hola","chau","como_estas"});
	fmt.Println(lastElementArray);
}