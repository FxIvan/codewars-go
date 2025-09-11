package borrador

// INICIO Faberger
import (
	"fmt"
	"strconv"
	"sync"
)

type MaxMin struct {
	Init uint64
	Max  uint64
}

func goTo(min int, max int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := min; i < max; i++ {
		fmt.Println("Thread", id, ":", i)
	}
}

/*
	func goToTwo(min int, max int, id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := min; i < max; i++ {
			fmt.Println("Thread", id, ":", i)
		}
	}
*/
/*
func goTo(min int, max int, id int) {
	for i := min; i < max; i++ {
		fmt.Println("Thread", id, ":", i)
	}
}
*/
func UpsideDown(low, high string) uint64 {
	var wg sync.WaitGroup
	wg.Add(2)
	// ----------------------------------------

	count := uint64(0)
	lowBig, _ := strconv.Atoi(low)
	highBig, _ := strconv.Atoi(high)

	left := highBig / 2 // lowBig to left

	go goTo(lowBig, left, 1, &wg)
	go goTo(left, highBig, 2, &wg)

	/*
		goTo(lowBig, left, 1)
		goTo(left, highBig, 2)
	*/
	// ----------------------------------------
	wg.Wait()
	fmt.Println("Done!")

	return count
}

/*
var mirrorPairs = map[byte]byte{
	'0': '0', '1': '1', '6': '9', '8': '8', '9': '6',
}

// Función para contar upside-down numbers dentro de un rango
func UpsideDown(low, high string) uint64 {
	count := uint64(0)
	lowBig, _ := new(big.Int).SetString(low, 10)
	highBig, _ := new(big.Int).SetString(high, 10)

	for length := len(low); length <= len(high); length++ {
		count += generateUpsideDown("", length, lowBig, highBig)
	}

	return count
}

// Backtracking para generar números upside-down válidos
func generateUpsideDown(current string, targetLen int, low, high *big.Int) uint64 {
	if len(current) > targetLen {
		return 0
	}

	if len(current) == targetLen {
		// Verificamos si el número está dentro del rango
		num := new(big.Int)
		num.SetString(current, 10)
		if num.Cmp(low) >= 0 && num.Cmp(high) <= 0 {
			return 1
		}
		return 0
	}

	count := uint64(0)
	for digit, mirrored := range mirrorPairs {
		// No permitir números con 0 al inicio (excepto "0" solo)
		if len(current) == 0 && digit == '0' && targetLen > 1 {
			continue
		}

		next := string(digit) + current + string(mirrored)
		count += generateUpsideDown(next, targetLen, low, high)
	}

	return count
}
*/
// FIN Faberger
