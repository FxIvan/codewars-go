package montecarloPi

import (
	"math/rand"
	"time"
)

func MontecarloPi(numPoints int) float64 {
	rand.Seed(time.Now().UnixNano())
	insideCircle := 0

	for i := 0; i < numPoints; i++ {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1

		resultX := x * x
		resultY := y * y

		resultIntoCircle := resultX + resultY
		if resultIntoCircle <= 1 {
			insideCircle++
		}
	}

	insideCircleProporcion := float64(insideCircle) / float64(numPoints)
	return 4.0 * insideCircleProporcion
}
