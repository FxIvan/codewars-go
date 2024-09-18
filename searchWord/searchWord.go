package searchWord

import (
	"fmt"
	_"io/ioutil"
	"log"
	"os"
	"strings"
	"bufio"
	_"strconv"
)
/******
1. Buscar Palabras y imprimir de esa palabra 20 palabras para detras y 20 palabras para delante

Programa SearchWord
	Formulario frmSearchWord
		variable wordInsert string
		BotonComando cmdBuscar
	Fin formulario
	Procedimiento cmdBuscar
		Hacer
			Iteracion de cada elemento del array
		Hasta (elemento == wordInsert)
	Fin procedimiento
Fin programa
******/


func SearchWord() []string{
	//go run main.go nombreArchivo palabraClave
	filePath := "default"
	word := os.Args[2]
	fmt.Println("word:",word)
	scanner := bufio.NewScanner(os.Stdin)

	if(len(os.Args) == 3){
		filePath = os.Args[1]
		fmt.Println("filePath:",filePath)
		file , err := os.Open(filePath)
		if err != nil{
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	var streamerLines[] string
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		if(strings.Contains(scanner.Text(), fmt.Sprint(word))){
			streamerLines = append(streamerLines, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}

	return streamerLines
}
