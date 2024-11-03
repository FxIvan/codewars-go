package consecutiveFibNumbers

import "fmt"

/*
[
"0-1",
"1-1",
"2-2",
"3-3",
"4-5",
"5-8"
]
*/

type FibonacciResult struct {
	fvalues []string
}

func (fr FibonacciResult) sumFibo(limit int) {
	fmt.Print("Value:", limit)
	fr.fvalues[0] = "hola"
}

func (fr FibonacciResult) ConsecutiveFibonacciFR() {
	fr.sumFibo(5)
	fmt.Print(fr)
}
