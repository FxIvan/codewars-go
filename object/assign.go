package object

import (
	"encoding/json"
	"fmt"
	"log"
)

type CommandLinux struct{
	ls			string
	cd			string
	cp			string
	mv			string
	cat_comm	string
}

type CommandWindow struct{
	dir			string
	cd			string
	copy		string
	move		string
	type_comm	string
}

func ObjectAssign(){
	commandLinux := `
	{
		"ls": "List the contents of a directory",
		"cd": "Change the current directory",
		"cp": "Copy files or directories",
		"mv": "Move or rename files or directories",
		"cat_comm": "Concatenate and display file content"
	}
	`

	commandWindows := `
	{
		"dir": "List the contents of a directory",
		"cd": "Change the current directory",
		"copy": "Copy files or directories",
		"move": "Move or rename files or directories",
		"type_comm": "Display the content of a file"
	}
	`

	var cw *CommandWindow
	var cl *CommandLinux

	err := json.Unmarshal([]byte(commandWindows), &cw)
	if(err != nil){
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(commandLinux), &cl)
	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println("Info comando:",cw,cl)
	
}

