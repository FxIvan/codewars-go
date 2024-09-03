package library

import (
	_"fmt"
)	

type CharacterSpecial struct {
	Character string
	ValueNumber string
}

func GetValueNumber(character string) string {
	arrayOfCharacter := [10]CharacterSpecial{
		{Character: "(", ValueNumber: "0"},
		{Character: ")", ValueNumber: "1"},
		{Character: "+", ValueNumber: "2"},
		{Character: "-", ValueNumber: "3"},
		{Character: "*", ValueNumber: "4"},
		{Character: "/", ValueNumber: "5"},
		{Character: "^", ValueNumber: "6"},
		{Character: "%", ValueNumber: "7"},
		{Character: "#", ValueNumber: "8"},
		{Character: "$", ValueNumber: "9"},
	}

	for _, item := range arrayOfCharacter {
		if item.Character == character {
			return item.ValueNumber
		}
	}

	return character
}