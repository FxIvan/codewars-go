package faberger

import (
	"fmt"
	"strconv"
)

type MaxMin struct {
	Init uint64
	Max  uint64
}

func UpsideDown(low, high string) uint64 {

	var count uint64

	if high == "0" {
		return count
	}

	lowBig, _ := strconv.Atoi(low)
	highBig, _ := strconv.Atoi(high)

	fmt.Print("LowBig:", lowBig, "\n")
	fmt.Print("HighBig:", highBig, "\n")

	middBig := highBig / 2
	fmt.Print("MiddBig:", middBig, "\n")
	// UpsideDown(string(lowBig), string(middBig))

	return count
}
